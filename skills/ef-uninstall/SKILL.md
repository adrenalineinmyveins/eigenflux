---
name: ef-uninstall
description: Remove an EigenFlux CLI installation, its owned native triggers, host plugin, and managed Skills when the owner requests uninstallation. Delete EigenFlux-owned local identity, credentials, data, and edited content.
metadata:
  author: "Phronesis AI"
  version: "1.0.3"
  requires:
    bins: ["eigenflux"]
---

# EigenFlux Uninstall

Handle owner requests to uninstall or remove EigenFlux, including “卸载 EigenFlux”. Require CLI 0.6.4. Use the stable Agent Home and resolved absolute CLI path.

Preview through the official `<CLI>.uninstall.sh` or `<CLI>.uninstall.ps1` beside the CLI. Pass the absolute CLI path first, then `--homedir`, `--server`, `--source agent`, and `--format json`. Inspect registered Homes, hosts, Skills targets, and preserved content. For a missing wrapper, recover ownership through CLI preview and run `installation record` with the verified host and Skills target. Use the registered wrapper paths; retain unknown paths.

Run the wrapper with `--begin` before native cleanup. Carry its ID through `--attempt`. Record native cleanup failures through `--record-result blocked` or `--record-result failed`.

Stop watch owners. Use host-native tools to remove owned EigenFlux triggers and plugins, matching task identity, Home, and server. Preserve unrelated tasks and shared configuration. Resolve ambiguous ownership with the owner.

Remove EigenFlux-owned local accounts, credentials, identity keys, caches, history, edited Skills, scripts, native triggers, and plugin configuration. Use `--all-homes` when scope covers every registration. Preserve unrelated host configuration.

After native cleanup succeeds, run the wrapper with `--apply --host-cleanup-confirmed`. Include `--notifications --notification-cleanup-confirmed` when scope covers all notification users. Preserve shared or changed notification assets. Verify Windows finalization after the original CLI exits.

Pass voluntary reason text through `--reason` and a known category through `--reason-category`. Keep text local until successful cleanup deletes the audit. Proceed without asking when no reason was supplied. Report upload failures separately; offline delivery after removal is not guaranteed.

Read back native state and verify executable removal. Verify local credentials, configuration, cached history, registered Skills, and scripts were deleted. Report permission failures and partial or blocked cleanup. Retain shared lock files and unrelated host content.
