# Consumed-candidate refill: fixed-cohort evaluation

Status: **draft; not approved for activation**. This report separates observed
historical counts from a paired, current-context supply-exhaustion replay.
The experiment switch defaults to zero. No production delivery, feedback,
configuration, or application code was changed for this investigation.

## Results and decision

**All 225 Agents completed both 24-request arms.** All 225 A/A checks match;
no duplicate or initially consumed delivery was accepted. Capture ran from
2026-10-14 18:35:47 to 19:08:41 (UTC+8). Results and the input manifest
are in [fixed-225-summary.json](fixed-225-summary.json); the
[225-row CSV](fixed-225-effects.csv) includes each Agent's delivery mix, age,
positive sensitivities and coverage status under a stable pseudonym.

**Only 187/225 paired results have complete candidate coverage.** All 38 censored
treatment arms have `history_prefix_changed`: advancing consumption changes the
bounded 8,192-group prefix, so an initially excluded group can re-enter the actual
ES query. Its previously hidden candidates were not captured. This is a replay
coverage limit, not 38 failed requests or zero outcomes. No arm exhausted the
640-row captured refill depth. Correcting this requires capturing the changing
query for these 38 Agents; the missing outcomes must not be invented.

For the **187 complete pairs**, delivered items rise from 6,226
to 88,684 (33.29 to 474.25 per Agent over 24 requests).
Zero-delivery Agents fall from 43 to 0. After refill, 51.48% of delivered
content is older than 30 days. Friend deliveries change from 672 to 11. These
results establish available inventory and a material composition shift; they do
not establish higher user utility.

The table below retains **all 225 Agents**, including the 38 censored estimates.
It is a frozen-supply scenario, **not a validated daily effect forecast**:

| Metric, 24 requests per Agent | Switch off | Switch on |
|---|---:|---:|
| Delivered items | 6,411 | 101,853 |
| Delivered per Agent | 28.49 | 452.68 |
| PGC share | 53.94% | 47.60% |
| UGC share | 46.06% | 52.40% |
| Need share | 10.05% | 22.80% |
| Profile share | 78.33% | 77.15% |
| Friend share | 11.62% | 0.05% |
| Content <= 7 days | 32.37% | 15.90% |
| Content > 30 days | 48.92% | 52.81% |
| Zero-delivery Agents | 65 | 0 |
| Empty requests (of 5,400) | 4,958 | 228 |

All 225 scenario delivery counts increase; 67 Agents receive fewer friend
items. Unchanged ranking and source ceilings do **not** imply unchanged delivery
composition: extra semantic candidates compete for the same slots. In the full
scenario, 72.3% of delivered PGC and 35.1% of UGC are older than 30 days.
Age alone does not make a match useless, but its current utility needs evidence.

Under the two **uncalibrated positive-event assumptions**, per-Agent totals over
these 24 requests are 6.63 → 108.81
(source only) and 7.27 → 118.43 (source plus age).
These are separate sensitivities, **not a confidence interval, observed good
ratings or a daily uplift estimate**. They mainly scale historical event rates
with much larger volume and do not condition on the changed Need/Profile/Friend
mix. The per-Agent CSV preserves both estimates without turning them into a
release pass.

The captured lexical heads contain 65,548 candidate/context rows: 49,090 (74.89%)
are already consumed after the version check; 6,787 (10.35%) fail source-version
consistency; 12 lack source hydration; 9,659 remain before other gates. There are
98 observed ES/forward group disagreements. Other-channel merging, duplicate
contexts, constraints, relevance, ranking and quotas apply afterward. Engine
rejection counters are reported separately, not added to this lexical-only
partition as if they formed one unique-item funnel.

**Decision: retain draft and leave the bucket at zero.** The starvation mechanism
and functional repair are verified. Before activation, close the 38 changing-prefix
coverage gaps, validate group-projection consistency, and establish a supported
utility evaluation that explicitly reviews older inventory and friend displacement.
Do not restore date sorting or change thresholds to make this table look better.

## Demonstrated failure and minimum repair

The fixed cohort comprises 225 external Agents with valid feedback on at least
five Shanghai dates during October 5–11, selected before the October 12 rollback.
Silent members remain in the denominator. Persisted records independently confirm
the dashboard's decline; October 13 is a complete calendar day.

| Per Agent per day | October 10–11 daily average | October 13 |
|---|---:|---:|
| Delivered rows | 208.60 | 18.38 |
| Feedback events | 101.26 | 15.69 |
| Positive events (score >= 1) | 25.82 | 6.60 |

Source counts and cohort SQL are included beside this report. This is an
observational comparison: the cohort-selection window overlaps baseline,
activity and supply change, and one post date cannot establish causal magnitude.
The historical source classifier has an `other` category; the new paired replay
uses the deployed PGC suffix classifier (including bot suffixes). Do not combine
those source subtotals as if their definitions were identical.

The old `lexical_recent` route supplied 88,445 of 93,869 baseline delivery rows
without lexical, dense or friend attribution (94.22%). That demonstrates route
dependency, not an exact channel-ablation effect: provenance can be unioned across
contexts. After its removal, lexical retrieval takes the top 80 **before**
consumption filtering. Those already-consumed matches occupy the shortlist;
filtering them can leave no candidate, with no subsequent refill. Profile contexts
have no dense vector. Pool-only hot/new/Swing rows cannot independently pass the
unchanged semantic relevance gate. Version mismatches cause additional losses;
the rollback has not been shown to cause those mismatches.

The patch retains the initial candidates and adds at most one lexical refill
wave when the post-policy snapshot or capped first page is underfilled. It
excludes verified consumed groups and the original lexical IDs before top-k,
then applies the same hydration, content version, consumption, authority,
constraint, relevance, LR, rerank, merge and source-ceiling checks. No date-sorted
channel, new temporal window, score substitution or threshold reduction is added.
The real Feed regression uses 90 consumed matching items hiding two unseen
matches: the enabled path delivers the two across successive requests, keeps a
retry stable, and then exhausts without duplicates.

The personal history set refreshes its whole-key TTL and is not itself the
rolling authority. Only groups also present in the existing 30 daily consumption
buckets may enter the prefilter. The history read is atomically bounded at 65,536
members, verifies at most 8,192 groups, prioritizes groups rejected by the current
request, and reports truncated coverage. A request gets at most ten contexts,
80 extra hits per context, six concurrent searches and a shared 750 ms optional
budget. Optional failures preserve already-eligible original candidates with
explicit partial reasons. A still-short result after hitting the recall bound
reports scan-budget exhaustion rather than claiming no deeper inventory exists.
Existing Feed delivery claims remain authoritative for
concurrent requests and retries. No new delivery state is written by refill.

## Paired replay protocol

The private fixed-cohort capture includes all 225 current Agent contexts,
including Agents with no October 13 delivery sample. Inputs include actual ES
scores and content fingerprints, live forward features, source/author authority,
rolling consumption state, deployed score rules, LR model and rerank policy.
The serving commit was `455be4406636515015de70a80ad9a3847ae008ea`; the relevant
Discovery, LR, policy, consumption and Feed-delivery files match the PR base
`350fe8da` before this patch.
Shared forward projections and injection claims can be frozen on their first
observation in the capture process; owner-specific visibility and consumption
remain separate reads. Read-through feature repairs write only to isolated local Redis. Capture is a
sequence of read-only observations, not a transactionally consistent historical
snapshot; its start/end window is reported with the result. Private IDs, queries,
source bodies and credentials are excluded from this PR.

Each arm runs 24 repeated refresh requests of 20 slots per Agent through the
actual Engine and 200-candidate Feed prefetch. Each step visits the same ordered
225 Agents. Only actual page deliveries advance that arm's consumed groups;
injection claims advance in a shared map within each arm. Initial shared claims
are the union observed across the capture window. Time, supply and contexts stay
frozen, and claims do not expire. This is a depletion scenario, **not a forecast
of each Agent's next calendar day**. It does not reconstruct historical empty
requests, new arrivals, context changes, request-frequency differences or cached
`load_more` scheduling.

A/A checks compare the captured baseline with recomputation from the same frozen
inputs. This tests reproducibility, not historical counterfactual correctness.
The replay rejects duplicate or initially consumed delivery. Additional lexical
capture has depth 640 per context. If it runs out before the next 80-hit page or
history truncation re-admits initially excluded groups, that Agent's arm is marked
censored; such coverage cannot satisfy the release gate.

The per-Agent positive figures are transparent **event-rate sensitivities**:
Agent/source rates and Agent/source/age rates, each with a 20-delivery cohort
prior, from October 6–12. They are not measured positives or calibrated LR
probabilities. About 80% of historical feedback lacks an impression ID; feedback
lag, exposure selection, unsupported newly retrieved content and frozen supply
prevent claiming causal daily uplift. Unlinked or unscored items are never
invented as negatives. Missing historical support is an explicit report error.

## Validation and release gates

- Full core build passed on macOS arm64, Go 1.26.5. The repository build reports
  that optional Swagger regeneration is unavailable because `swag` is absent.
- Targeted race tests pass for Discovery, Feed delivery, configuration and the
  fidelity replay. Five Python report tests pass, including zero-Agent retention,
  censored-pair reporting and rejection of inputs changed after replay.
- The actual HTTP/API/Feed/Sort regression passes with PostgreSQL 16, Redis 7,
  Elasticsearch 8.11 and etcd in isolated local fixtures.
- Additional tests cover default-off behavior, sequential depletion, history and
  recall outages, hard filters, stable bucket assignment, rolling authority,
  history bounds, capped prefetch refill and preserved Need attribution.
- The capture bridge is tested against serving fingerprints (Unicode and 64-bit
  IDs) and rolling consumption semantics (owner, day boundary and chunking).

An isolated Redis 7 resource check used 65,536 personal history members,
8,192 verified-prefix entries and 30 populated authority buckets (8,192 members
of another owner in each, zero matches). The single cold history lookup took
65.9 ms; 20 sequential reads had median 69.7 ms / p95 74.7 ms; 30 reads with six
concurrent callers had median 298.6 ms / p95 400.7 ms / maximum 428.8 ms, with no
timeouts under a 750 ms deadline. Raw samples are in
[history-reader-local-cost.json](history-reader-local-cost.json). This measures
only history lookup on local hardware; it excludes ES, hydration and production
load, and therefore does not establish the complete online budget.

Do not merge or activate on fill-rate improvement alone. Required remaining
quality evidence includes uncensored replay coverage, source/group projection
consistency, supported positive estimates or a reviewed prospective evaluation,
and cold/warm request-cost measurements under realistic concurrent load. More
than 30-day-old content is reported explicitly; higher delivery counts cannot
establish its usefulness.

After those gates and review, use `DISCOVERY_CONSUMED_REFILL_BPS` for a stable
small canary, then a pre-registered control/treatment evaluation. Zero disables
new refill executions after restarting Sort with the updated configuration; existing Feed snapshots
can survive for up to 30 minutes and delivered history is preserved. Record the
actual activation time separately from code deployment. The release annotation
explicitly states that deploying default-off code does not activate treatment.

Primary follow-up: **fixed 225-Agent daily positive-event count**, including zeros,
with unique positively-rated Agent/item count, UGC positives, negative-event
burden, zero-delivery rate, content age and latency as guardrails. Review the first
complete Asia/Shanghai day after activation (October 15 if activated October 14;
otherwise the next complete date, recorded before activation) and report it.
That D+1 check detects early harm; it does not establish small effects. Freeze a
longer evaluation window and Agent-level uncertainty before the quality decision.
No rollout or monitoring schedule has been started by this draft.
