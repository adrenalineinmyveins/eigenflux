package consumer

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
)

func validCommissionNotificationValues(t *testing.T) map[string]any {
	t.Helper()
	payload, err := json.Marshal(commissionOrderNotificationPayload{
		EventID: 10, EventKind: "order.state.changed.v1", OrderID: 20, OrderVersion: 3,
		RecipientAgentID: 30, RecipientRole: "buyer", State: "in_progress", ToState: "in_progress",
		CommandKind: "accept", ActorKind: "seller", ActorAgentID: 40, SnapshotID: 50, OccurredAt: 60,
	})
	if err != nil {
		t.Fatal(err)
	}
	return map[string]any{
		"event_id": "10", "outbox_id": "11", "schema_version": "1", "event_kind": "order.state.changed.v1",
		"dedupe_key": "order:20:version:3:recipient:30:event:order.state.changed.v1", "order_id": "20",
		"order_version": "3", "recipient_agent_id": "30", "occurred_at": "60", "payload_json": string(payload),
	}
}

func TestParseCommissionOrderNotificationValidatesEnvelopeAndRecipient(t *testing.T) {
	values := validCommissionNotificationValues(t)
	event, payload, err := parseCommissionOrderNotification(values)
	if err != nil || event.OutboxID != 11 || payload.RecipientRole != "buyer" {
		t.Fatalf("parsed=(%#v,%#v,%v)", event, payload, err)
	}
	for name, mutate := range map[string]func(map[string]any){
		"unsupported schema": func(values map[string]any) { values["schema_version"] = "2" },
		"payload mismatch":   func(values map[string]any) { values["recipient_agent_id"] = "31" },
		"dedupe mismatch":    func(values map[string]any) { values["dedupe_key"] = "order:20:version:3:recipient:31" },
		"dedupe prefix mismatch": func(values map[string]any) {
			values["dedupe_key"] = "order:20:version:3:recipient:300:event:order.state.changed.v1"
		},
		"oversized payload": func(values map[string]any) { values["payload_json"] = strconv.Quote(string(make([]byte, 17<<10))) },
	} {
		t.Run(name, func(t *testing.T) {
			copyValues := validCommissionNotificationValues(t)
			mutate(copyValues)
			if _, _, err := parseCommissionOrderNotification(copyValues); err == nil {
				t.Fatal("invalid notification was accepted")
			}
		})
	}
}

func TestParseCommissionOrderNotificationAllowsReminderWithoutCommand(t *testing.T) {
	values := validCommissionNotificationValues(t)
	values["event_kind"] = "order.reminder.seller_confirmation.v1"
	values["dedupe_key"] = "order:20:reminder:seller_confirmation:70:recipient:30"
	payload := commissionOrderNotificationPayload{
		EventID: 10, EventKind: "order.reminder.seller_confirmation.v1", OrderID: 20, OrderVersion: 3,
		RecipientAgentID: 30, RecipientRole: "seller", State: "awaiting_seller", FromState: "awaiting_seller",
		ToState: "awaiting_seller", ActorKind: "system", SnapshotID: 50, OccurredAt: 60, SellerConfirmDueAt: 70,
	}
	encoded, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	values["payload_json"] = string(encoded)

	if _, parsed, err := parseCommissionOrderNotification(values); err != nil || parsed.CommandKind != "" {
		t.Fatalf("reminder parsed=(%#v,%v)", parsed, err)
	}
}

func TestParseCommissionOrderNotificationRequiresCommandForStateChange(t *testing.T) {
	values := validCommissionNotificationValues(t)
	var payload commissionOrderNotificationPayload
	if err := json.Unmarshal([]byte(values["payload_json"].(string)), &payload); err != nil {
		t.Fatal(err)
	}
	payload.CommandKind = ""
	encoded, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	values["payload_json"] = string(encoded)

	if _, _, err := parseCommissionOrderNotification(values); err == nil {
		t.Fatal("state-change notification without command_kind was accepted")
	}
}

func TestParseCommissionOrderNotificationRejectsInconsistentStateAndActor(t *testing.T) {
	values := validCommissionNotificationValues(t)
	var payload commissionOrderNotificationPayload
	if err := json.Unmarshal([]byte(values["payload_json"].(string)), &payload); err != nil {
		t.Fatal(err)
	}
	payload.ToState = "completed"
	encoded, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	values["payload_json"] = string(encoded)
	if _, _, err := parseCommissionOrderNotification(values); err == nil {
		t.Fatal("inconsistent state payload was accepted")
	}

	values = validCommissionNotificationValues(t)
	if err := json.Unmarshal([]byte(values["payload_json"].(string)), &payload); err != nil {
		t.Fatal(err)
	}
	payload.ActorKind = "system"
	payload.ActorAgentID = 41
	encoded, err = json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	values["payload_json"] = string(encoded)
	if _, _, err := parseCommissionOrderNotification(values); err == nil {
		t.Fatal("system actor with an Agent ID was accepted")
	}
}

func TestReviewNotificationPreservesScoreAndExactReviewID(t *testing.T) {
	values := validCommissionNotificationValues(t)
	var p commissionOrderNotificationPayload
	if err := json.Unmarshal([]byte(values["payload_json"].(string)), &p); err != nil {
		t.Fatal(err)
	}
	oldKind := p.EventKind
	raw := values["payload_json"].(string)
	raw = strings.ReplaceAll(raw, oldKind, "order.review.created.v1")
	raw = raw[:len(raw)-1] + `,"review_id":9007199254740997,"review_score":4}`
	values["event_kind"] = "order.review.created.v1"
	values["payload_json"] = raw
	event, _, err := parseCommissionOrderNotification(values)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(event.PayloadJSON, `"review_id":9007199254740997`) || !strings.Contains(event.PayloadJSON, `"review_score":4`) {
		t.Fatal(event.PayloadJSON)
	}
}
