# CLI installation removal

`eigenflux uninstall` previews the registered installation without writing files or contacting the network. `--begin` records an attempt before native cleanup; resume it through `--attempt`. Apply requires `--host-cleanup-confirmed`. Multiple registered Homes require `--all-homes`.

Use host-native tools to stop and remove EigenFlux triggers, plugins, and their configuration. Retain unrelated host settings and third-party content. The CLI checks installation, update, and watch locks before removal. It rejects linked or non-canonical Agent Home roots before deleting data.

Successful removal deletes the registered CLI, its owned sidecars and wrappers, registered EigenFlux-managed Skills including edited content, and each registered Home's local configuration, credentials, identity keys, caches, history, maintenance state, and uninstall audits. Skills targets require a valid EigenFlux ownership manifest; only registered entries are removed. Leaf symlinks are unlinked without removing their targets. Local data cleanup does not delete a cloud account.

Terminal upload clients capture existing credentials in memory before cleanup.
Successful cleanup deletes local files before uploading the actual terminal outcome; it does not recreate credentials or queues. Free-text reasons remain local during an attempt and are deleted after successful cleanup. Upload has one five-second budget across accounts and does not block removal. Offline uninstall has no retained delivery queue after successful data cleanup.

Installation/update/watch lock files remain to prevent replacement of live lock inodes. Successful removal retains no business data or credentials in those files. Partial or blocked removal keeps retry/audit state and reports failure accurately.

Windows reports pending removal until the original process exits. The official wrapper copies a trusted reporter to a temporary directory. Finalization reacquires locks, verifies receipt file hashes, removes original files, captures terminal upload credentials, clears local data, and reports the actual outcome. An intervening update or reinstallation blocks deletion. The temporary reporter is then removed. Direct CLI invocation leaves removal pending.

Notification cleanup remains explicit with `--notifications --notification-cleanup-confirmed` because the integration may be shared. Native tools remove EigenFlux-owned modified notification assets when appropriate. Shared PATH entries and OS permission records are not removed.

Verification compiles and runs isolated CLI copies. Tests cover immutable preview, multi-Home scope, active locks, edited managed content removal, foreign content retention, symlink boundaries, local data cleanup, and official wrapper execution. Keep embedded wrappers byte-identical to `static/`.
