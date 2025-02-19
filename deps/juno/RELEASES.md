# Release and Upgrade procedures and communications

The `#validators-private` channel on discord will be used for all communications
from the team. Only **active validators** should be allowed access, for security
reasons.

**The core team will endeavor to always make sure there is 48-72 hours notice of
an impending upgrade unless there is no alternative.**

Most of our validator communications are done on the
[Juno Discord](https://discord.gg/Juno). You should join, and change your server
name to `nick | validator-name`, then ask a mod for permission to see the
private validator channels.

## Release versioning

**If a change crosses a major version number, i.e. `3.x.x -> 4.x.x` then it is
consensus-breaking.**

In the past, some releases have been consensus-breaking but only incremented a
minor version, if indicated. In the future, we will look to be clearer. the
**Only patch versions, i.e. `x.x.1 -> x.x.2`, or `3.1.0 -> 3.1.1` are guaranteed
to be non-consensus breaking.**

## Scheduled upgrade via governance

For a SoftwareUpgradeProposal via governance:

1.  Validators will be told via the announcements channel when the prop is live
2.  Validators will be told via the announcements channel if it passes
3.  Validators will be told via the announcements channel when the upgrade
    instructions are available, and the upgrade will be coordinated in the
    private validators channel as the target upgrade block nears.

## Emergency upgrade or security patch

If the team needs to upgrade the chain faster than the cadence of governance,
then a special procedure applies.

This procedure minimizes the amount that is publicly shared about a potential
issue.

1.  An announcement calling validators to check in on the private validators
    channel will be posted on the *validator's announcement channel* on discord.
    No specifics will be shared here, as it is public.
2.  Details of the patch and the upgrade plan will be shared on the private
    channel, as well as an expected ETA.
3.  When instructions are available, they will be pinned, and a second
    announcement sent on the announcements channel. A thread for acknowledgments
    will be created for validators to signal readiness.
4.  The team will compile a spreadsheet of validator readiness to check we are
    past 67%.

There are two further considerations:

1.  If the change is consensus-breaking, a halt height will be applied and
    validators will manually upgrade at that block, after the halt.
2.  If the change is non-consensus breaking, validators will apply when ready,
    and then signal readiness.

## Syncing from genesis

The team will be putting together instructions that will be kept up-to-date for
syncing without using a backup.

Note that pre-Lupercalia (`< v3.0.0`) is effectively a different chain, as the
old-style CosmosHub upgrade path was used, meaning there is a disjoint of one
block.
