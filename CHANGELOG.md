# Change Log

## [Unreleased](https://github.com/asteris-llc/converge/tree/HEAD)

[Full Changelog](https://github.com/asteris-llc/converge/compare/0.6.0-beta1...HEAD)

### Closed Pull Requests

- unarchive: update sample hcl to use file.directory [\#606](https://github.com/asteris-llc/converge/pull/606) ([arichardet](https://github.com/arichardet))

## [0.6.0-beta1](https://github.com/asteris-llc/converge/tree/0.6.0-beta1) (2017-02-24)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.5.0...0.6.0-beta1)

### Enhancements

- Add a status method to raise level if there are differences [\#586](https://github.com/asteris-llc/converge/issues/586)
- unarchive resource [\#513](https://github.com/asteris-llc/converge/issues/513)
- Improve User Diffs [\#444](https://github.com/asteris-llc/converge/issues/444)
- Add enabled/disabled to user [\#277](https://github.com/asteris-llc/converge/issues/277)
- make name available to module calls [\#112](https://github.com/asteris-llc/converge/issues/112)
- Benchmarking operations [\#24](https://github.com/asteris-llc/converge/issues/24)
- Feature/file.unarchive [\#592](https://github.com/asteris-llc/converge/pull/592) ([arichardet](https://github.com/arichardet))
- have wercker install script on get.converge.sh [\#579](https://github.com/asteris-llc/converge/pull/579) ([stevendborrelli](https://github.com/stevendborrelli))
- Add various generators [\#571](https://github.com/asteris-llc/converge/pull/571) ([BrianHicks](https://github.com/BrianHicks))

### Bugs

- update pip package to fix failing example [\#580](https://github.com/asteris-llc/converge/pull/580) ([stevendborrelli](https://github.com/stevendborrelli))

### Closed Issues

- Autocompletion and man pages [\#560](https://github.com/asteris-llc/converge/issues/560)
- Refactor User Delete to use Diffs [\#445](https://github.com/asteris-llc/converge/issues/445)

### Closed Pull Requests

- Feature/systemd unit state [\#602](https://github.com/asteris-llc/converge/pull/602) ([rebeccaskinner](https://github.com/rebeccaskinner))
- no longer xcompile freebsd & solaris [\#601](https://github.com/asteris-llc/converge/pull/601) ([arichardet](https://github.com/arichardet))
- load: remove addGroupDependenciesToGroup [\#599](https://github.com/asteris-llc/converge/pull/599) ([BrianHicks](https://github.com/BrianHicks))
- refactor readWrite to use io.Copy [\#596](https://github.com/asteris-llc/converge/pull/596) ([rebeccaskinner](https://github.com/rebeccaskinner))
- use home\_dir value for user add diffs [\#588](https://github.com/asteris-llc/converge/pull/588) ([arichardet](https://github.com/arichardet))
- add status method-raise level to will change when there are diffs [\#587](https://github.com/asteris-llc/converge/pull/587) ([arichardet](https://github.com/arichardet))
- Feature/userdel diffs [\#584](https://github.com/asteris-llc/converge/pull/584) ([arichardet](https://github.com/arichardet))
- Feature/user expiry [\#581](https://github.com/asteris-llc/converge/pull/581) ([arichardet](https://github.com/arichardet))
- Update installer to 0.5.0 [\#577](https://github.com/asteris-llc/converge/pull/577) ([stevendborrelli](https://github.com/stevendborrelli))
- add converge logo with name to images [\#576](https://github.com/asteris-llc/converge/pull/576) ([arichardet](https://github.com/arichardet))
- Sleep in Graph Generation Test [\#570](https://github.com/asteris-llc/converge/pull/570) ([BrianHicks](https://github.com/BrianHicks))

## [0.5.0](https://github.com/asteris-llc/converge/tree/0.5.0) (2016-12-29)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.5.0-rc1...0.5.0)

### Enhancements

- Handle time.Time in preparer [\#490](https://github.com/asteris-llc/converge/issues/490)
- k8s demo [\#480](https://github.com/asteris-llc/converge/issues/480)
- no mo' spaces [\#320](https://github.com/asteris-llc/converge/issues/320)
- We're using too many ports [\#286](https://github.com/asteris-llc/converge/issues/286)
- better status display on plan and execution [\#269](https://github.com/asteris-llc/converge/issues/269)
- builtin file owner module [\#74](https://github.com/asteris-llc/converge/issues/74)
- demo: kubernetes - use latest version of converge [\#524](https://github.com/asteris-llc/converge/pull/524) ([ryane](https://github.com/ryane))
- fix shasum filename [\#522](https://github.com/asteris-llc/converge/pull/522) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/improve output [\#516](https://github.com/asteris-llc/converge/pull/516) ([arichardet](https://github.com/arichardet))

### Bugs

- docker.container panics if image not present [\#538](https://github.com/asteris-llc/converge/issues/538)
- Return error instead of message from docker resource plan [\#529](https://github.com/asteris-llc/converge/issues/529)
- misleading error message in converge graph command [\#520](https://github.com/asteris-llc/converge/issues/520)
- exit code is 0 when converge plan has errors [\#492](https://github.com/asteris-llc/converge/issues/492)
- error specifying file mode [\#487](https://github.com/asteris-llc/converge/issues/487)
- docker.network misleading error messages [\#485](https://github.com/asteris-llc/converge/issues/485)
- inconsistent indentation on multi-line diffs [\#478](https://github.com/asteris-llc/converge/issues/478)
- pass typed params through to modules [\#409](https://github.com/asteris-llc/converge/issues/409)
- Mutable TaskStatus results in apply losing differences [\#374](https://github.com/asteris-llc/converge/issues/374)
- auth code is repeated [\#273](https://github.com/asteris-llc/converge/issues/273)
- Remove indent for multi-line diffs [\#530](https://github.com/asteris-llc/converge/pull/530) ([arichardet](https://github.com/arichardet))
- status: add MayChange status and Warning display [\#528](https://github.com/asteris-llc/converge/pull/528) ([ryane](https://github.com/ryane))
- Fix/disable docker solaris [\#527](https://github.com/asteris-llc/converge/pull/527) ([ryane](https://github.com/ryane))
- plan cmd: return non-zero exit code when plan has error\(s\) [\#525](https://github.com/asteris-llc/converge/pull/525) ([ryane](https://github.com/ryane))
- graph cmd: fix wrong argument length error [\#521](https://github.com/asteris-llc/converge/pull/521) ([ryane](https://github.com/ryane))
- Fix/file.mode errors [\#519](https://github.com/asteris-llc/converge/pull/519) ([ryane](https://github.com/ryane))

### Closed Issues

- do not output private fields in docs [\#552](https://github.com/asteris-llc/converge/issues/552)
- Better output for cascading failures [\#367](https://github.com/asteris-llc/converge/issues/367)
- file download/fetch module [\#250](https://github.com/asteris-llc/converge/issues/250)

### Closed Pull Requests

- do not merge: release commit for 0.5.0 [\#574](https://github.com/asteris-llc/converge/pull/574) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Highlight HCL [\#573](https://github.com/asteris-llc/converge/pull/573) ([BrianHicks](https://github.com/BrianHicks))
- Update docs style with new branding [\#572](https://github.com/asteris-llc/converge/pull/572) ([BrianHicks](https://github.com/BrianHicks))
- add logo to Readme [\#532](https://github.com/asteris-llc/converge/pull/532) ([stevendborrelli](https://github.com/stevendborrelli))
- update readme with graph and download info [\#518](https://github.com/asteris-llc/converge/pull/518) ([stevendborrelli](https://github.com/stevendborrelli))
- kubernetes example [\#507](https://github.com/asteris-llc/converge/pull/507) ([ryane](https://github.com/ryane))

## [0.5.0-rc1](https://github.com/asteris-llc/converge/tree/0.5.0-rc1) (2016-12-28)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.5.0-beta1...0.5.0-rc1)

### Closed Issues

- task.query behavior [\#567](https://github.com/asteris-llc/converge/issues/567)
- Get rid of vendor directory [\#545](https://github.com/asteris-llc/converge/issues/545)

### Closed Pull Requests

- no merge: update changelog [\#569](https://github.com/asteris-llc/converge/pull/569) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docs: only list exported fields in param and export sections [\#561](https://github.com/asteris-llc/converge/pull/561) ([arichardet](https://github.com/arichardet))
- Fix Linting Bugs [\#554](https://github.com/asteris-llc/converge/pull/554) ([BrianHicks](https://github.com/BrianHicks))

## [0.5.0-beta1](https://github.com/asteris-llc/converge/tree/0.5.0-beta1) (2016-12-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0...0.5.0-beta1)

### Enhancements

- Add interceptors for RPC auth [\#531](https://github.com/asteris-llc/converge/pull/531) ([BrianHicks](https://github.com/BrianHicks))

### Bugs

- LVM FS module reports failure due to still having changes [\#555](https://github.com/asteris-llc/converge/issues/555)
- handle newlines in docs generation for exported fields [\#548](https://github.com/asteris-llc/converge/issues/548)
- docker: update docker dependency [\#512](https://github.com/asteris-llc/converge/pull/512) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- changelog branch; tag and push tags before merging [\#559](https://github.com/asteris-llc/converge/pull/559) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Minor Fixes After Makefile rewrite [\#558](https://github.com/asteris-llc/converge/pull/558) ([BrianHicks](https://github.com/BrianHicks))
- reset updates needed after apply in lvm [\#557](https://github.com/asteris-llc/converge/pull/557) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Makefile: remove ref to VERSION in PACKAGE\_VERSION [\#556](https://github.com/asteris-llc/converge/pull/556) ([BrianHicks](https://github.com/BrianHicks))
- Better Makefiles [\#553](https://github.com/asteris-llc/converge/pull/553) ([BrianHicks](https://github.com/BrianHicks))
- when apply fails due to changes still being present; use those diffs … [\#551](https://github.com/asteris-llc/converge/pull/551) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docs: use same formatting for \(re\)exported fields as parameters [\#550](https://github.com/asteris-llc/converge/pull/550) ([arichardet](https://github.com/arichardet))
- Feature/74 file owner [\#549](https://github.com/asteris-llc/converge/pull/549) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/file.fetch [\#543](https://github.com/asteris-llc/converge/pull/543) ([arichardet](https://github.com/arichardet))
- Fix/538 docker container panic [\#542](https://github.com/asteris-llc/converge/pull/542) ([rebeccaskinner](https://github.com/rebeccaskinner))
- return a better error for planned network failures [\#541](https://github.com/asteris-llc/converge/pull/541) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Timer status display [\#537](https://github.com/asteris-llc/converge/pull/537) ([BrianHicks](https://github.com/BrianHicks))
- Add fuzzing and benchmarks to CI [\#536](https://github.com/asteris-llc/converge/pull/536) ([BrianHicks](https://github.com/BrianHicks))
- disallow some characters as resource names [\#535](https://github.com/asteris-llc/converge/pull/535) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use a single port [\#534](https://github.com/asteris-llc/converge/pull/534) ([BrianHicks](https://github.com/BrianHicks))
- Fix/374 immutable status [\#533](https://github.com/asteris-llc/converge/pull/533) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/docs [\#526](https://github.com/asteris-llc/converge/pull/526) ([arichardet](https://github.com/arichardet))
- Feature/preparer time.time [\#523](https://github.com/asteris-llc/converge/pull/523) ([arichardet](https://github.com/arichardet))

## [0.4.0](https://github.com/asteris-llc/converge/tree/0.4.0) (2016-11-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-rc1...0.4.0)

### Enhancements

- Add `nonempty` tag to preparer [\#486](https://github.com/asteris-llc/converge/issues/486)
- apt module [\#404](https://github.com/asteris-llc/converge/issues/404)
- Handle time.Duration in Preparer [\#358](https://github.com/asteris-llc/converge/issues/358)
- binary distribution [\#306](https://github.com/asteris-llc/converge/issues/306)
- docker.volume resource [\#299](https://github.com/asteris-llc/converge/issues/299)
- docker.network resource [\#298](https://github.com/asteris-llc/converge/issues/298)
- Add ability to modify user [\#278](https://github.com/asteris-llc/converge/issues/278)
- LVM module [\#270](https://github.com/asteris-llc/converge/issues/270)
- update installer to point to 0.4.0 [\#517](https://github.com/asteris-llc/converge/pull/517) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/preparer time.duration [\#483](https://github.com/asteris-llc/converge/pull/483) ([arichardet](https://github.com/arichardet))
- check rpm for empty packages [\#482](https://github.com/asteris-llc/converge/pull/482) ([stevendborrelli](https://github.com/stevendborrelli))
- only extract binary from tarball [\#481](https://github.com/asteris-llc/converge/pull/481) ([stevendborrelli](https://github.com/stevendborrelli))
- docker.network resource [\#477](https://github.com/asteris-llc/converge/pull/477) ([ryane](https://github.com/ryane))
- Feature/apt package [\#461](https://github.com/asteris-llc/converge/pull/461) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docker.volume [\#453](https://github.com/asteris-llc/converge/pull/453) ([ryane](https://github.com/ryane))
- Add the ability to modify a user [\#434](https://github.com/asteris-llc/converge/pull/434) ([arichardet](https://github.com/arichardet))

### Bugs

- validate non-empty string values on some resources [\#337](https://github.com/asteris-llc/converge/issues/337)
- Intermittent test failure with `make test` [\#476](https://github.com/asteris-llc/converge/issues/476)
- Converge loses dependencies between modules within a branch when  [\#427](https://github.com/asteris-llc/converge/issues/427)
- diff output is partially bold [\#402](https://github.com/asteris-llc/converge/issues/402)
- Apply doesn't show diff output [\#399](https://github.com/asteris-llc/converge/issues/399)
- values that don't implement Stringer log badly [\#398](https://github.com/asteris-llc/converge/issues/398)
- resolving conditional macros line [\#389](https://github.com/asteris-llc/converge/issues/389)
- Case predicates fail when they include `lookup` [\#386](https://github.com/asteris-llc/converge/issues/386)
- converge exit code on a failed run [\#365](https://github.com/asteris-llc/converge/issues/365)
- Address name conflicts [\#361](https://github.com/asteris-llc/converge/issues/361)
- resources should have the ability to gracefully stop on interrupt [\#319](https://github.com/asteris-llc/converge/issues/319)
- converge panics when interrupted [\#318](https://github.com/asteris-llc/converge/issues/318)
- Fix/conditional nodes [\#484](https://github.com/asteris-llc/converge/pull/484) ([rebeccaskinner](https://github.com/rebeccaskinner))
- grouping: fix issues with groups and conditionals [\#460](https://github.com/asteris-llc/converge/pull/460) ([ryane](https://github.com/ryane))
- return non-zero exit code if apply graph has an error [\#436](https://github.com/asteris-llc/converge/pull/436) ([ryane](https://github.com/ryane))
- replace "context" with "golang.org/x/net/context" [\#433](https://github.com/asteris-llc/converge/pull/433) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- update changelog, add script [\#511](https://github.com/asteris-llc/converge/pull/511) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: fix link in resource authors guide [\#479](https://github.com/asteris-llc/converge/pull/479) ([ryane](https://github.com/ryane))
- Fix/resource duplication [\#458](https://github.com/asteris-llc/converge/pull/458) ([arichardet](https://github.com/arichardet))
- Fix/graceful exit [\#447](https://github.com/asteris-llc/converge/pull/447) ([arichardet](https://github.com/arichardet))
- LVM module [\#184](https://github.com/asteris-llc/converge/pull/184) ([avnik](https://github.com/avnik))

## [0.4.0-rc1](https://github.com/asteris-llc/converge/tree/0.4.0-rc1) (2016-11-17)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-beta1...0.4.0-rc1)

### Closed Pull Requests

- update changelog to reflect rc1 status [\#510](https://github.com/asteris-llc/converge/pull/510) ([rebeccaskinner](https://github.com/rebeccaskinner))
- move rendering plant creation inside of the transform block to avoid … [\#509](https://github.com/asteris-llc/converge/pull/509) ([rebeccaskinner](https://github.com/rebeccaskinner))
- set github link to the converge repo [\#508](https://github.com/asteris-llc/converge/pull/508) ([rebeccaskinner](https://github.com/rebeccaskinner))
- empty commit for CI [\#506](https://github.com/asteris-llc/converge/pull/506) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.4.0-beta1](https://github.com/asteris-llc/converge/tree/0.4.0-beta1) (2016-11-15)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0...0.4.0-beta1)

### Enhancements

- file.owner resource [\#503](https://github.com/asteris-llc/converge/issues/503)

### Bugs

- package.rpm state does not default to present [\#446](https://github.com/asteris-llc/converge/issues/446)
- Error when adding a user when a group of the same name exists [\#425](https://github.com/asteris-llc/converge/issues/425)

### Closed Issues

- nonempty check in preparer for user/group [\#489](https://github.com/asteris-llc/converge/issues/489)

### Closed Pull Requests

- Fix/0.4.0 beta1 fix [\#505](https://github.com/asteris-llc/converge/pull/505) ([rebeccaskinner](https://github.com/rebeccaskinner))
- release notes [\#497](https://github.com/asteris-llc/converge/pull/497) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use %v for logging [\#496](https://github.com/asteris-llc/converge/pull/496) ([arichardet](https://github.com/arichardet))
- update docs with conditional changes [\#495](https://github.com/asteris-llc/converge/pull/495) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/nonempty tag [\#494](https://github.com/asteris-llc/converge/pull/494) ([arichardet](https://github.com/arichardet))
- Remove extract executable [\#488](https://github.com/asteris-llc/converge/pull/488) ([arichardet](https://github.com/arichardet))
- Manifest / index page [\#464](https://github.com/asteris-llc/converge/pull/464) ([BrianHicks](https://github.com/BrianHicks))
- default to state 'present' for rpm module [\#463](https://github.com/asteris-llc/converge/pull/463) ([rebeccaskinner](https://github.com/rebeccaskinner))
- vendor: upgrade all deps [\#459](https://github.com/asteris-llc/converge/pull/459) ([BrianHicks](https://github.com/BrianHicks))

## [0.3.0](https://github.com/asteris-llc/converge/tree/0.3.0) (2016-10-27)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-rc1...0.3.0)

### Enhancements

- "param is required" error should include param name [\#335](https://github.com/asteris-llc/converge/issues/335)
- Add ability to modify group [\#279](https://github.com/asteris-llc/converge/issues/279)
- add conditionals to module resource [\#268](https://github.com/asteris-llc/converge/issues/268)
- named locks [\#249](https://github.com/asteris-llc/converge/issues/249)
- pretty printed changes should align values [\#244](https://github.com/asteris-llc/converge/issues/244)
- ability to wait for a condition [\#238](https://github.com/asteris-llc/converge/issues/238)
- proper arrays [\#110](https://github.com/asteris-llc/converge/issues/110)
- makefile: add hashes to files [\#440](https://github.com/asteris-llc/converge/pull/440) ([BrianHicks](https://github.com/BrianHicks))
- \[389\] switch log level from Info to Debug [\#428](https://github.com/asteris-llc/converge/pull/428) ([mason-fish](https://github.com/mason-fish))
- Feature/examples 0.3.0 [\#419](https://github.com/asteris-llc/converge/pull/419) ([ryane](https://github.com/ryane))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- named groups [\#392](https://github.com/asteris-llc/converge/pull/392) ([ryane](https://github.com/ryane))
- Don't render module dependencies [\#387](https://github.com/asteris-llc/converge/pull/387) ([BrianHicks](https://github.com/BrianHicks))
- Added docs, renamed packages to package to  for consistency with other modules [\#384](https://github.com/asteris-llc/converge/pull/384) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use status.RaiseLevel in group [\#377](https://github.com/asteris-llc/converge/pull/377) ([arichardet](https://github.com/arichardet))
- Feature/rpm module [\#373](https://github.com/asteris-llc/converge/pull/373) ([rebeccaskinner](https://github.com/rebeccaskinner))
- use a container for CI [\#372](https://github.com/asteris-llc/converge/pull/372) ([BrianHicks](https://github.com/BrianHicks))
- create a metadata envelope for nodes [\#369](https://github.com/asteris-llc/converge/pull/369) ([BrianHicks](https://github.com/BrianHicks))
- cmd/server.go: use errgroup instead of waitgroup [\#363](https://github.com/asteris-llc/converge/pull/363) ([QuentinPerez](https://github.com/QuentinPerez))
- Feature/conditionals [\#362](https://github.com/asteris-llc/converge/pull/362) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/example swarm wait [\#346](https://github.com/asteris-llc/converge/pull/346) ([ryane](https://github.com/ryane))
- Add Compound Parameters [\#340](https://github.com/asteris-llc/converge/pull/340) ([BrianHicks](https://github.com/BrianHicks))
- load overlay module in elk example [\#326](https://github.com/asteris-llc/converge/pull/326) ([feniix](https://github.com/feniix))
- Refactor deserializers [\#321](https://github.com/asteris-llc/converge/pull/321) ([BrianHicks](https://github.com/BrianHicks))
- Use text/tabwriter to align human output [\#317](https://github.com/asteris-llc/converge/pull/317) ([sehqlr](https://github.com/sehqlr))
- Fix/225 pipeline function refactor [\#307](https://github.com/asteris-llc/converge/pull/307) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Bugs

- Looking up a node from a branch that is also a part of a group introduced a deadlock [\#415](https://github.com/asteris-llc/converge/issues/415)
- user/group not activating changes [\#417](https://github.com/asteris-llc/converge/issues/417)
- docker.container reporting incorrect entrypoint diff [\#410](https://github.com/asteris-llc/converge/issues/410)
- converge panics when encountering an empty conditional [\#407](https://github.com/asteris-llc/converge/issues/407)
- Use lists as params in modules is broken [\#397](https://github.com/asteris-llc/converge/issues/397)
- Explicit dependencies fail inside of case statements [\#385](https://github.com/asteris-llc/converge/issues/385)
- Use `package.rpm` not `rpm.package` for rpm module [\#382](https://github.com/asteris-llc/converge/issues/382)
- shell module should not set StatusLevel based on process exit code [\#323](https://github.com/asteris-llc/converge/issues/323)
- StatusLevel not taken into account during graph execution [\#322](https://github.com/asteris-llc/converge/issues/322)
- param dependency fails in `samples/shellContext.hcl` [\#313](https://github.com/asteris-llc/converge/issues/313)
- Execution Engine Ignoring Warning Levels [\#243](https://github.com/asteris-llc/converge/issues/243)
- Make pipeline functions use mult-return instead of Either [\#225](https://github.com/asteris-llc/converge/issues/225)
- Fix 420 go test race [\#424](https://github.com/asteris-llc/converge/pull/424) ([ryane](https://github.com/ryane))
- Replace `AreSiblingIDs` with `AreSiblings` to prevent deadlocking [\#423](https://github.com/asteris-llc/converge/pull/423) ([rebeccaskinner](https://github.com/rebeccaskinner))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- Handle thunked branches [\#403](https://github.com/asteris-llc/converge/pull/403) ([rebeccaskinner](https://github.com/rebeccaskinner))
- don't exclude modules in getNearestAncestor [\#396](https://github.com/asteris-llc/converge/pull/396) ([ryane](https://github.com/ryane))
- Fix/385 dependencies in conditionals [\#391](https://github.com/asteris-llc/converge/pull/391) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix lookup calls to use os/user in SetAddUserOptions [\#390](https://github.com/asteris-llc/converge/pull/390) ([arichardet](https://github.com/arichardet))
- Change `rpm.package` to `package.rpm` [\#383](https://github.com/asteris-llc/converge/pull/383) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Update user status level and errors, add checks in Apply for group [\#375](https://github.com/asteris-llc/converge/pull/375) ([arichardet](https://github.com/arichardet))
- Ability to use pointers as a preparer value [\#339](https://github.com/asteris-llc/converge/pull/339) ([BrianHicks](https://github.com/BrianHicks))
- Status Error Codes [\#333](https://github.com/asteris-llc/converge/pull/333) ([BrianHicks](https://github.com/BrianHicks))
- fix codeclimate yaml [\#328](https://github.com/asteris-llc/converge/pull/328) ([BrianHicks](https://github.com/BrianHicks))

### Closed Issues

- docs for rpm module [\#381](https://github.com/asteris-llc/converge/issues/381)

### Closed Pull Requests

- update for 0.3.0 [\#443](https://github.com/asteris-llc/converge/pull/443) ([stevendborrelli](https://github.com/stevendborrelli))
- prep for 0.3.0-rc1 [\#437](https://github.com/asteris-llc/converge/pull/437) ([stevendborrelli](https://github.com/stevendborrelli))
- update release notes [\#435](https://github.com/asteris-llc/converge/pull/435) ([stevendborrelli](https://github.com/stevendborrelli))
- beta3 release notes [\#431](https://github.com/asteris-llc/converge/pull/431) ([stevendborrelli](https://github.com/stevendborrelli))
- update docs to include currently identified switch statement limitations [\#388](https://github.com/asteris-llc/converge/pull/388) ([rebeccaskinner](https://github.com/rebeccaskinner))
- update vendored dependencies [\#378](https://github.com/asteris-llc/converge/pull/378) ([BrianHicks](https://github.com/BrianHicks))
- 0.3.0 release documentation [\#376](https://github.com/asteris-llc/converge/pull/376) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: add multi-version control [\#371](https://github.com/asteris-llc/converge/pull/371) ([BrianHicks](https://github.com/BrianHicks))
- add docs for the platform module [\#342](https://github.com/asteris-llc/converge/pull/342) ([stevendborrelli](https://github.com/stevendborrelli))
- update for PR \#307 [\#316](https://github.com/asteris-llc/converge/pull/316) ([stevendborrelli](https://github.com/stevendborrelli))

## [0.3.0-rc1](https://github.com/asteris-llc/converge/tree/0.3.0-rc1) (2016-10-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta3...0.3.0-rc1)

### Closed Issues

- Update documentation regarding conditionals and groups [\#430](https://github.com/asteris-llc/converge/issues/430)

## [0.3.0-beta3](https://github.com/asteris-llc/converge/tree/0.3.0-beta3) (2016-10-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta2...0.3.0-beta3)

### Enhancements

- add installation script [\#416](https://github.com/asteris-llc/converge/pull/416) ([stevendborrelli](https://github.com/stevendborrelli))

### Bugs

- go tests race condition [\#420](https://github.com/asteris-llc/converge/issues/420)

### Closed Pull Requests

- document known bug between conditionals and groups [\#432](https://github.com/asteris-llc/converge/pull/432) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.3.0-beta2](https://github.com/asteris-llc/converge/tree/0.3.0-beta2) (2016-10-24)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta1...0.3.0-beta2)

### Bugs

- Fix container entrypoint diffs [\#412](https://github.com/asteris-llc/converge/pull/412) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- wercker: don't delete removed [\#414](https://github.com/asteris-llc/converge/pull/414) ([BrianHicks](https://github.com/BrianHicks))
- Nightly/test builds [\#413](https://github.com/asteris-llc/converge/pull/413) ([BrianHicks](https://github.com/BrianHicks))
- fix panic [\#408](https://github.com/asteris-llc/converge/pull/408) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.3.0-beta1](https://github.com/asteris-llc/converge/tree/0.3.0-beta1) (2016-10-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/before-moving-resolve-conditionals...0.3.0-beta1)

### Enhancements

- file/dir refactor [\#327](https://github.com/asteris-llc/converge/issues/327)
- Use errgroup in server [\#295](https://github.com/asteris-llc/converge/issues/295)
- Allow ability to indicate group name when adding a user [\#276](https://github.com/asteris-llc/converge/issues/276)
- builtin file group module [\#75](https://github.com/asteris-llc/converge/issues/75)

### Bugs

- Conditional Regression [\#401](https://github.com/asteris-llc/converge/issues/401)
- module dependencies now failing [\#395](https://github.com/asteris-llc/converge/issues/395)
- docker.container regression [\#343](https://github.com/asteris-llc/converge/issues/343)
- samples in the README are formatted incorrectly. [\#104](https://github.com/asteris-llc/converge/issues/104)

### Closed Issues

- Tidy up codebase by fixing `make lint` reports [\#348](https://github.com/asteris-llc/converge/issues/348)
- Can we simplify the visualization of large modules? [\#280](https://github.com/asteris-llc/converge/issues/280)
- How should we manage module docs and examples? [\#173](https://github.com/asteris-llc/converge/issues/173)

### Closed Pull Requests

- Allow lists in parameters when called in a module [\#400](https://github.com/asteris-llc/converge/pull/400) ([BrianHicks](https://github.com/BrianHicks))
- comment typo [\#380](https://github.com/asteris-llc/converge/pull/380) ([rebeccaskinner](https://github.com/rebeccaskinner))
- FIX local command flag added to example [\#370](https://github.com/asteris-llc/converge/pull/370) ([philcryer](https://github.com/philcryer))
- codeclimate: remove markdown linting [\#368](https://github.com/asteris-llc/converge/pull/368) ([BrianHicks](https://github.com/BrianHicks))
- Add testing section to contribution guidelines [\#366](https://github.com/asteris-llc/converge/pull/366) ([arichardet](https://github.com/arichardet))
- remove binary [\#354](https://github.com/asteris-llc/converge/pull/354) ([stevendborrelli](https://github.com/stevendborrelli))
- wait: fix error handling in retrier [\#353](https://github.com/asteris-llc/converge/pull/353) ([ryane](https://github.com/ryane))
- Feature/param name in error [\#352](https://github.com/asteris-llc/converge/pull/352) ([ryane](https://github.com/ryane))
- docs: add section on error values to resource author's guide [\#347](https://github.com/asteris-llc/converge/pull/347) ([BrianHicks](https://github.com/BrianHicks))
- Fix docker regressions [\#345](https://github.com/asteris-llc/converge/pull/345) ([ryane](https://github.com/ryane))
- wait.query and wait.port resources [\#334](https://github.com/asteris-llc/converge/pull/334) ([ryane](https://github.com/ryane))

## [before-moving-resolve-conditionals](https://github.com/asteris-llc/converge/tree/before-moving-resolve-conditionals) (2016-10-05)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0...before-moving-resolve-conditionals)

### Enhancements

- Use github.com/pkg/errors exclusively [\#300](https://github.com/asteris-llc/converge/issues/300)

### Bugs

- handle parameters with valid zero value [\#338](https://github.com/asteris-llc/converge/issues/338)

### Closed Issues

- Universal Benchmarking [\#308](https://github.com/asteris-llc/converge/issues/308)
- Use table tests instead of helpers.PreparerValidator [\#155](https://github.com/asteris-llc/converge/issues/155)

### Closed Pull Requests

- Handle preparer fields where zero is valid [\#344](https://github.com/asteris-llc/converge/pull/344) ([arichardet](https://github.com/arichardet))
- plan: test, which illustrate "return nil, err" problem [\#336](https://github.com/asteris-llc/converge/pull/336) ([avnik](https://github.com/avnik))
- Support for modifying linux groups [\#331](https://github.com/asteris-llc/converge/pull/331) ([arichardet](https://github.com/arichardet))
- Add contributing.md and code of conduct [\#330](https://github.com/asteris-llc/converge/pull/330) ([BrianHicks](https://github.com/BrianHicks))
- shell: non-0 exit code returns StatusWillChange [\#324](https://github.com/asteris-llc/converge/pull/324) ([ryane](https://github.com/ryane))
- map keys are considered "strings" in parse.Node [\#315](https://github.com/asteris-llc/converge/pull/315) ([rebeccaskinner](https://github.com/rebeccaskinner))
- helpers,docker: move AssertDiff to common testhelpers [\#312](https://github.com/asteris-llc/converge/pull/312) ([avnik](https://github.com/avnik))

## [0.2.0](https://github.com/asteris-llc/converge/tree/0.2.0) (2016-09-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-rc1...0.2.0)

### Enhancements

- Calls to `lookup` should be all lower-case [\#223](https://github.com/asteris-llc/converge/issues/223)

### Bugs

- field name conflicts during lookups [\#294](https://github.com/asteris-llc/converge/issues/294)
- panic when trying to lookup against file.directory [\#292](https://github.com/asteris-llc/converge/issues/292)
- Allow adding and deleting a group without specifying gid [\#275](https://github.com/asteris-llc/converge/issues/275)
- param lookup race condition [\#266](https://github.com/asteris-llc/converge/issues/266)
- Node is sometimes unresolvable, depending on ordering within a template [\#254](https://github.com/asteris-llc/converge/issues/254)
- params should handle boolean values [\#251](https://github.com/asteris-llc/converge/issues/251)
- remove token handling logic when only a client and not a server [\#247](https://github.com/asteris-llc/converge/issues/247)
- Check shouldn't be called from Apply [\#240](https://github.com/asteris-llc/converge/issues/240)
- value passing causes dependency resolution failures in child modules [\#239](https://github.com/asteris-llc/converge/issues/239)
- task.Task has extra argument [\#237](https://github.com/asteris-llc/converge/issues/237)
- Lookups fail for file.content [\#236](https://github.com/asteris-llc/converge/issues/236)
- RPC isn't used in healthchecks [\#199](https://github.com/asteris-llc/converge/issues/199)

### Closed Issues

- module author guide [\#291](https://github.com/asteris-llc/converge/issues/291)
- logging is a little too chatty [\#248](https://github.com/asteris-llc/converge/issues/248)
- Add RPC auth documentation [\#246](https://github.com/asteris-llc/converge/issues/246)

### Closed Pull Requests

- Add ability to indicate the group name when adding a user [\#310](https://github.com/asteris-llc/converge/pull/310) ([arichardet](https://github.com/arichardet))
- 0.2.0 release [\#311](https://github.com/asteris-llc/converge/pull/311) ([stevendborrelli](https://github.com/stevendborrelli))
- Add "Basic Usage" section to Server docs [\#287](https://github.com/asteris-llc/converge/pull/287) ([sehqlr](https://github.com/sehqlr))
- Simplify Status implementation [\#284](https://github.com/asteris-llc/converge/pull/284) ([BrianHicks](https://github.com/BrianHicks))
- Fix climate issues [\#282](https://github.com/asteris-llc/converge/pull/282) ([BrianHicks](https://github.com/BrianHicks))
- Add codeclimate checks [\#281](https://github.com/asteris-llc/converge/pull/281) ([BrianHicks](https://github.com/BrianHicks))
- Make the logs quieter [\#274](https://github.com/asteris-llc/converge/pull/274) ([BrianHicks](https://github.com/BrianHicks))
- example: elasticsearch, kibana, and filebeat ~elk [\#272](https://github.com/asteris-llc/converge/pull/272) ([ryane](https://github.com/ryane))
- param: accept boolean values [\#271](https://github.com/asteris-llc/converge/pull/271) ([BrianHicks](https://github.com/BrianHicks))
- Use `MakeLanguage` for dependecy resolution to ensure that no [\#265](https://github.com/asteris-llc/converge/pull/265) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Make notes match sample code [\#262](https://github.com/asteris-llc/converge/pull/262) ([tomduckering](https://github.com/tomduckering))

## [0.2.0-rc1](https://github.com/asteris-llc/converge/tree/0.2.0-rc1) (2016-09-23)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-beta2...0.2.0-rc1)

### Bugs

- docker.image incorrectly reporting changes after apply [\#304](https://github.com/asteris-llc/converge/issues/304)

### Closed Issues

- preprocessor race condition [\#302](https://github.com/asteris-llc/converge/issues/302)

### Closed Pull Requests

- Fixes docker resources and updates examples [\#305](https://github.com/asteris-llc/converge/pull/305) ([ryane](https://github.com/ryane))
- use thread-safe field cache [\#303](https://github.com/asteris-llc/converge/pull/303) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/294 field name conflicts [\#301](https://github.com/asteris-llc/converge/pull/301) ([rebeccaskinner](https://github.com/rebeccaskinner))
- preproc: don't add fields for non-struct anon field [\#293](https://github.com/asteris-llc/converge/pull/293) ([ryane](https://github.com/ryane))
- docs: add draft resource authors guide [\#290](https://github.com/asteris-llc/converge/pull/290) ([BrianHicks](https://github.com/BrianHicks))
- perform healthchecks over RPC [\#289](https://github.com/asteris-llc/converge/pull/289) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/239 [\#288](https://github.com/asteris-llc/converge/pull/288) ([rebeccaskinner](https://github.com/rebeccaskinner))
- cmd: don't set local token if only a client [\#285](https://github.com/asteris-llc/converge/pull/285) ([BrianHicks](https://github.com/BrianHicks))
- Fix/user group - Allow adding/deleting without gid [\#283](https://github.com/asteris-llc/converge/pull/283) ([arichardet](https://github.com/arichardet))

## [0.2.0-beta2](https://github.com/asteris-llc/converge/tree/0.2.0-beta2) (2016-09-20)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-beta1...0.2.0-beta2)

### Enhancements

- module signing [\#16](https://github.com/asteris-llc/converge/issues/16)

### Bugs

- This graph is rendered kinda funky [\#141](https://github.com/asteris-llc/converge/issues/141)

### Closed Issues

- Update all deps to the latest versions [\#171](https://github.com/asteris-llc/converge/issues/171)
- human printer should be profiled and optimized [\#156](https://github.com/asteris-llc/converge/issues/156)

### Closed Pull Requests

- example: docker swarm mode [\#267](https://github.com/asteris-llc/converge/pull/267) ([ryane](https://github.com/ryane))

## [0.2.0-beta1](https://github.com/asteris-llc/converge/tree/0.2.0-beta1) (2016-09-16)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.1.1...0.2.0-beta1)

### Enhancements

- Create a query module [\#227](https://github.com/asteris-llc/converge/issues/227)
- Graphing should work over RPC [\#207](https://github.com/asteris-llc/converge/issues/207)
- Allow Value-passing between modules [\#206](https://github.com/asteris-llc/converge/issues/206)
- logging needs improvement [\#200](https://github.com/asteris-llc/converge/issues/200)
- REST layer on RPC [\#170](https://github.com/asteris-llc/converge/issues/170)
- CLI RPC client [\#169](https://github.com/asteris-llc/converge/issues/169)
- RPC authentication [\#168](https://github.com/asteris-llc/converge/issues/168)
- RPC server [\#167](https://github.com/asteris-llc/converge/issues/167)
- Docker module [\#159](https://github.com/asteris-llc/converge/issues/159)
- Clear/readable errors [\#140](https://github.com/asteris-llc/converge/issues/140)
- Macros? [\#111](https://github.com/asteris-llc/converge/issues/111)
- should template be called file.content? [\#103](https://github.com/asteris-llc/converge/issues/103)
- execution in new graph code [\#92](https://github.com/asteris-llc/converge/issues/92)
- execution planning in new graph code [\#91](https://github.com/asteris-llc/converge/issues/91)
- implement file.group [\#90](https://github.com/asteris-llc/converge/issues/90)
- reimplement file.user [\#89](https://github.com/asteris-llc/converge/issues/89)
- reimplement file.mode [\#88](https://github.com/asteris-llc/converge/issues/88)
- reimplement templates [\#87](https://github.com/asteris-llc/converge/issues/87)
- reimplement shell tasks [\#86](https://github.com/asteris-llc/converge/issues/86)
- implement graph visualization on top of new graph code [\#85](https://github.com/asteris-llc/converge/issues/85)
- resources should keep track of their parents [\#83](https://github.com/asteris-llc/converge/issues/83)
- infer dependencies based on template contents [\#65](https://github.com/asteris-llc/converge/issues/65)
- "native" filemode module [\#40](https://github.com/asteris-llc/converge/issues/40)
- "native" modules addressing scheme [\#39](https://github.com/asteris-llc/converge/issues/39)
- basic diff finding [\#20](https://github.com/asteris-llc/converge/issues/20)
- authenticated module server [\#18](https://github.com/asteris-llc/converge/issues/18)
- module server [\#17](https://github.com/asteris-llc/converge/issues/17)
- HTTP fetching [\#13](https://github.com/asteris-llc/converge/issues/13)
- graph: keep track of parentage inside edges [\#163](https://github.com/asteris-llc/converge/pull/163) ([BrianHicks](https://github.com/BrianHicks))

### Bugs

- lookup should be fully case insensitive [\#232](https://github.com/asteris-llc/converge/issues/232)
- rpc server does not support graph operation anymore [\#205](https://github.com/asteris-llc/converge/issues/205)
- graph is walked out of order [\#160](https://github.com/asteris-llc/converge/issues/160)
- Concurrent Map Read/Write Panic When Running Tests [\#158](https://github.com/asteris-llc/converge/issues/158)
- Shell Doesn't Validate Against User-Defined Interpreter [\#119](https://github.com/asteris-llc/converge/issues/119)
- invalid script syntax doesn't return a good error [\#113](https://github.com/asteris-llc/converge/issues/113)
- Blackbox/test\_apply\_remote.sh Fails [\#102](https://github.com/asteris-llc/converge/issues/102)
- `helpers` is kind of an awful module name [\#96](https://github.com/asteris-llc/converge/issues/96)
- Failing tasks don't block tasks that depend on them [\#82](https://github.com/asteris-llc/converge/issues/82)
- all modules in samples must be valid and useful [\#66](https://github.com/asteris-llc/converge/issues/66)
- params should not have default dependencies [\#62](https://github.com/asteris-llc/converge/issues/62)
- resource names with dashes break graph rendering [\#61](https://github.com/asteris-llc/converge/issues/61)
- Fix error handling when fail to print results during plan or apply [\#183](https://github.com/asteris-llc/converge/pull/183) ([arichardet](https://github.com/arichardet))
- Return better error information for Shell module failures [\#121](https://github.com/asteris-llc/converge/pull/121) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Closed Issues

- errors from check silently swallowed during apply [\#213](https://github.com/asteris-llc/converge/issues/213)
- errors are printed with full stacks [\#208](https://github.com/asteris-llc/converge/issues/208)
- nil pointer panic when file.content resource errors [\#179](https://github.com/asteris-llc/converge/issues/179)
- intermittent concurrent map write panic with dependent params [\#176](https://github.com/asteris-llc/converge/issues/176)
- use standard library net/context library [\#161](https://github.com/asteris-llc/converge/issues/161)
- Subgraph PrettyPrinter should use `nil` as the `SubgraphBottomID` [\#149](https://github.com/asteris-llc/converge/issues/149)
- DigraphPrettyPrinter \(prettyprinters\) should shrink [\#147](https://github.com/asteris-llc/converge/issues/147)
- Renderable \(prettyprinters\) is larger than necessary [\#146](https://github.com/asteris-llc/converge/issues/146)
- Shell command context [\#145](https://github.com/asteris-llc/converge/issues/145)
- Add a 'Health' Module [\#144](https://github.com/asteris-llc/converge/issues/144)
- Builtin file.directory module [\#137](https://github.com/asteris-llc/converge/issues/137)
- Graphviz Generation Fails With Empty Param [\#129](https://github.com/asteris-llc/converge/issues/129)
- parameters should be required if no default is set [\#123](https://github.com/asteris-llc/converge/issues/123)
- Shell Module Error Codes Could Have Better Formatting [\#120](https://github.com/asteris-llc/converge/issues/120)
- params should allow commas [\#114](https://github.com/asteris-llc/converge/issues/114)
- modules and params should not show up by default [\#106](https://github.com/asteris-llc/converge/issues/106)
- use a context in loading [\#94](https://github.com/asteris-llc/converge/issues/94)
- builtin file owner and file permissions module [\#84](https://github.com/asteris-llc/converge/issues/84)

### Closed Pull Requests

- We got an assigned port so here's docs. [\#261](https://github.com/asteris-llc/converge/pull/261) ([BrianHicks](https://github.com/BrianHicks))
- update readme [\#260](https://github.com/asteris-llc/converge/pull/260) ([stevendborrelli](https://github.com/stevendborrelli))
- Add user support [\#259](https://github.com/asteris-llc/converge/pull/259) ([arichardet](https://github.com/arichardet))
- MOAR DOCS PLS [\#258](https://github.com/asteris-llc/converge/pull/258) ([BrianHicks](https://github.com/BrianHicks))
- Makefile: refine xcompile targets [\#257](https://github.com/asteris-llc/converge/pull/257) ([BrianHicks](https://github.com/BrianHicks))
- Fixed issue where apply swallowed error. [\#256](https://github.com/asteris-llc/converge/pull/256) ([Dacode45](https://github.com/Dacode45))
- add mutex to render [\#255](https://github.com/asteris-llc/converge/pull/255) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Add default group implementation for unsupported systems [\#252](https://github.com/asteris-llc/converge/pull/252) ([arichardet](https://github.com/arichardet))
- Feature/module verification [\#245](https://github.com/asteris-llc/converge/pull/245) ([ajhager](https://github.com/ajhager))
- Feature/fully lowercase field names [\#242](https://github.com/asteris-llc/converge/pull/242) ([rebeccaskinner](https://github.com/rebeccaskinner))
- file.directory resource [\#241](https://github.com/asteris-llc/converge/pull/241) ([BrianHicks](https://github.com/BrianHicks))
- Add user group support [\#234](https://github.com/asteris-llc/converge/pull/234) ([arichardet](https://github.com/arichardet))
- support lowercase field names [\#231](https://github.com/asteris-llc/converge/pull/231) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/query refactor [\#230](https://github.com/asteris-llc/converge/pull/230) ([rebeccaskinner](https://github.com/rebeccaskinner))
- handle params that are thunked [\#229](https://github.com/asteris-llc/converge/pull/229) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/query module [\#228](https://github.com/asteris-llc/converge/pull/228) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Refactor helpers into descriptive names [\#222](https://github.com/asteris-llc/converge/pull/222) ([BrianHicks](https://github.com/BrianHicks))
- Feature/value passing [\#220](https://github.com/asteris-llc/converge/pull/220) ([rebeccaskinner](https://github.com/rebeccaskinner))
- rpc: flatten details field [\#218](https://github.com/asteris-llc/converge/pull/218) ([BrianHicks](https://github.com/BrianHicks))
- Feature/formatter \#208 [\#217](https://github.com/asteris-llc/converge/pull/217) ([BrianHicks](https://github.com/BrianHicks))
- docker.container: fix entrypoint diff logic [\#215](https://github.com/asteris-llc/converge/pull/215) ([ryane](https://github.com/ryane))
- Graphing over RPC [\#214](https://github.com/asteris-llc/converge/pull/214) ([BrianHicks](https://github.com/BrianHicks))
- docker.container resource [\#203](https://github.com/asteris-llc/converge/pull/203) ([ryane](https://github.com/ryane))
- Rework logs to use logrus [\#202](https://github.com/asteris-llc/converge/pull/202) ([BrianHicks](https://github.com/BrianHicks))
- Feature/platform refactor [\#197](https://github.com/asteris-llc/converge/pull/197) ([stevendborrelli](https://github.com/stevendborrelli))
- Added host environment variable support [\#195](https://github.com/asteris-llc/converge/pull/195) ([arichardet](https://github.com/arichardet))
- resources\(docker image\): document [\#193](https://github.com/asteris-llc/converge/pull/193) ([BrianHicks](https://github.com/BrianHicks))
- Documentation Site [\#192](https://github.com/asteris-llc/converge/pull/192) ([BrianHicks](https://github.com/BrianHicks))
- Feature/health module [\#189](https://github.com/asteris-llc/converge/pull/189) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docker.image resource redux [\#188](https://github.com/asteris-llc/converge/pull/188) ([ryane](https://github.com/ryane))
- Basic RPC [\#187](https://github.com/asteris-llc/converge/pull/187) ([BrianHicks](https://github.com/BrianHicks))
- shell env and working dir support [\#185](https://github.com/asteris-llc/converge/pull/185) ([ryane](https://github.com/ryane))
- \[GRAPH\] Fix \#158, add triggering files, add tests [\#181](https://github.com/asteris-llc/converge/pull/181) ([sehqlr](https://github.com/sehqlr))
- don't panic when result status is nil [\#180](https://github.com/asteris-llc/converge/pull/180) ([ryane](https://github.com/ryane))
- Feature/optimize human printer [\#175](https://github.com/asteris-llc/converge/pull/175) ([arichardet](https://github.com/arichardet))
- Shrink binaries with some magical ldflags [\#174](https://github.com/asteris-llc/converge/pull/174) ([BrianHicks](https://github.com/BrianHicks))
- Feature/shell refactor [\#172](https://github.com/asteris-llc/converge/pull/172) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Dependency ordered graph walking [\#166](https://github.com/asteris-llc/converge/pull/166) ([BrianHicks](https://github.com/BrianHicks))
- Migrate context for go1.7, refactor fetch/http.go [\#164](https://github.com/asteris-llc/converge/pull/164) ([sehqlr](https://github.com/sehqlr))
- Feature/resource return refactor [\#157](https://github.com/asteris-llc/converge/pull/157) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Move errors into the tree [\#154](https://github.com/asteris-llc/converge/pull/154) ([BrianHicks](https://github.com/BrianHicks))
- Human-readable pretty printer [\#152](https://github.com/asteris-llc/converge/pull/152) ([BrianHicks](https://github.com/BrianHicks))
- make nil a useful default bottom value for subgraphdi [\#150](https://github.com/asteris-llc/converge/pull/150) ([rebeccaskinner](https://github.com/rebeccaskinner))
- JSONL output and interface refactoring [\#148](https://github.com/asteris-llc/converge/pull/148) ([BrianHicks](https://github.com/BrianHicks))
- Feature/template split [\#136](https://github.com/asteris-llc/converge/pull/136) ([rebeccaskinner](https://github.com/rebeccaskinner))
- fetch: add context [\#135](https://github.com/asteris-llc/converge/pull/135) ([BrianHicks](https://github.com/BrianHicks))
- Render merged subtrees [\#133](https://github.com/asteris-llc/converge/pull/133) ([BrianHicks](https://github.com/BrianHicks))
- Fix/119 [\#132](https://github.com/asteris-llc/converge/pull/132) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Trim Duplicate Subtrees [\#131](https://github.com/asteris-llc/converge/pull/131) ([BrianHicks](https://github.com/BrianHicks))
- Appropriately handle required parameters instead of crashing [\#130](https://github.com/asteris-llc/converge/pull/130) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use concurrent map implementation for graph values [\#128](https://github.com/asteris-llc/converge/pull/128) ([BrianHicks](https://github.com/BrianHicks))
- Make params missing defaults required [\#127](https://github.com/asteris-llc/converge/pull/127) ([BrianHicks](https://github.com/BrianHicks))
- makefile: simplify linting [\#126](https://github.com/asteris-llc/converge/pull/126) ([BrianHicks](https://github.com/BrianHicks))
- Rename "template" to "file.content" [\#124](https://github.com/asteris-llc/converge/pull/124) ([BrianHicks](https://github.com/BrianHicks))
- gometalinter fixes [\#122](https://github.com/asteris-llc/converge/pull/122) ([BrianHicks](https://github.com/BrianHicks))
- Graphviz [\#118](https://github.com/asteris-llc/converge/pull/118) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Rewrite graph system [\#116](https://github.com/asteris-llc/converge/pull/116) ([BrianHicks](https://github.com/BrianHicks))
- resource\(shell\): add interpreter [\#109](https://github.com/asteris-llc/converge/pull/109) ([BrianHicks](https://github.com/BrianHicks))
- switch to glide [\#108](https://github.com/asteris-llc/converge/pull/108) ([BrianHicks](https://github.com/BrianHicks))
- Feature/context [\#107](https://github.com/asteris-llc/converge/pull/107) ([BrianHicks](https://github.com/BrianHicks))
- Format samples in readme [\#105](https://github.com/asteris-llc/converge/pull/105) ([ajhager](https://github.com/ajhager))
- Port shell resource and tests [\#101](https://github.com/asteris-llc/converge/pull/101) ([ajhager](https://github.com/ajhager))
- resource\(template\): implement and test template [\#98](https://github.com/asteris-llc/converge/pull/98) ([BrianHicks](https://github.com/BrianHicks))
- Application for preparers [\#97](https://github.com/asteris-llc/converge/pull/97) ([BrianHicks](https://github.com/BrianHicks))
- Plan in new preparers [\#95](https://github.com/asteris-llc/converge/pull/95) ([BrianHicks](https://github.com/BrianHicks))
- dependencies should block execution [\#81](https://github.com/asteris-llc/converge/pull/81) ([siddharthist](https://github.com/siddharthist))
- dependency tracker [\#80](https://github.com/asteris-llc/converge/pull/80) ([BrianHicks](https://github.com/BrianHicks))
- Self Serve [\#79](https://github.com/asteris-llc/converge/pull/79) ([BrianHicks](https://github.com/BrianHicks))
- Module server [\#78](https://github.com/asteris-llc/converge/pull/78) ([BrianHicks](https://github.com/BrianHicks))
- File mode module [\#76](https://github.com/asteris-llc/converge/pull/76) ([BrianHicks](https://github.com/BrianHicks))
- Simplify Application [\#73](https://github.com/asteris-llc/converge/pull/73) ([BrianHicks](https://github.com/BrianHicks))
- Fix graph with dashes [\#71](https://github.com/asteris-llc/converge/pull/71) ([BrianHicks](https://github.com/BrianHicks))
- load: from http [\#23](https://github.com/asteris-llc/converge/pull/23) ([siddharthist](https://github.com/siddharthist))

## [0.1.1](https://github.com/asteris-llc/converge/tree/0.1.1) (2016-06-13)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.1...0.1.1)

### Enhancements

- log lines from DAG walking should not appear [\#67](https://github.com/asteris-llc/converge/issues/67)

### Bugs

- color is not working? [\#64](https://github.com/asteris-llc/converge/issues/64)
- recursive modules are valid [\#63](https://github.com/asteris-llc/converge/issues/63)

### Closed Pull Requests

- Feature/logging [\#70](https://github.com/asteris-llc/converge/pull/70) ([BrianHicks](https://github.com/BrianHicks))
- add color printing again [\#69](https://github.com/asteris-llc/converge/pull/69) ([siddharthist](https://github.com/siddharthist))
- Feature/basic blackbox testing [\#68](https://github.com/asteris-llc/converge/pull/68) ([BrianHicks](https://github.com/BrianHicks))

## [0.1](https://github.com/asteris-llc/converge/tree/0.1) (2016-06-09)
[Full Changelog](https://github.com/asteris-llc/converge/compare/Unreleased...0.1)

### Enhancements

- converge format [\#55](https://github.com/asteris-llc/converge/issues/55)
- colorized summary lines [\#45](https://github.com/asteris-llc/converge/issues/45)
- print a nice summary of changes \(or no changes\) after commands [\#35](https://github.com/asteris-llc/converge/issues/35)
- Progressive output in check/apply [\#34](https://github.com/asteris-llc/converge/issues/34)
- Supplying params is hard [\#33](https://github.com/asteris-llc/converge/issues/33)
- Template should accept file permissions [\#31](https://github.com/asteris-llc/converge/issues/31)
- add net/context to Check [\#30](https://github.com/asteris-llc/converge/issues/30)
- templating in param defaults [\#27](https://github.com/asteris-llc/converge/issues/27)
- passing in parameters to top-level module [\#26](https://github.com/asteris-llc/converge/issues/26)
- Apply [\#15](https://github.com/asteris-llc/converge/issues/15)
- HTTPS fetching [\#14](https://github.com/asteris-llc/converge/issues/14)
- Nicer plan output [\#12](https://github.com/asteris-llc/converge/issues/12)
- Checker [\#11](https://github.com/asteris-llc/converge/issues/11)
- Requirements [\#10](https://github.com/asteris-llc/converge/issues/10)

### Bugs

- duplicate names are lost in the graph [\#56](https://github.com/asteris-llc/converge/issues/56)
- fill in short and long descriptions [\#51](https://github.com/asteris-llc/converge/issues/51)
- don't ignore error when Viper binds pflags [\#48](https://github.com/asteris-llc/converge/issues/48)
- param names should have limited characters [\#41](https://github.com/asteris-llc/converge/issues/41)
- template permissions should be 600 by default [\#38](https://github.com/asteris-llc/converge/issues/38)
- exec.Check tests [\#21](https://github.com/asteris-llc/converge/issues/21)
- Incomplete comment  [\#5](https://github.com/asteris-llc/converge/issues/5)

### Closed Issues

- README is out of date [\#52](https://github.com/asteris-llc/converge/issues/52)

### Closed Pull Requests

- Update 0.1 docs [\#60](https://github.com/asteris-llc/converge/pull/60) ([BrianHicks](https://github.com/BrianHicks))
- fixed souceFile.hcl [\#59](https://github.com/asteris-llc/converge/pull/59) ([Dacode45](https://github.com/Dacode45))
- Bug/names [\#58](https://github.com/asteris-llc/converge/pull/58) ([Dacode45](https://github.com/Dacode45))
- fmt [\#57](https://github.com/asteris-llc/converge/pull/57) ([BrianHicks](https://github.com/BrianHicks))
- Feature/requirements [\#54](https://github.com/asteris-llc/converge/pull/54) ([Dacode45](https://github.com/Dacode45))
- don't ignore the error when Viper binds to PFlags [\#50](https://github.com/asteris-llc/converge/pull/50) ([siddharthist](https://github.com/siddharthist))
- color summaries [\#49](https://github.com/asteris-llc/converge/pull/49) ([siddharthist](https://github.com/siddharthist))
- add key=value parameter parsing [\#47](https://github.com/asteris-llc/converge/pull/47) ([siddharthist](https://github.com/siddharthist))
- Summary lines [\#46](https://github.com/asteris-llc/converge/pull/46) ([BrianHicks](https://github.com/BrianHicks))
- render: fix swallowed errors [\#44](https://github.com/asteris-llc/converge/pull/44) ([BrianHicks](https://github.com/BrianHicks))
- limit identifiers to word chars, dashes, and dots [\#43](https://github.com/asteris-llc/converge/pull/43) ([BrianHicks](https://github.com/BrianHicks))
- resource\(template\): fix permissions [\#42](https://github.com/asteris-llc/converge/pull/42) ([BrianHicks](https://github.com/BrianHicks))
- Lazily load parameters - defaults can be templated [\#37](https://github.com/asteris-llc/converge/pull/37) ([siddharthist](https://github.com/siddharthist))
- commands: add streaming output [\#36](https://github.com/asteris-llc/converge/pull/36) ([BrianHicks](https://github.com/BrianHicks))
- Apply [\#32](https://github.com/asteris-llc/converge/pull/32) ([BrianHicks](https://github.com/BrianHicks))
- load: accept initial parameters [\#28](https://github.com/asteris-llc/converge/pull/28) ([BrianHicks](https://github.com/BrianHicks))
- Add tests for Plan [\#25](https://github.com/asteris-llc/converge/pull/25) ([BrianHicks](https://github.com/BrianHicks))
- plan: pretty-printing terminal output [\#22](https://github.com/asteris-llc/converge/pull/22) ([siddharthist](https://github.com/siddharthist))
- fix incomplete comment [\#19](https://github.com/asteris-llc/converge/pull/19) ([BrianHicks](https://github.com/BrianHicks))
- Checker [\#9](https://github.com/asteris-llc/converge/pull/9) ([BrianHicks](https://github.com/BrianHicks))
- Name\(\) -\> String\(\) [\#8](https://github.com/asteris-llc/converge/pull/8) ([siddharthist](https://github.com/siddharthist))
- \[WIP\] Graphviz Support [\#7](https://github.com/asteris-llc/converge/pull/7) ([BrianHicks](https://github.com/BrianHicks))
- Feature/validation error type [\#3](https://github.com/asteris-llc/converge/pull/3) ([siddharthist](https://github.com/siddharthist))
- shelltask: add 'validate' method [\#2](https://github.com/asteris-llc/converge/pull/2) ([siddharthist](https://github.com/siddharthist))
- readme: status -\> check [\#1](https://github.com/asteris-llc/converge/pull/1) ([siddharthist](https://github.com/siddharthist))

# Change Log

## [Unreleased](https://github.com/asteris-llc/converge/tree/HEAD)

[Full Changelog](https://github.com/asteris-llc/converge/compare/0.5.0-rc1...HEAD)

### Closed Issues

- do not output private fields in docs [\#552](https://github.com/asteris-llc/converge/issues/552)

### Closed Pull Requests

- Highlight HCL [\#573](https://github.com/asteris-llc/converge/pull/573) ([BrianHicks](https://github.com/BrianHicks))
- Update docs style with new branding [\#572](https://github.com/asteris-llc/converge/pull/572) ([BrianHicks](https://github.com/BrianHicks))

## [0.5.0-rc1](https://github.com/asteris-llc/converge/tree/0.5.0-rc1) (2016-12-28)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.5.0-beta1...0.5.0-rc1)

### Closed Issues

- task.query behavior [\#567](https://github.com/asteris-llc/converge/issues/567)
- Get rid of vendor directory [\#545](https://github.com/asteris-llc/converge/issues/545)

### Closed Pull Requests

- no merge: update changelog [\#569](https://github.com/asteris-llc/converge/pull/569) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docs: only list exported fields in param and export sections [\#561](https://github.com/asteris-llc/converge/pull/561) ([arichardet](https://github.com/arichardet))
- Fix Linting Bugs [\#554](https://github.com/asteris-llc/converge/pull/554) ([BrianHicks](https://github.com/BrianHicks))

## [0.5.0-beta1](https://github.com/asteris-llc/converge/tree/0.5.0-beta1) (2016-12-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0...0.5.0-beta1)

### Enhancements

- Handle time.Time in preparer [\#490](https://github.com/asteris-llc/converge/issues/490)
- k8s demo [\#480](https://github.com/asteris-llc/converge/issues/480)
- no mo' spaces [\#320](https://github.com/asteris-llc/converge/issues/320)
- We're using too many ports [\#286](https://github.com/asteris-llc/converge/issues/286)
- better status display on plan and execution [\#269](https://github.com/asteris-llc/converge/issues/269)
- builtin file owner module [\#74](https://github.com/asteris-llc/converge/issues/74)
- Add interceptors for RPC auth [\#531](https://github.com/asteris-llc/converge/pull/531) ([BrianHicks](https://github.com/BrianHicks))
- demo: kubernetes - use latest version of converge [\#524](https://github.com/asteris-llc/converge/pull/524) ([ryane](https://github.com/ryane))
- fix shasum filename [\#522](https://github.com/asteris-llc/converge/pull/522) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/improve output [\#516](https://github.com/asteris-llc/converge/pull/516) ([arichardet](https://github.com/arichardet))

### Bugs

- LVM FS module reports failure due to still having changes [\#555](https://github.com/asteris-llc/converge/issues/555)
- handle newlines in docs generation for exported fields [\#548](https://github.com/asteris-llc/converge/issues/548)
- docker.container panics if image not present [\#538](https://github.com/asteris-llc/converge/issues/538)
- Return error instead of message from docker resource plan [\#529](https://github.com/asteris-llc/converge/issues/529)
- misleading error message in converge graph command [\#520](https://github.com/asteris-llc/converge/issues/520)
- exit code is 0 when converge plan has errors [\#492](https://github.com/asteris-llc/converge/issues/492)
- error specifying file mode [\#487](https://github.com/asteris-llc/converge/issues/487)
- docker.network misleading error messages [\#485](https://github.com/asteris-llc/converge/issues/485)
- inconsistent indentation on multi-line diffs [\#478](https://github.com/asteris-llc/converge/issues/478)
- pass typed params through to modules [\#409](https://github.com/asteris-llc/converge/issues/409)
- auth code is repeated [\#273](https://github.com/asteris-llc/converge/issues/273)
- Remove indent for multi-line diffs [\#530](https://github.com/asteris-llc/converge/pull/530) ([arichardet](https://github.com/arichardet))
- status: add MayChange status and Warning display [\#528](https://github.com/asteris-llc/converge/pull/528) ([ryane](https://github.com/ryane))
- Fix/disable docker solaris [\#527](https://github.com/asteris-llc/converge/pull/527) ([ryane](https://github.com/ryane))
- plan cmd: return non-zero exit code when plan has error\(s\) [\#525](https://github.com/asteris-llc/converge/pull/525) ([ryane](https://github.com/ryane))
- graph cmd: fix wrong argument length error [\#521](https://github.com/asteris-llc/converge/pull/521) ([ryane](https://github.com/ryane))
- Fix/file.mode errors [\#519](https://github.com/asteris-llc/converge/pull/519) ([ryane](https://github.com/ryane))
- docker: update docker dependency [\#512](https://github.com/asteris-llc/converge/pull/512) ([ryane](https://github.com/ryane))

### Closed Issues

- Better output for cascading failures [\#367](https://github.com/asteris-llc/converge/issues/367)
- file download/fetch module [\#250](https://github.com/asteris-llc/converge/issues/250)

### Closed Pull Requests

- changelog branch; tag and push tags before merging [\#559](https://github.com/asteris-llc/converge/pull/559) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Minor Fixes After Makefile rewrite [\#558](https://github.com/asteris-llc/converge/pull/558) ([BrianHicks](https://github.com/BrianHicks))
- reset updates needed after apply in lvm [\#557](https://github.com/asteris-llc/converge/pull/557) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Makefile: remove ref to VERSION in PACKAGE\_VERSION [\#556](https://github.com/asteris-llc/converge/pull/556) ([BrianHicks](https://github.com/BrianHicks))
- Better Makefiles [\#553](https://github.com/asteris-llc/converge/pull/553) ([BrianHicks](https://github.com/BrianHicks))
- when apply fails due to changes still being present; use those diffs … [\#551](https://github.com/asteris-llc/converge/pull/551) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docs: use same formatting for \(re\)exported fields as parameters [\#550](https://github.com/asteris-llc/converge/pull/550) ([arichardet](https://github.com/arichardet))
- Feature/74 file owner [\#549](https://github.com/asteris-llc/converge/pull/549) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/file.fetch [\#543](https://github.com/asteris-llc/converge/pull/543) ([arichardet](https://github.com/arichardet))
- Fix/538 docker container panic [\#542](https://github.com/asteris-llc/converge/pull/542) ([rebeccaskinner](https://github.com/rebeccaskinner))
- return a better error for planned network failures [\#541](https://github.com/asteris-llc/converge/pull/541) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Timer status display [\#537](https://github.com/asteris-llc/converge/pull/537) ([BrianHicks](https://github.com/BrianHicks))
- Add fuzzing and benchmarks to CI [\#536](https://github.com/asteris-llc/converge/pull/536) ([BrianHicks](https://github.com/BrianHicks))
- disallow some characters as resource names [\#535](https://github.com/asteris-llc/converge/pull/535) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use a single port [\#534](https://github.com/asteris-llc/converge/pull/534) ([BrianHicks](https://github.com/BrianHicks))
- Fix/374 immutable status [\#533](https://github.com/asteris-llc/converge/pull/533) ([rebeccaskinner](https://github.com/rebeccaskinner))
- add logo to Readme [\#532](https://github.com/asteris-llc/converge/pull/532) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docs [\#526](https://github.com/asteris-llc/converge/pull/526) ([arichardet](https://github.com/arichardet))
- Feature/preparer time.time [\#523](https://github.com/asteris-llc/converge/pull/523) ([arichardet](https://github.com/arichardet))
- update readme with graph and download info [\#518](https://github.com/asteris-llc/converge/pull/518) ([stevendborrelli](https://github.com/stevendborrelli))
- kubernetes example [\#507](https://github.com/asteris-llc/converge/pull/507) ([ryane](https://github.com/ryane))

## [0.4.0](https://github.com/asteris-llc/converge/tree/0.4.0) (2016-11-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-rc1...0.4.0)

### Enhancements

- Add `nonempty` tag to preparer [\#486](https://github.com/asteris-llc/converge/issues/486)
- apt module [\#404](https://github.com/asteris-llc/converge/issues/404)
- Handle time.Duration in Preparer [\#358](https://github.com/asteris-llc/converge/issues/358)
- binary distribution [\#306](https://github.com/asteris-llc/converge/issues/306)
- docker.volume resource [\#299](https://github.com/asteris-llc/converge/issues/299)
- docker.network resource [\#298](https://github.com/asteris-llc/converge/issues/298)
- Add ability to modify user [\#278](https://github.com/asteris-llc/converge/issues/278)
- LVM module [\#270](https://github.com/asteris-llc/converge/issues/270)
- update installer to point to 0.4.0 [\#517](https://github.com/asteris-llc/converge/pull/517) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/preparer time.duration [\#483](https://github.com/asteris-llc/converge/pull/483) ([arichardet](https://github.com/arichardet))
- check rpm for empty packages [\#482](https://github.com/asteris-llc/converge/pull/482) ([stevendborrelli](https://github.com/stevendborrelli))
- only extract binary from tarball [\#481](https://github.com/asteris-llc/converge/pull/481) ([stevendborrelli](https://github.com/stevendborrelli))
- docker.network resource [\#477](https://github.com/asteris-llc/converge/pull/477) ([ryane](https://github.com/ryane))
- Feature/apt package [\#461](https://github.com/asteris-llc/converge/pull/461) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docker.volume [\#453](https://github.com/asteris-llc/converge/pull/453) ([ryane](https://github.com/ryane))
- Add the ability to modify a user [\#434](https://github.com/asteris-llc/converge/pull/434) ([arichardet](https://github.com/arichardet))

### Bugs

- validate non-empty string values on some resources [\#337](https://github.com/asteris-llc/converge/issues/337)
- Intermittent test failure with `make test` [\#476](https://github.com/asteris-llc/converge/issues/476)
- Converge loses dependencies between modules within a branch when  [\#427](https://github.com/asteris-llc/converge/issues/427)
- diff output is partially bold [\#402](https://github.com/asteris-llc/converge/issues/402)
- Apply doesn't show diff output [\#399](https://github.com/asteris-llc/converge/issues/399)
- values that don't implement Stringer log badly [\#398](https://github.com/asteris-llc/converge/issues/398)
- resolving conditional macros line [\#389](https://github.com/asteris-llc/converge/issues/389)
- Case predicates fail when they include `lookup` [\#386](https://github.com/asteris-llc/converge/issues/386)
- converge exit code on a failed run [\#365](https://github.com/asteris-llc/converge/issues/365)
- Address name conflicts [\#361](https://github.com/asteris-llc/converge/issues/361)
- resources should have the ability to gracefully stop on interrupt [\#319](https://github.com/asteris-llc/converge/issues/319)
- converge panics when interrupted [\#318](https://github.com/asteris-llc/converge/issues/318)
- Fix/conditional nodes [\#484](https://github.com/asteris-llc/converge/pull/484) ([rebeccaskinner](https://github.com/rebeccaskinner))
- grouping: fix issues with groups and conditionals [\#460](https://github.com/asteris-llc/converge/pull/460) ([ryane](https://github.com/ryane))
- return non-zero exit code if apply graph has an error [\#436](https://github.com/asteris-llc/converge/pull/436) ([ryane](https://github.com/ryane))
- replace "context" with "golang.org/x/net/context" [\#433](https://github.com/asteris-llc/converge/pull/433) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- update changelog, add script [\#511](https://github.com/asteris-llc/converge/pull/511) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: fix link in resource authors guide [\#479](https://github.com/asteris-llc/converge/pull/479) ([ryane](https://github.com/ryane))
- Fix/resource duplication [\#458](https://github.com/asteris-llc/converge/pull/458) ([arichardet](https://github.com/arichardet))
- Fix/graceful exit [\#447](https://github.com/asteris-llc/converge/pull/447) ([arichardet](https://github.com/arichardet))
- LVM module [\#184](https://github.com/asteris-llc/converge/pull/184) ([avnik](https://github.com/avnik))

## [0.4.0-rc1](https://github.com/asteris-llc/converge/tree/0.4.0-rc1) (2016-11-17)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-beta1...0.4.0-rc1)

### Closed Pull Requests

- update changelog to reflect rc1 status [\#510](https://github.com/asteris-llc/converge/pull/510) ([rebeccaskinner](https://github.com/rebeccaskinner))
- move rendering plant creation inside of the transform block to avoid … [\#509](https://github.com/asteris-llc/converge/pull/509) ([rebeccaskinner](https://github.com/rebeccaskinner))
- set github link to the converge repo [\#508](https://github.com/asteris-llc/converge/pull/508) ([rebeccaskinner](https://github.com/rebeccaskinner))
- empty commit for CI [\#506](https://github.com/asteris-llc/converge/pull/506) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.4.0-beta1](https://github.com/asteris-llc/converge/tree/0.4.0-beta1) (2016-11-15)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0...0.4.0-beta1)

### Enhancements

- file.owner resource [\#503](https://github.com/asteris-llc/converge/issues/503)

### Bugs

- package.rpm state does not default to present [\#446](https://github.com/asteris-llc/converge/issues/446)
- Error when adding a user when a group of the same name exists [\#425](https://github.com/asteris-llc/converge/issues/425)

### Closed Issues

- nonempty check in preparer for user/group [\#489](https://github.com/asteris-llc/converge/issues/489)

### Closed Pull Requests

- Fix/0.4.0 beta1 fix [\#505](https://github.com/asteris-llc/converge/pull/505) ([rebeccaskinner](https://github.com/rebeccaskinner))
- release notes [\#497](https://github.com/asteris-llc/converge/pull/497) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use %v for logging [\#496](https://github.com/asteris-llc/converge/pull/496) ([arichardet](https://github.com/arichardet))
- update docs with conditional changes [\#495](https://github.com/asteris-llc/converge/pull/495) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/nonempty tag [\#494](https://github.com/asteris-llc/converge/pull/494) ([arichardet](https://github.com/arichardet))
- Remove extract executable [\#488](https://github.com/asteris-llc/converge/pull/488) ([arichardet](https://github.com/arichardet))
- Manifest / index page [\#464](https://github.com/asteris-llc/converge/pull/464) ([BrianHicks](https://github.com/BrianHicks))
- default to state 'present' for rpm module [\#463](https://github.com/asteris-llc/converge/pull/463) ([rebeccaskinner](https://github.com/rebeccaskinner))
- vendor: upgrade all deps [\#459](https://github.com/asteris-llc/converge/pull/459) ([BrianHicks](https://github.com/BrianHicks))

## [0.3.0](https://github.com/asteris-llc/converge/tree/0.3.0) (2016-10-27)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-rc1...0.3.0)

### Enhancements

- "param is required" error should include param name [\#335](https://github.com/asteris-llc/converge/issues/335)
- Add ability to modify group [\#279](https://github.com/asteris-llc/converge/issues/279)
- add conditionals to module resource [\#268](https://github.com/asteris-llc/converge/issues/268)
- named locks [\#249](https://github.com/asteris-llc/converge/issues/249)
- pretty printed changes should align values [\#244](https://github.com/asteris-llc/converge/issues/244)
- ability to wait for a condition [\#238](https://github.com/asteris-llc/converge/issues/238)
- proper arrays [\#110](https://github.com/asteris-llc/converge/issues/110)
- makefile: add hashes to files [\#440](https://github.com/asteris-llc/converge/pull/440) ([BrianHicks](https://github.com/BrianHicks))
- \[389\] switch log level from Info to Debug [\#428](https://github.com/asteris-llc/converge/pull/428) ([mason-fish](https://github.com/mason-fish))
- Feature/examples 0.3.0 [\#419](https://github.com/asteris-llc/converge/pull/419) ([ryane](https://github.com/ryane))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- named groups [\#392](https://github.com/asteris-llc/converge/pull/392) ([ryane](https://github.com/ryane))
- Don't render module dependencies [\#387](https://github.com/asteris-llc/converge/pull/387) ([BrianHicks](https://github.com/BrianHicks))
- Added docs, renamed packages to package to  for consistency with other modules [\#384](https://github.com/asteris-llc/converge/pull/384) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use status.RaiseLevel in group [\#377](https://github.com/asteris-llc/converge/pull/377) ([arichardet](https://github.com/arichardet))
- Feature/rpm module [\#373](https://github.com/asteris-llc/converge/pull/373) ([rebeccaskinner](https://github.com/rebeccaskinner))
- use a container for CI [\#372](https://github.com/asteris-llc/converge/pull/372) ([BrianHicks](https://github.com/BrianHicks))
- create a metadata envelope for nodes [\#369](https://github.com/asteris-llc/converge/pull/369) ([BrianHicks](https://github.com/BrianHicks))
- cmd/server.go: use errgroup instead of waitgroup [\#363](https://github.com/asteris-llc/converge/pull/363) ([QuentinPerez](https://github.com/QuentinPerez))
- Feature/conditionals [\#362](https://github.com/asteris-llc/converge/pull/362) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/example swarm wait [\#346](https://github.com/asteris-llc/converge/pull/346) ([ryane](https://github.com/ryane))
- Add Compound Parameters [\#340](https://github.com/asteris-llc/converge/pull/340) ([BrianHicks](https://github.com/BrianHicks))
- load overlay module in elk example [\#326](https://github.com/asteris-llc/converge/pull/326) ([feniix](https://github.com/feniix))
- Refactor deserializers [\#321](https://github.com/asteris-llc/converge/pull/321) ([BrianHicks](https://github.com/BrianHicks))
- Use text/tabwriter to align human output [\#317](https://github.com/asteris-llc/converge/pull/317) ([sehqlr](https://github.com/sehqlr))
- Fix/225 pipeline function refactor [\#307](https://github.com/asteris-llc/converge/pull/307) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Bugs

- Looking up a node from a branch that is also a part of a group introduced a deadlock [\#415](https://github.com/asteris-llc/converge/issues/415)
- user/group not activating changes [\#417](https://github.com/asteris-llc/converge/issues/417)
- docker.container reporting incorrect entrypoint diff [\#410](https://github.com/asteris-llc/converge/issues/410)
- converge panics when encountering an empty conditional [\#407](https://github.com/asteris-llc/converge/issues/407)
- Use lists as params in modules is broken [\#397](https://github.com/asteris-llc/converge/issues/397)
- Explicit dependencies fail inside of case statements [\#385](https://github.com/asteris-llc/converge/issues/385)
- Use `package.rpm` not `rpm.package` for rpm module [\#382](https://github.com/asteris-llc/converge/issues/382)
- shell module should not set StatusLevel based on process exit code [\#323](https://github.com/asteris-llc/converge/issues/323)
- StatusLevel not taken into account during graph execution [\#322](https://github.com/asteris-llc/converge/issues/322)
- param dependency fails in `samples/shellContext.hcl` [\#313](https://github.com/asteris-llc/converge/issues/313)
- Execution Engine Ignoring Warning Levels [\#243](https://github.com/asteris-llc/converge/issues/243)
- Make pipeline functions use mult-return instead of Either [\#225](https://github.com/asteris-llc/converge/issues/225)
- Fix 420 go test race [\#424](https://github.com/asteris-llc/converge/pull/424) ([ryane](https://github.com/ryane))
- Replace `AreSiblingIDs` with `AreSiblings` to prevent deadlocking [\#423](https://github.com/asteris-llc/converge/pull/423) ([rebeccaskinner](https://github.com/rebeccaskinner))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- Handle thunked branches [\#403](https://github.com/asteris-llc/converge/pull/403) ([rebeccaskinner](https://github.com/rebeccaskinner))
- don't exclude modules in getNearestAncestor [\#396](https://github.com/asteris-llc/converge/pull/396) ([ryane](https://github.com/ryane))
- Fix/385 dependencies in conditionals [\#391](https://github.com/asteris-llc/converge/pull/391) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix lookup calls to use os/user in SetAddUserOptions [\#390](https://github.com/asteris-llc/converge/pull/390) ([arichardet](https://github.com/arichardet))
- Change `rpm.package` to `package.rpm` [\#383](https://github.com/asteris-llc/converge/pull/383) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Update user status level and errors, add checks in Apply for group [\#375](https://github.com/asteris-llc/converge/pull/375) ([arichardet](https://github.com/arichardet))
- Ability to use pointers as a preparer value [\#339](https://github.com/asteris-llc/converge/pull/339) ([BrianHicks](https://github.com/BrianHicks))
- Status Error Codes [\#333](https://github.com/asteris-llc/converge/pull/333) ([BrianHicks](https://github.com/BrianHicks))
- fix codeclimate yaml [\#328](https://github.com/asteris-llc/converge/pull/328) ([BrianHicks](https://github.com/BrianHicks))

### Closed Issues

- docs for rpm module [\#381](https://github.com/asteris-llc/converge/issues/381)

### Closed Pull Requests

- update for 0.3.0 [\#443](https://github.com/asteris-llc/converge/pull/443) ([stevendborrelli](https://github.com/stevendborrelli))
- prep for 0.3.0-rc1 [\#437](https://github.com/asteris-llc/converge/pull/437) ([stevendborrelli](https://github.com/stevendborrelli))
- update release notes [\#435](https://github.com/asteris-llc/converge/pull/435) ([stevendborrelli](https://github.com/stevendborrelli))
- beta3 release notes [\#431](https://github.com/asteris-llc/converge/pull/431) ([stevendborrelli](https://github.com/stevendborrelli))
- update docs to include currently identified switch statement limitations [\#388](https://github.com/asteris-llc/converge/pull/388) ([rebeccaskinner](https://github.com/rebeccaskinner))
- update vendored dependencies [\#378](https://github.com/asteris-llc/converge/pull/378) ([BrianHicks](https://github.com/BrianHicks))
- 0.3.0 release documentation [\#376](https://github.com/asteris-llc/converge/pull/376) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: add multi-version control [\#371](https://github.com/asteris-llc/converge/pull/371) ([BrianHicks](https://github.com/BrianHicks))
- add docs for the platform module [\#342](https://github.com/asteris-llc/converge/pull/342) ([stevendborrelli](https://github.com/stevendborrelli))
- update for PR \#307 [\#316](https://github.com/asteris-llc/converge/pull/316) ([stevendborrelli](https://github.com/stevendborrelli))

## [0.3.0-rc1](https://github.com/asteris-llc/converge/tree/0.3.0-rc1) (2016-10-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta3...0.3.0-rc1)

### Closed Issues

- Update documentation regarding conditionals and groups [\#430](https://github.com/asteris-llc/converge/issues/430)

## [0.3.0-beta3](https://github.com/asteris-llc/converge/tree/0.3.0-beta3) (2016-10-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta2...0.3.0-beta3)

### Enhancements

- add installation script [\#416](https://github.com/asteris-llc/converge/pull/416) ([stevendborrelli](https://github.com/stevendborrelli))

### Bugs

- go tests race condition [\#420](https://github.com/asteris-llc/converge/issues/420)

### Closed Pull Requests

- document known bug between conditionals and groups [\#432](https://github.com/asteris-llc/converge/pull/432) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.3.0-beta2](https://github.com/asteris-llc/converge/tree/0.3.0-beta2) (2016-10-24)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta1...0.3.0-beta2)

### Bugs

- Fix container entrypoint diffs [\#412](https://github.com/asteris-llc/converge/pull/412) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- wercker: don't delete removed [\#414](https://github.com/asteris-llc/converge/pull/414) ([BrianHicks](https://github.com/BrianHicks))
- Nightly/test builds [\#413](https://github.com/asteris-llc/converge/pull/413) ([BrianHicks](https://github.com/BrianHicks))
- fix panic [\#408](https://github.com/asteris-llc/converge/pull/408) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.3.0-beta1](https://github.com/asteris-llc/converge/tree/0.3.0-beta1) (2016-10-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/before-moving-resolve-conditionals...0.3.0-beta1)

### Enhancements

- file/dir refactor [\#327](https://github.com/asteris-llc/converge/issues/327)
- Use errgroup in server [\#295](https://github.com/asteris-llc/converge/issues/295)
- Allow ability to indicate group name when adding a user [\#276](https://github.com/asteris-llc/converge/issues/276)
- builtin file group module [\#75](https://github.com/asteris-llc/converge/issues/75)

### Bugs

- Conditional Regression [\#401](https://github.com/asteris-llc/converge/issues/401)
- module dependencies now failing [\#395](https://github.com/asteris-llc/converge/issues/395)
- docker.container regression [\#343](https://github.com/asteris-llc/converge/issues/343)
- samples in the README are formatted incorrectly. [\#104](https://github.com/asteris-llc/converge/issues/104)

### Closed Issues

- Tidy up codebase by fixing `make lint` reports [\#348](https://github.com/asteris-llc/converge/issues/348)
- Can we simplify the visualization of large modules? [\#280](https://github.com/asteris-llc/converge/issues/280)
- How should we manage module docs and examples? [\#173](https://github.com/asteris-llc/converge/issues/173)

### Closed Pull Requests

- Allow lists in parameters when called in a module [\#400](https://github.com/asteris-llc/converge/pull/400) ([BrianHicks](https://github.com/BrianHicks))
- comment typo [\#380](https://github.com/asteris-llc/converge/pull/380) ([rebeccaskinner](https://github.com/rebeccaskinner))
- FIX local command flag added to example [\#370](https://github.com/asteris-llc/converge/pull/370) ([philcryer](https://github.com/philcryer))
- codeclimate: remove markdown linting [\#368](https://github.com/asteris-llc/converge/pull/368) ([BrianHicks](https://github.com/BrianHicks))
- Add testing section to contribution guidelines [\#366](https://github.com/asteris-llc/converge/pull/366) ([arichardet](https://github.com/arichardet))
- remove binary [\#354](https://github.com/asteris-llc/converge/pull/354) ([stevendborrelli](https://github.com/stevendborrelli))
- wait: fix error handling in retrier [\#353](https://github.com/asteris-llc/converge/pull/353) ([ryane](https://github.com/ryane))
- Feature/param name in error [\#352](https://github.com/asteris-llc/converge/pull/352) ([ryane](https://github.com/ryane))
- docs: add section on error values to resource author's guide [\#347](https://github.com/asteris-llc/converge/pull/347) ([BrianHicks](https://github.com/BrianHicks))
- Fix docker regressions [\#345](https://github.com/asteris-llc/converge/pull/345) ([ryane](https://github.com/ryane))
- wait.query and wait.port resources [\#334](https://github.com/asteris-llc/converge/pull/334) ([ryane](https://github.com/ryane))

## [before-moving-resolve-conditionals](https://github.com/asteris-llc/converge/tree/before-moving-resolve-conditionals) (2016-10-05)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0...before-moving-resolve-conditionals)

### Enhancements

- Use github.com/pkg/errors exclusively [\#300](https://github.com/asteris-llc/converge/issues/300)

### Bugs

- handle parameters with valid zero value [\#338](https://github.com/asteris-llc/converge/issues/338)

### Closed Issues

- Universal Benchmarking [\#308](https://github.com/asteris-llc/converge/issues/308)
- Use table tests instead of helpers.PreparerValidator [\#155](https://github.com/asteris-llc/converge/issues/155)

### Closed Pull Requests

- Handle preparer fields where zero is valid [\#344](https://github.com/asteris-llc/converge/pull/344) ([arichardet](https://github.com/arichardet))
- plan: test, which illustrate "return nil, err" problem [\#336](https://github.com/asteris-llc/converge/pull/336) ([avnik](https://github.com/avnik))
- Support for modifying linux groups [\#331](https://github.com/asteris-llc/converge/pull/331) ([arichardet](https://github.com/arichardet))
- Add contributing.md and code of conduct [\#330](https://github.com/asteris-llc/converge/pull/330) ([BrianHicks](https://github.com/BrianHicks))
- shell: non-0 exit code returns StatusWillChange [\#324](https://github.com/asteris-llc/converge/pull/324) ([ryane](https://github.com/ryane))
- map keys are considered "strings" in parse.Node [\#315](https://github.com/asteris-llc/converge/pull/315) ([rebeccaskinner](https://github.com/rebeccaskinner))
- helpers,docker: move AssertDiff to common testhelpers [\#312](https://github.com/asteris-llc/converge/pull/312) ([avnik](https://github.com/avnik))

## [0.2.0](https://github.com/asteris-llc/converge/tree/0.2.0) (2016-09-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-rc1...0.2.0)

### Enhancements

- Calls to `lookup` should be all lower-case [\#223](https://github.com/asteris-llc/converge/issues/223)

### Bugs

- field name conflicts during lookups [\#294](https://github.com/asteris-llc/converge/issues/294)
- panic when trying to lookup against file.directory [\#292](https://github.com/asteris-llc/converge/issues/292)
- Allow adding and deleting a group without specifying gid [\#275](https://github.com/asteris-llc/converge/issues/275)
- param lookup race condition [\#266](https://github.com/asteris-llc/converge/issues/266)
- Node is sometimes unresolvable, depending on ordering within a template [\#254](https://github.com/asteris-llc/converge/issues/254)
- params should handle boolean values [\#251](https://github.com/asteris-llc/converge/issues/251)
- remove token handling logic when only a client and not a server [\#247](https://github.com/asteris-llc/converge/issues/247)
- Check shouldn't be called from Apply [\#240](https://github.com/asteris-llc/converge/issues/240)
- value passing causes dependency resolution failures in child modules [\#239](https://github.com/asteris-llc/converge/issues/239)
- task.Task has extra argument [\#237](https://github.com/asteris-llc/converge/issues/237)
- Lookups fail for file.content [\#236](https://github.com/asteris-llc/converge/issues/236)
- RPC isn't used in healthchecks [\#199](https://github.com/asteris-llc/converge/issues/199)

### Closed Issues

- module author guide [\#291](https://github.com/asteris-llc/converge/issues/291)
- logging is a little too chatty [\#248](https://github.com/asteris-llc/converge/issues/248)
- Add RPC auth documentation [\#246](https://github.com/asteris-llc/converge/issues/246)

### Closed Pull Requests

- Add ability to indicate the group name when adding a user [\#310](https://github.com/asteris-llc/converge/pull/310) ([arichardet](https://github.com/arichardet))
- 0.2.0 release [\#311](https://github.com/asteris-llc/converge/pull/311) ([stevendborrelli](https://github.com/stevendborrelli))
- Add "Basic Usage" section to Server docs [\#287](https://github.com/asteris-llc/converge/pull/287) ([sehqlr](https://github.com/sehqlr))
- Simplify Status implementation [\#284](https://github.com/asteris-llc/converge/pull/284) ([BrianHicks](https://github.com/BrianHicks))
- Fix climate issues [\#282](https://github.com/asteris-llc/converge/pull/282) ([BrianHicks](https://github.com/BrianHicks))
- Add codeclimate checks [\#281](https://github.com/asteris-llc/converge/pull/281) ([BrianHicks](https://github.com/BrianHicks))
- Make the logs quieter [\#274](https://github.com/asteris-llc/converge/pull/274) ([BrianHicks](https://github.com/BrianHicks))
- example: elasticsearch, kibana, and filebeat ~elk [\#272](https://github.com/asteris-llc/converge/pull/272) ([ryane](https://github.com/ryane))
- param: accept boolean values [\#271](https://github.com/asteris-llc/converge/pull/271) ([BrianHicks](https://github.com/BrianHicks))
- Use `MakeLanguage` for dependecy resolution to ensure that no [\#265](https://github.com/asteris-llc/converge/pull/265) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Make notes match sample code [\#262](https://github.com/asteris-llc/converge/pull/262) ([tomduckering](https://github.com/tomduckering))

## [0.2.0-rc1](https://github.com/asteris-llc/converge/tree/0.2.0-rc1) (2016-09-23)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-beta2...0.2.0-rc1)

### Bugs

- docker.image incorrectly reporting changes after apply [\#304](https://github.com/asteris-llc/converge/issues/304)

### Closed Issues

- preprocessor race condition [\#302](https://github.com/asteris-llc/converge/issues/302)

### Closed Pull Requests

- Fixes docker resources and updates examples [\#305](https://github.com/asteris-llc/converge/pull/305) ([ryane](https://github.com/ryane))
- use thread-safe field cache [\#303](https://github.com/asteris-llc/converge/pull/303) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/294 field name conflicts [\#301](https://github.com/asteris-llc/converge/pull/301) ([rebeccaskinner](https://github.com/rebeccaskinner))
- preproc: don't add fields for non-struct anon field [\#293](https://github.com/asteris-llc/converge/pull/293) ([ryane](https://github.com/ryane))
- docs: add draft resource authors guide [\#290](https://github.com/asteris-llc/converge/pull/290) ([BrianHicks](https://github.com/BrianHicks))
- perform healthchecks over RPC [\#289](https://github.com/asteris-llc/converge/pull/289) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/239 [\#288](https://github.com/asteris-llc/converge/pull/288) ([rebeccaskinner](https://github.com/rebeccaskinner))
- cmd: don't set local token if only a client [\#285](https://github.com/asteris-llc/converge/pull/285) ([BrianHicks](https://github.com/BrianHicks))
- Fix/user group - Allow adding/deleting without gid [\#283](https://github.com/asteris-llc/converge/pull/283) ([arichardet](https://github.com/arichardet))

## [0.2.0-beta2](https://github.com/asteris-llc/converge/tree/0.2.0-beta2) (2016-09-20)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-beta1...0.2.0-beta2)

### Enhancements

- module signing [\#16](https://github.com/asteris-llc/converge/issues/16)

### Bugs

- This graph is rendered kinda funky [\#141](https://github.com/asteris-llc/converge/issues/141)

### Closed Issues

- Update all deps to the latest versions [\#171](https://github.com/asteris-llc/converge/issues/171)
- human printer should be profiled and optimized [\#156](https://github.com/asteris-llc/converge/issues/156)

### Closed Pull Requests

- example: docker swarm mode [\#267](https://github.com/asteris-llc/converge/pull/267) ([ryane](https://github.com/ryane))

## [0.2.0-beta1](https://github.com/asteris-llc/converge/tree/0.2.0-beta1) (2016-09-16)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.1.1...0.2.0-beta1)

### Enhancements

- Create a query module [\#227](https://github.com/asteris-llc/converge/issues/227)
- Graphing should work over RPC [\#207](https://github.com/asteris-llc/converge/issues/207)
- Allow Value-passing between modules [\#206](https://github.com/asteris-llc/converge/issues/206)
- logging needs improvement [\#200](https://github.com/asteris-llc/converge/issues/200)
- REST layer on RPC [\#170](https://github.com/asteris-llc/converge/issues/170)
- CLI RPC client [\#169](https://github.com/asteris-llc/converge/issues/169)
- RPC authentication [\#168](https://github.com/asteris-llc/converge/issues/168)
- RPC server [\#167](https://github.com/asteris-llc/converge/issues/167)
- Docker module [\#159](https://github.com/asteris-llc/converge/issues/159)
- Clear/readable errors [\#140](https://github.com/asteris-llc/converge/issues/140)
- Macros? [\#111](https://github.com/asteris-llc/converge/issues/111)
- should template be called file.content? [\#103](https://github.com/asteris-llc/converge/issues/103)
- execution in new graph code [\#92](https://github.com/asteris-llc/converge/issues/92)
- execution planning in new graph code [\#91](https://github.com/asteris-llc/converge/issues/91)
- implement file.group [\#90](https://github.com/asteris-llc/converge/issues/90)
- reimplement file.user [\#89](https://github.com/asteris-llc/converge/issues/89)
- reimplement file.mode [\#88](https://github.com/asteris-llc/converge/issues/88)
- reimplement templates [\#87](https://github.com/asteris-llc/converge/issues/87)
- reimplement shell tasks [\#86](https://github.com/asteris-llc/converge/issues/86)
- implement graph visualization on top of new graph code [\#85](https://github.com/asteris-llc/converge/issues/85)
- resources should keep track of their parents [\#83](https://github.com/asteris-llc/converge/issues/83)
- infer dependencies based on template contents [\#65](https://github.com/asteris-llc/converge/issues/65)
- "native" filemode module [\#40](https://github.com/asteris-llc/converge/issues/40)
- "native" modules addressing scheme [\#39](https://github.com/asteris-llc/converge/issues/39)
- basic diff finding [\#20](https://github.com/asteris-llc/converge/issues/20)
- authenticated module server [\#18](https://github.com/asteris-llc/converge/issues/18)
- module server [\#17](https://github.com/asteris-llc/converge/issues/17)
- HTTP fetching [\#13](https://github.com/asteris-llc/converge/issues/13)
- graph: keep track of parentage inside edges [\#163](https://github.com/asteris-llc/converge/pull/163) ([BrianHicks](https://github.com/BrianHicks))

### Bugs

- lookup should be fully case insensitive [\#232](https://github.com/asteris-llc/converge/issues/232)
- rpc server does not support graph operation anymore [\#205](https://github.com/asteris-llc/converge/issues/205)
- graph is walked out of order [\#160](https://github.com/asteris-llc/converge/issues/160)
- Concurrent Map Read/Write Panic When Running Tests [\#158](https://github.com/asteris-llc/converge/issues/158)
- Shell Doesn't Validate Against User-Defined Interpreter [\#119](https://github.com/asteris-llc/converge/issues/119)
- invalid script syntax doesn't return a good error [\#113](https://github.com/asteris-llc/converge/issues/113)
- Blackbox/test\_apply\_remote.sh Fails [\#102](https://github.com/asteris-llc/converge/issues/102)
- `helpers` is kind of an awful module name [\#96](https://github.com/asteris-llc/converge/issues/96)
- Failing tasks don't block tasks that depend on them [\#82](https://github.com/asteris-llc/converge/issues/82)
- all modules in samples must be valid and useful [\#66](https://github.com/asteris-llc/converge/issues/66)
- params should not have default dependencies [\#62](https://github.com/asteris-llc/converge/issues/62)
- resource names with dashes break graph rendering [\#61](https://github.com/asteris-llc/converge/issues/61)
- Fix error handling when fail to print results during plan or apply [\#183](https://github.com/asteris-llc/converge/pull/183) ([arichardet](https://github.com/arichardet))
- Return better error information for Shell module failures [\#121](https://github.com/asteris-llc/converge/pull/121) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Closed Issues

- errors from check silently swallowed during apply [\#213](https://github.com/asteris-llc/converge/issues/213)
- errors are printed with full stacks [\#208](https://github.com/asteris-llc/converge/issues/208)
- nil pointer panic when file.content resource errors [\#179](https://github.com/asteris-llc/converge/issues/179)
- intermittent concurrent map write panic with dependent params [\#176](https://github.com/asteris-llc/converge/issues/176)
- use standard library net/context library [\#161](https://github.com/asteris-llc/converge/issues/161)
- Subgraph PrettyPrinter should use `nil` as the `SubgraphBottomID` [\#149](https://github.com/asteris-llc/converge/issues/149)
- DigraphPrettyPrinter \(prettyprinters\) should shrink [\#147](https://github.com/asteris-llc/converge/issues/147)
- Renderable \(prettyprinters\) is larger than necessary [\#146](https://github.com/asteris-llc/converge/issues/146)
- Shell command context [\#145](https://github.com/asteris-llc/converge/issues/145)
- Add a 'Health' Module [\#144](https://github.com/asteris-llc/converge/issues/144)
- Builtin file.directory module [\#137](https://github.com/asteris-llc/converge/issues/137)
- Graphviz Generation Fails With Empty Param [\#129](https://github.com/asteris-llc/converge/issues/129)
- parameters should be required if no default is set [\#123](https://github.com/asteris-llc/converge/issues/123)
- Shell Module Error Codes Could Have Better Formatting [\#120](https://github.com/asteris-llc/converge/issues/120)
- params should allow commas [\#114](https://github.com/asteris-llc/converge/issues/114)
- modules and params should not show up by default [\#106](https://github.com/asteris-llc/converge/issues/106)
- use a context in loading [\#94](https://github.com/asteris-llc/converge/issues/94)
- builtin file owner and file permissions module [\#84](https://github.com/asteris-llc/converge/issues/84)

### Closed Pull Requests

- We got an assigned port so here's docs. [\#261](https://github.com/asteris-llc/converge/pull/261) ([BrianHicks](https://github.com/BrianHicks))
- update readme [\#260](https://github.com/asteris-llc/converge/pull/260) ([stevendborrelli](https://github.com/stevendborrelli))
- Add user support [\#259](https://github.com/asteris-llc/converge/pull/259) ([arichardet](https://github.com/arichardet))
- MOAR DOCS PLS [\#258](https://github.com/asteris-llc/converge/pull/258) ([BrianHicks](https://github.com/BrianHicks))
- Makefile: refine xcompile targets [\#257](https://github.com/asteris-llc/converge/pull/257) ([BrianHicks](https://github.com/BrianHicks))
- Fixed issue where apply swallowed error. [\#256](https://github.com/asteris-llc/converge/pull/256) ([Dacode45](https://github.com/Dacode45))
- add mutex to render [\#255](https://github.com/asteris-llc/converge/pull/255) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Add default group implementation for unsupported systems [\#252](https://github.com/asteris-llc/converge/pull/252) ([arichardet](https://github.com/arichardet))
- Feature/module verification [\#245](https://github.com/asteris-llc/converge/pull/245) ([ajhager](https://github.com/ajhager))
- Feature/fully lowercase field names [\#242](https://github.com/asteris-llc/converge/pull/242) ([rebeccaskinner](https://github.com/rebeccaskinner))
- file.directory resource [\#241](https://github.com/asteris-llc/converge/pull/241) ([BrianHicks](https://github.com/BrianHicks))
- Add user group support [\#234](https://github.com/asteris-llc/converge/pull/234) ([arichardet](https://github.com/arichardet))
- support lowercase field names [\#231](https://github.com/asteris-llc/converge/pull/231) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/query refactor [\#230](https://github.com/asteris-llc/converge/pull/230) ([rebeccaskinner](https://github.com/rebeccaskinner))
- handle params that are thunked [\#229](https://github.com/asteris-llc/converge/pull/229) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/query module [\#228](https://github.com/asteris-llc/converge/pull/228) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Refactor helpers into descriptive names [\#222](https://github.com/asteris-llc/converge/pull/222) ([BrianHicks](https://github.com/BrianHicks))
- Feature/value passing [\#220](https://github.com/asteris-llc/converge/pull/220) ([rebeccaskinner](https://github.com/rebeccaskinner))
- rpc: flatten details field [\#218](https://github.com/asteris-llc/converge/pull/218) ([BrianHicks](https://github.com/BrianHicks))
- Feature/formatter \#208 [\#217](https://github.com/asteris-llc/converge/pull/217) ([BrianHicks](https://github.com/BrianHicks))
- docker.container: fix entrypoint diff logic [\#215](https://github.com/asteris-llc/converge/pull/215) ([ryane](https://github.com/ryane))
- Graphing over RPC [\#214](https://github.com/asteris-llc/converge/pull/214) ([BrianHicks](https://github.com/BrianHicks))
- docker.container resource [\#203](https://github.com/asteris-llc/converge/pull/203) ([ryane](https://github.com/ryane))
- Rework logs to use logrus [\#202](https://github.com/asteris-llc/converge/pull/202) ([BrianHicks](https://github.com/BrianHicks))
- Feature/platform refactor [\#197](https://github.com/asteris-llc/converge/pull/197) ([stevendborrelli](https://github.com/stevendborrelli))
- Added host environment variable support [\#195](https://github.com/asteris-llc/converge/pull/195) ([arichardet](https://github.com/arichardet))
- resources\(docker image\): document [\#193](https://github.com/asteris-llc/converge/pull/193) ([BrianHicks](https://github.com/BrianHicks))
- Documentation Site [\#192](https://github.com/asteris-llc/converge/pull/192) ([BrianHicks](https://github.com/BrianHicks))
- Feature/health module [\#189](https://github.com/asteris-llc/converge/pull/189) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docker.image resource redux [\#188](https://github.com/asteris-llc/converge/pull/188) ([ryane](https://github.com/ryane))
- Basic RPC [\#187](https://github.com/asteris-llc/converge/pull/187) ([BrianHicks](https://github.com/BrianHicks))
- shell env and working dir support [\#185](https://github.com/asteris-llc/converge/pull/185) ([ryane](https://github.com/ryane))
- \[GRAPH\] Fix \#158, add triggering files, add tests [\#181](https://github.com/asteris-llc/converge/pull/181) ([sehqlr](https://github.com/sehqlr))
- don't panic when result status is nil [\#180](https://github.com/asteris-llc/converge/pull/180) ([ryane](https://github.com/ryane))
- Feature/optimize human printer [\#175](https://github.com/asteris-llc/converge/pull/175) ([arichardet](https://github.com/arichardet))
- Shrink binaries with some magical ldflags [\#174](https://github.com/asteris-llc/converge/pull/174) ([BrianHicks](https://github.com/BrianHicks))
- Feature/shell refactor [\#172](https://github.com/asteris-llc/converge/pull/172) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Dependency ordered graph walking [\#166](https://github.com/asteris-llc/converge/pull/166) ([BrianHicks](https://github.com/BrianHicks))
- Migrate context for go1.7, refactor fetch/http.go [\#164](https://github.com/asteris-llc/converge/pull/164) ([sehqlr](https://github.com/sehqlr))
- Feature/resource return refactor [\#157](https://github.com/asteris-llc/converge/pull/157) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Move errors into the tree [\#154](https://github.com/asteris-llc/converge/pull/154) ([BrianHicks](https://github.com/BrianHicks))
- Human-readable pretty printer [\#152](https://github.com/asteris-llc/converge/pull/152) ([BrianHicks](https://github.com/BrianHicks))
- make nil a useful default bottom value for subgraphdi [\#150](https://github.com/asteris-llc/converge/pull/150) ([rebeccaskinner](https://github.com/rebeccaskinner))
- JSONL output and interface refactoring [\#148](https://github.com/asteris-llc/converge/pull/148) ([BrianHicks](https://github.com/BrianHicks))
- Feature/template split [\#136](https://github.com/asteris-llc/converge/pull/136) ([rebeccaskinner](https://github.com/rebeccaskinner))
- fetch: add context [\#135](https://github.com/asteris-llc/converge/pull/135) ([BrianHicks](https://github.com/BrianHicks))
- Render merged subtrees [\#133](https://github.com/asteris-llc/converge/pull/133) ([BrianHicks](https://github.com/BrianHicks))
- Fix/119 [\#132](https://github.com/asteris-llc/converge/pull/132) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Trim Duplicate Subtrees [\#131](https://github.com/asteris-llc/converge/pull/131) ([BrianHicks](https://github.com/BrianHicks))
- Appropriately handle required parameters instead of crashing [\#130](https://github.com/asteris-llc/converge/pull/130) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use concurrent map implementation for graph values [\#128](https://github.com/asteris-llc/converge/pull/128) ([BrianHicks](https://github.com/BrianHicks))
- Make params missing defaults required [\#127](https://github.com/asteris-llc/converge/pull/127) ([BrianHicks](https://github.com/BrianHicks))
- makefile: simplify linting [\#126](https://github.com/asteris-llc/converge/pull/126) ([BrianHicks](https://github.com/BrianHicks))
- Rename "template" to "file.content" [\#124](https://github.com/asteris-llc/converge/pull/124) ([BrianHicks](https://github.com/BrianHicks))
- gometalinter fixes [\#122](https://github.com/asteris-llc/converge/pull/122) ([BrianHicks](https://github.com/BrianHicks))
- Graphviz [\#118](https://github.com/asteris-llc/converge/pull/118) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Rewrite graph system [\#116](https://github.com/asteris-llc/converge/pull/116) ([BrianHicks](https://github.com/BrianHicks))
- resource\(shell\): add interpreter [\#109](https://github.com/asteris-llc/converge/pull/109) ([BrianHicks](https://github.com/BrianHicks))
- switch to glide [\#108](https://github.com/asteris-llc/converge/pull/108) ([BrianHicks](https://github.com/BrianHicks))
- Feature/context [\#107](https://github.com/asteris-llc/converge/pull/107) ([BrianHicks](https://github.com/BrianHicks))
- Format samples in readme [\#105](https://github.com/asteris-llc/converge/pull/105) ([ajhager](https://github.com/ajhager))
- Port shell resource and tests [\#101](https://github.com/asteris-llc/converge/pull/101) ([ajhager](https://github.com/ajhager))
- resource\(template\): implement and test template [\#98](https://github.com/asteris-llc/converge/pull/98) ([BrianHicks](https://github.com/BrianHicks))
- Application for preparers [\#97](https://github.com/asteris-llc/converge/pull/97) ([BrianHicks](https://github.com/BrianHicks))
- Plan in new preparers [\#95](https://github.com/asteris-llc/converge/pull/95) ([BrianHicks](https://github.com/BrianHicks))
- dependencies should block execution [\#81](https://github.com/asteris-llc/converge/pull/81) ([siddharthist](https://github.com/siddharthist))
- dependency tracker [\#80](https://github.com/asteris-llc/converge/pull/80) ([BrianHicks](https://github.com/BrianHicks))
- Self Serve [\#79](https://github.com/asteris-llc/converge/pull/79) ([BrianHicks](https://github.com/BrianHicks))
- Module server [\#78](https://github.com/asteris-llc/converge/pull/78) ([BrianHicks](https://github.com/BrianHicks))
- File mode module [\#76](https://github.com/asteris-llc/converge/pull/76) ([BrianHicks](https://github.com/BrianHicks))
- Simplify Application [\#73](https://github.com/asteris-llc/converge/pull/73) ([BrianHicks](https://github.com/BrianHicks))
- Fix graph with dashes [\#71](https://github.com/asteris-llc/converge/pull/71) ([BrianHicks](https://github.com/BrianHicks))
- load: from http [\#23](https://github.com/asteris-llc/converge/pull/23) ([siddharthist](https://github.com/siddharthist))

## [0.1.1](https://github.com/asteris-llc/converge/tree/0.1.1) (2016-06-13)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.1...0.1.1)

### Enhancements

- log lines from DAG walking should not appear [\#67](https://github.com/asteris-llc/converge/issues/67)

### Bugs

- color is not working? [\#64](https://github.com/asteris-llc/converge/issues/64)
- recursive modules are valid [\#63](https://github.com/asteris-llc/converge/issues/63)

### Closed Pull Requests

- Feature/logging [\#70](https://github.com/asteris-llc/converge/pull/70) ([BrianHicks](https://github.com/BrianHicks))
- add color printing again [\#69](https://github.com/asteris-llc/converge/pull/69) ([siddharthist](https://github.com/siddharthist))
- Feature/basic blackbox testing [\#68](https://github.com/asteris-llc/converge/pull/68) ([BrianHicks](https://github.com/BrianHicks))

## [0.1](https://github.com/asteris-llc/converge/tree/0.1) (2016-06-09)
[Full Changelog](https://github.com/asteris-llc/converge/compare/Unreleased...0.1)

### Enhancements

- converge format [\#55](https://github.com/asteris-llc/converge/issues/55)
- colorized summary lines [\#45](https://github.com/asteris-llc/converge/issues/45)
- print a nice summary of changes \(or no changes\) after commands [\#35](https://github.com/asteris-llc/converge/issues/35)
- Progressive output in check/apply [\#34](https://github.com/asteris-llc/converge/issues/34)
- Supplying params is hard [\#33](https://github.com/asteris-llc/converge/issues/33)
- Template should accept file permissions [\#31](https://github.com/asteris-llc/converge/issues/31)
- add net/context to Check [\#30](https://github.com/asteris-llc/converge/issues/30)
- templating in param defaults [\#27](https://github.com/asteris-llc/converge/issues/27)
- passing in parameters to top-level module [\#26](https://github.com/asteris-llc/converge/issues/26)
- Apply [\#15](https://github.com/asteris-llc/converge/issues/15)
- HTTPS fetching [\#14](https://github.com/asteris-llc/converge/issues/14)
- Nicer plan output [\#12](https://github.com/asteris-llc/converge/issues/12)
- Checker [\#11](https://github.com/asteris-llc/converge/issues/11)
- Requirements [\#10](https://github.com/asteris-llc/converge/issues/10)

### Bugs

- duplicate names are lost in the graph [\#56](https://github.com/asteris-llc/converge/issues/56)
- fill in short and long descriptions [\#51](https://github.com/asteris-llc/converge/issues/51)
- don't ignore error when Viper binds pflags [\#48](https://github.com/asteris-llc/converge/issues/48)
- param names should have limited characters [\#41](https://github.com/asteris-llc/converge/issues/41)
- template permissions should be 600 by default [\#38](https://github.com/asteris-llc/converge/issues/38)
- exec.Check tests [\#21](https://github.com/asteris-llc/converge/issues/21)
- Incomplete comment  [\#5](https://github.com/asteris-llc/converge/issues/5)

### Closed Issues

- README is out of date [\#52](https://github.com/asteris-llc/converge/issues/52)

### Closed Pull Requests

- Update 0.1 docs [\#60](https://github.com/asteris-llc/converge/pull/60) ([BrianHicks](https://github.com/BrianHicks))
- fixed souceFile.hcl [\#59](https://github.com/asteris-llc/converge/pull/59) ([Dacode45](https://github.com/Dacode45))
- Bug/names [\#58](https://github.com/asteris-llc/converge/pull/58) ([Dacode45](https://github.com/Dacode45))
- fmt [\#57](https://github.com/asteris-llc/converge/pull/57) ([BrianHicks](https://github.com/BrianHicks))
- Feature/requirements [\#54](https://github.com/asteris-llc/converge/pull/54) ([Dacode45](https://github.com/Dacode45))
- don't ignore the error when Viper binds to PFlags [\#50](https://github.com/asteris-llc/converge/pull/50) ([siddharthist](https://github.com/siddharthist))
- color summaries [\#49](https://github.com/asteris-llc/converge/pull/49) ([siddharthist](https://github.com/siddharthist))
- add key=value parameter parsing [\#47](https://github.com/asteris-llc/converge/pull/47) ([siddharthist](https://github.com/siddharthist))
- Summary lines [\#46](https://github.com/asteris-llc/converge/pull/46) ([BrianHicks](https://github.com/BrianHicks))
- render: fix swallowed errors [\#44](https://github.com/asteris-llc/converge/pull/44) ([BrianHicks](https://github.com/BrianHicks))
- limit identifiers to word chars, dashes, and dots [\#43](https://github.com/asteris-llc/converge/pull/43) ([BrianHicks](https://github.com/BrianHicks))
- resource\(template\): fix permissions [\#42](https://github.com/asteris-llc/converge/pull/42) ([BrianHicks](https://github.com/BrianHicks))
- Lazily load parameters - defaults can be templated [\#37](https://github.com/asteris-llc/converge/pull/37) ([siddharthist](https://github.com/siddharthist))
- commands: add streaming output [\#36](https://github.com/asteris-llc/converge/pull/36) ([BrianHicks](https://github.com/BrianHicks))
- Apply [\#32](https://github.com/asteris-llc/converge/pull/32) ([BrianHicks](https://github.com/BrianHicks))
- load: accept initial parameters [\#28](https://github.com/asteris-llc/converge/pull/28) ([BrianHicks](https://github.com/BrianHicks))
- Add tests for Plan [\#25](https://github.com/asteris-llc/converge/pull/25) ([BrianHicks](https://github.com/BrianHicks))
- plan: pretty-printing terminal output [\#22](https://github.com/asteris-llc/converge/pull/22) ([siddharthist](https://github.com/siddharthist))
- fix incomplete comment [\#19](https://github.com/asteris-llc/converge/pull/19) ([BrianHicks](https://github.com/BrianHicks))
- Checker [\#9](https://github.com/asteris-llc/converge/pull/9) ([BrianHicks](https://github.com/BrianHicks))
- Name\(\) -\> String\(\) [\#8](https://github.com/asteris-llc/converge/pull/8) ([siddharthist](https://github.com/siddharthist))
- \[WIP\] Graphviz Support [\#7](https://github.com/asteris-llc/converge/pull/7) ([BrianHicks](https://github.com/BrianHicks))
- Feature/validation error type [\#3](https://github.com/asteris-llc/converge/pull/3) ([siddharthist](https://github.com/siddharthist))
- shelltask: add 'validate' method [\#2](https://github.com/asteris-llc/converge/pull/2) ([siddharthist](https://github.com/siddharthist))
- readme: status -\> check [\#1](https://github.com/asteris-llc/converge/pull/1) ([siddharthist](https://github.com/siddharthist))

# Change Log

## [Unreleased](https://github.com/asteris-llc/converge/tree/HEAD)

[Full Changelog](https://github.com/asteris-llc/converge/compare/0.5.0-beta1...HEAD)

### Closed Issues

- task.query behavior [\#567](https://github.com/asteris-llc/converge/issues/567)
- Get rid of vendor directory [\#545](https://github.com/asteris-llc/converge/issues/545)

### Closed Pull Requests

- docs: only list exported fields in param and export sections [\#561](https://github.com/asteris-llc/converge/pull/561) ([arichardet](https://github.com/arichardet))
- Fix Linting Bugs [\#554](https://github.com/asteris-llc/converge/pull/554) ([BrianHicks](https://github.com/BrianHicks))

## [0.5.0-beta1](https://github.com/asteris-llc/converge/tree/0.5.0-beta1) (2016-12-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0...0.5.0-beta1)

### Enhancements

- Handle time.Time in preparer [\#490](https://github.com/asteris-llc/converge/issues/490)
- k8s demo [\#480](https://github.com/asteris-llc/converge/issues/480)
- no mo' spaces [\#320](https://github.com/asteris-llc/converge/issues/320)
- We're using too many ports [\#286](https://github.com/asteris-llc/converge/issues/286)
- better status display on plan and execution [\#269](https://github.com/asteris-llc/converge/issues/269)
- builtin file owner module [\#74](https://github.com/asteris-llc/converge/issues/74)
- Add interceptors for RPC auth [\#531](https://github.com/asteris-llc/converge/pull/531) ([BrianHicks](https://github.com/BrianHicks))
- demo: kubernetes - use latest version of converge [\#524](https://github.com/asteris-llc/converge/pull/524) ([ryane](https://github.com/ryane))
- fix shasum filename [\#522](https://github.com/asteris-llc/converge/pull/522) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/improve output [\#516](https://github.com/asteris-llc/converge/pull/516) ([arichardet](https://github.com/arichardet))

### Bugs

- LVM FS module reports failure due to still having changes [\#555](https://github.com/asteris-llc/converge/issues/555)
- handle newlines in docs generation for exported fields [\#548](https://github.com/asteris-llc/converge/issues/548)
- docker.container panics if image not present [\#538](https://github.com/asteris-llc/converge/issues/538)
- Return error instead of message from docker resource plan [\#529](https://github.com/asteris-llc/converge/issues/529)
- misleading error message in converge graph command [\#520](https://github.com/asteris-llc/converge/issues/520)
- exit code is 0 when converge plan has errors [\#492](https://github.com/asteris-llc/converge/issues/492)
- error specifying file mode [\#487](https://github.com/asteris-llc/converge/issues/487)
- docker.network misleading error messages [\#485](https://github.com/asteris-llc/converge/issues/485)
- inconsistent indentation on multi-line diffs [\#478](https://github.com/asteris-llc/converge/issues/478)
- pass typed params through to modules [\#409](https://github.com/asteris-llc/converge/issues/409)
- auth code is repeated [\#273](https://github.com/asteris-llc/converge/issues/273)
- Remove indent for multi-line diffs [\#530](https://github.com/asteris-llc/converge/pull/530) ([arichardet](https://github.com/arichardet))
- status: add MayChange status and Warning display [\#528](https://github.com/asteris-llc/converge/pull/528) ([ryane](https://github.com/ryane))
- Fix/disable docker solaris [\#527](https://github.com/asteris-llc/converge/pull/527) ([ryane](https://github.com/ryane))
- plan cmd: return non-zero exit code when plan has error\(s\) [\#525](https://github.com/asteris-llc/converge/pull/525) ([ryane](https://github.com/ryane))
- graph cmd: fix wrong argument length error [\#521](https://github.com/asteris-llc/converge/pull/521) ([ryane](https://github.com/ryane))
- Fix/file.mode errors [\#519](https://github.com/asteris-llc/converge/pull/519) ([ryane](https://github.com/ryane))
- docker: update docker dependency [\#512](https://github.com/asteris-llc/converge/pull/512) ([ryane](https://github.com/ryane))

### Closed Issues

- Better output for cascading failures [\#367](https://github.com/asteris-llc/converge/issues/367)
- file download/fetch module [\#250](https://github.com/asteris-llc/converge/issues/250)

### Closed Pull Requests

- changelog branch; tag and push tags before merging [\#559](https://github.com/asteris-llc/converge/pull/559) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Minor Fixes After Makefile rewrite [\#558](https://github.com/asteris-llc/converge/pull/558) ([BrianHicks](https://github.com/BrianHicks))
- reset updates needed after apply in lvm [\#557](https://github.com/asteris-llc/converge/pull/557) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Makefile: remove ref to VERSION in PACKAGE\_VERSION [\#556](https://github.com/asteris-llc/converge/pull/556) ([BrianHicks](https://github.com/BrianHicks))
- Better Makefiles [\#553](https://github.com/asteris-llc/converge/pull/553) ([BrianHicks](https://github.com/BrianHicks))
- when apply fails due to changes still being present; use those diffs … [\#551](https://github.com/asteris-llc/converge/pull/551) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docs: use same formatting for \(re\)exported fields as parameters [\#550](https://github.com/asteris-llc/converge/pull/550) ([arichardet](https://github.com/arichardet))
- Feature/74 file owner [\#549](https://github.com/asteris-llc/converge/pull/549) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/file.fetch [\#543](https://github.com/asteris-llc/converge/pull/543) ([arichardet](https://github.com/arichardet))
- Fix/538 docker container panic [\#542](https://github.com/asteris-llc/converge/pull/542) ([rebeccaskinner](https://github.com/rebeccaskinner))
- return a better error for planned network failures [\#541](https://github.com/asteris-llc/converge/pull/541) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Timer status display [\#537](https://github.com/asteris-llc/converge/pull/537) ([BrianHicks](https://github.com/BrianHicks))
- Add fuzzing and benchmarks to CI [\#536](https://github.com/asteris-llc/converge/pull/536) ([BrianHicks](https://github.com/BrianHicks))
- disallow some characters as resource names [\#535](https://github.com/asteris-llc/converge/pull/535) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use a single port [\#534](https://github.com/asteris-llc/converge/pull/534) ([BrianHicks](https://github.com/BrianHicks))
- Fix/374 immutable status [\#533](https://github.com/asteris-llc/converge/pull/533) ([rebeccaskinner](https://github.com/rebeccaskinner))
- add logo to Readme [\#532](https://github.com/asteris-llc/converge/pull/532) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docs [\#526](https://github.com/asteris-llc/converge/pull/526) ([arichardet](https://github.com/arichardet))
- Feature/preparer time.time [\#523](https://github.com/asteris-llc/converge/pull/523) ([arichardet](https://github.com/arichardet))
- update readme with graph and download info [\#518](https://github.com/asteris-llc/converge/pull/518) ([stevendborrelli](https://github.com/stevendborrelli))
- kubernetes example [\#507](https://github.com/asteris-llc/converge/pull/507) ([ryane](https://github.com/ryane))

## [0.4.0](https://github.com/asteris-llc/converge/tree/0.4.0) (2016-11-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-rc1...0.4.0)

### Enhancements

- Add `nonempty` tag to preparer [\#486](https://github.com/asteris-llc/converge/issues/486)
- apt module [\#404](https://github.com/asteris-llc/converge/issues/404)
- Handle time.Duration in Preparer [\#358](https://github.com/asteris-llc/converge/issues/358)
- binary distribution [\#306](https://github.com/asteris-llc/converge/issues/306)
- docker.volume resource [\#299](https://github.com/asteris-llc/converge/issues/299)
- docker.network resource [\#298](https://github.com/asteris-llc/converge/issues/298)
- Add ability to modify user [\#278](https://github.com/asteris-llc/converge/issues/278)
- LVM module [\#270](https://github.com/asteris-llc/converge/issues/270)
- update installer to point to 0.4.0 [\#517](https://github.com/asteris-llc/converge/pull/517) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/preparer time.duration [\#483](https://github.com/asteris-llc/converge/pull/483) ([arichardet](https://github.com/arichardet))
- check rpm for empty packages [\#482](https://github.com/asteris-llc/converge/pull/482) ([stevendborrelli](https://github.com/stevendborrelli))
- only extract binary from tarball [\#481](https://github.com/asteris-llc/converge/pull/481) ([stevendborrelli](https://github.com/stevendborrelli))
- docker.network resource [\#477](https://github.com/asteris-llc/converge/pull/477) ([ryane](https://github.com/ryane))
- Feature/apt package [\#461](https://github.com/asteris-llc/converge/pull/461) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docker.volume [\#453](https://github.com/asteris-llc/converge/pull/453) ([ryane](https://github.com/ryane))
- Add the ability to modify a user [\#434](https://github.com/asteris-llc/converge/pull/434) ([arichardet](https://github.com/arichardet))

### Bugs

- validate non-empty string values on some resources [\#337](https://github.com/asteris-llc/converge/issues/337)
- Intermittent test failure with `make test` [\#476](https://github.com/asteris-llc/converge/issues/476)
- Converge loses dependencies between modules within a branch when  [\#427](https://github.com/asteris-llc/converge/issues/427)
- diff output is partially bold [\#402](https://github.com/asteris-llc/converge/issues/402)
- Apply doesn't show diff output [\#399](https://github.com/asteris-llc/converge/issues/399)
- values that don't implement Stringer log badly [\#398](https://github.com/asteris-llc/converge/issues/398)
- resolving conditional macros line [\#389](https://github.com/asteris-llc/converge/issues/389)
- Case predicates fail when they include `lookup` [\#386](https://github.com/asteris-llc/converge/issues/386)
- converge exit code on a failed run [\#365](https://github.com/asteris-llc/converge/issues/365)
- Address name conflicts [\#361](https://github.com/asteris-llc/converge/issues/361)
- resources should have the ability to gracefully stop on interrupt [\#319](https://github.com/asteris-llc/converge/issues/319)
- converge panics when interrupted [\#318](https://github.com/asteris-llc/converge/issues/318)
- Fix/conditional nodes [\#484](https://github.com/asteris-llc/converge/pull/484) ([rebeccaskinner](https://github.com/rebeccaskinner))
- grouping: fix issues with groups and conditionals [\#460](https://github.com/asteris-llc/converge/pull/460) ([ryane](https://github.com/ryane))
- return non-zero exit code if apply graph has an error [\#436](https://github.com/asteris-llc/converge/pull/436) ([ryane](https://github.com/ryane))
- replace "context" with "golang.org/x/net/context" [\#433](https://github.com/asteris-llc/converge/pull/433) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- update changelog, add script [\#511](https://github.com/asteris-llc/converge/pull/511) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: fix link in resource authors guide [\#479](https://github.com/asteris-llc/converge/pull/479) ([ryane](https://github.com/ryane))
- Fix/resource duplication [\#458](https://github.com/asteris-llc/converge/pull/458) ([arichardet](https://github.com/arichardet))
- Fix/graceful exit [\#447](https://github.com/asteris-llc/converge/pull/447) ([arichardet](https://github.com/arichardet))
- LVM module [\#184](https://github.com/asteris-llc/converge/pull/184) ([avnik](https://github.com/avnik))

## [0.4.0-rc1](https://github.com/asteris-llc/converge/tree/0.4.0-rc1) (2016-11-17)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-beta1...0.4.0-rc1)

### Closed Pull Requests

- update changelog to reflect rc1 status [\#510](https://github.com/asteris-llc/converge/pull/510) ([rebeccaskinner](https://github.com/rebeccaskinner))
- move rendering plant creation inside of the transform block to avoid … [\#509](https://github.com/asteris-llc/converge/pull/509) ([rebeccaskinner](https://github.com/rebeccaskinner))
- set github link to the converge repo [\#508](https://github.com/asteris-llc/converge/pull/508) ([rebeccaskinner](https://github.com/rebeccaskinner))
- empty commit for CI [\#506](https://github.com/asteris-llc/converge/pull/506) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.4.0-beta1](https://github.com/asteris-llc/converge/tree/0.4.0-beta1) (2016-11-15)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0...0.4.0-beta1)

### Enhancements

- file.owner resource [\#503](https://github.com/asteris-llc/converge/issues/503)

### Bugs

- package.rpm state does not default to present [\#446](https://github.com/asteris-llc/converge/issues/446)
- Error when adding a user when a group of the same name exists [\#425](https://github.com/asteris-llc/converge/issues/425)

### Closed Issues

- nonempty check in preparer for user/group [\#489](https://github.com/asteris-llc/converge/issues/489)

### Closed Pull Requests

- Fix/0.4.0 beta1 fix [\#505](https://github.com/asteris-llc/converge/pull/505) ([rebeccaskinner](https://github.com/rebeccaskinner))
- release notes [\#497](https://github.com/asteris-llc/converge/pull/497) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use %v for logging [\#496](https://github.com/asteris-llc/converge/pull/496) ([arichardet](https://github.com/arichardet))
- update docs with conditional changes [\#495](https://github.com/asteris-llc/converge/pull/495) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/nonempty tag [\#494](https://github.com/asteris-llc/converge/pull/494) ([arichardet](https://github.com/arichardet))
- Remove extract executable [\#488](https://github.com/asteris-llc/converge/pull/488) ([arichardet](https://github.com/arichardet))
- Manifest / index page [\#464](https://github.com/asteris-llc/converge/pull/464) ([BrianHicks](https://github.com/BrianHicks))
- default to state 'present' for rpm module [\#463](https://github.com/asteris-llc/converge/pull/463) ([rebeccaskinner](https://github.com/rebeccaskinner))
- vendor: upgrade all deps [\#459](https://github.com/asteris-llc/converge/pull/459) ([BrianHicks](https://github.com/BrianHicks))

## [0.3.0](https://github.com/asteris-llc/converge/tree/0.3.0) (2016-10-27)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-rc1...0.3.0)

### Enhancements

- "param is required" error should include param name [\#335](https://github.com/asteris-llc/converge/issues/335)
- Add ability to modify group [\#279](https://github.com/asteris-llc/converge/issues/279)
- add conditionals to module resource [\#268](https://github.com/asteris-llc/converge/issues/268)
- named locks [\#249](https://github.com/asteris-llc/converge/issues/249)
- pretty printed changes should align values [\#244](https://github.com/asteris-llc/converge/issues/244)
- ability to wait for a condition [\#238](https://github.com/asteris-llc/converge/issues/238)
- proper arrays [\#110](https://github.com/asteris-llc/converge/issues/110)
- makefile: add hashes to files [\#440](https://github.com/asteris-llc/converge/pull/440) ([BrianHicks](https://github.com/BrianHicks))
- \[389\] switch log level from Info to Debug [\#428](https://github.com/asteris-llc/converge/pull/428) ([mason-fish](https://github.com/mason-fish))
- Feature/examples 0.3.0 [\#419](https://github.com/asteris-llc/converge/pull/419) ([ryane](https://github.com/ryane))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- named groups [\#392](https://github.com/asteris-llc/converge/pull/392) ([ryane](https://github.com/ryane))
- Don't render module dependencies [\#387](https://github.com/asteris-llc/converge/pull/387) ([BrianHicks](https://github.com/BrianHicks))
- Added docs, renamed packages to package to  for consistency with other modules [\#384](https://github.com/asteris-llc/converge/pull/384) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use status.RaiseLevel in group [\#377](https://github.com/asteris-llc/converge/pull/377) ([arichardet](https://github.com/arichardet))
- Feature/rpm module [\#373](https://github.com/asteris-llc/converge/pull/373) ([rebeccaskinner](https://github.com/rebeccaskinner))
- use a container for CI [\#372](https://github.com/asteris-llc/converge/pull/372) ([BrianHicks](https://github.com/BrianHicks))
- create a metadata envelope for nodes [\#369](https://github.com/asteris-llc/converge/pull/369) ([BrianHicks](https://github.com/BrianHicks))
- cmd/server.go: use errgroup instead of waitgroup [\#363](https://github.com/asteris-llc/converge/pull/363) ([QuentinPerez](https://github.com/QuentinPerez))
- Feature/conditionals [\#362](https://github.com/asteris-llc/converge/pull/362) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/example swarm wait [\#346](https://github.com/asteris-llc/converge/pull/346) ([ryane](https://github.com/ryane))
- Add Compound Parameters [\#340](https://github.com/asteris-llc/converge/pull/340) ([BrianHicks](https://github.com/BrianHicks))
- load overlay module in elk example [\#326](https://github.com/asteris-llc/converge/pull/326) ([feniix](https://github.com/feniix))
- Refactor deserializers [\#321](https://github.com/asteris-llc/converge/pull/321) ([BrianHicks](https://github.com/BrianHicks))
- Use text/tabwriter to align human output [\#317](https://github.com/asteris-llc/converge/pull/317) ([sehqlr](https://github.com/sehqlr))
- Fix/225 pipeline function refactor [\#307](https://github.com/asteris-llc/converge/pull/307) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Bugs

- Looking up a node from a branch that is also a part of a group introduced a deadlock [\#415](https://github.com/asteris-llc/converge/issues/415)
- user/group not activating changes [\#417](https://github.com/asteris-llc/converge/issues/417)
- docker.container reporting incorrect entrypoint diff [\#410](https://github.com/asteris-llc/converge/issues/410)
- converge panics when encountering an empty conditional [\#407](https://github.com/asteris-llc/converge/issues/407)
- Use lists as params in modules is broken [\#397](https://github.com/asteris-llc/converge/issues/397)
- Explicit dependencies fail inside of case statements [\#385](https://github.com/asteris-llc/converge/issues/385)
- Use `package.rpm` not `rpm.package` for rpm module [\#382](https://github.com/asteris-llc/converge/issues/382)
- shell module should not set StatusLevel based on process exit code [\#323](https://github.com/asteris-llc/converge/issues/323)
- StatusLevel not taken into account during graph execution [\#322](https://github.com/asteris-llc/converge/issues/322)
- param dependency fails in `samples/shellContext.hcl` [\#313](https://github.com/asteris-llc/converge/issues/313)
- Execution Engine Ignoring Warning Levels [\#243](https://github.com/asteris-llc/converge/issues/243)
- Make pipeline functions use mult-return instead of Either [\#225](https://github.com/asteris-llc/converge/issues/225)
- Fix 420 go test race [\#424](https://github.com/asteris-llc/converge/pull/424) ([ryane](https://github.com/ryane))
- Replace `AreSiblingIDs` with `AreSiblings` to prevent deadlocking [\#423](https://github.com/asteris-llc/converge/pull/423) ([rebeccaskinner](https://github.com/rebeccaskinner))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- Handle thunked branches [\#403](https://github.com/asteris-llc/converge/pull/403) ([rebeccaskinner](https://github.com/rebeccaskinner))
- don't exclude modules in getNearestAncestor [\#396](https://github.com/asteris-llc/converge/pull/396) ([ryane](https://github.com/ryane))
- Fix/385 dependencies in conditionals [\#391](https://github.com/asteris-llc/converge/pull/391) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix lookup calls to use os/user in SetAddUserOptions [\#390](https://github.com/asteris-llc/converge/pull/390) ([arichardet](https://github.com/arichardet))
- Change `rpm.package` to `package.rpm` [\#383](https://github.com/asteris-llc/converge/pull/383) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Update user status level and errors, add checks in Apply for group [\#375](https://github.com/asteris-llc/converge/pull/375) ([arichardet](https://github.com/arichardet))
- Ability to use pointers as a preparer value [\#339](https://github.com/asteris-llc/converge/pull/339) ([BrianHicks](https://github.com/BrianHicks))
- Status Error Codes [\#333](https://github.com/asteris-llc/converge/pull/333) ([BrianHicks](https://github.com/BrianHicks))
- fix codeclimate yaml [\#328](https://github.com/asteris-llc/converge/pull/328) ([BrianHicks](https://github.com/BrianHicks))

### Closed Issues

- docs for rpm module [\#381](https://github.com/asteris-llc/converge/issues/381)

### Closed Pull Requests

- update for 0.3.0 [\#443](https://github.com/asteris-llc/converge/pull/443) ([stevendborrelli](https://github.com/stevendborrelli))
- prep for 0.3.0-rc1 [\#437](https://github.com/asteris-llc/converge/pull/437) ([stevendborrelli](https://github.com/stevendborrelli))
- update release notes [\#435](https://github.com/asteris-llc/converge/pull/435) ([stevendborrelli](https://github.com/stevendborrelli))
- beta3 release notes [\#431](https://github.com/asteris-llc/converge/pull/431) ([stevendborrelli](https://github.com/stevendborrelli))
- update docs to include currently identified switch statement limitations [\#388](https://github.com/asteris-llc/converge/pull/388) ([rebeccaskinner](https://github.com/rebeccaskinner))
- update vendored dependencies [\#378](https://github.com/asteris-llc/converge/pull/378) ([BrianHicks](https://github.com/BrianHicks))
- 0.3.0 release documentation [\#376](https://github.com/asteris-llc/converge/pull/376) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: add multi-version control [\#371](https://github.com/asteris-llc/converge/pull/371) ([BrianHicks](https://github.com/BrianHicks))
- add docs for the platform module [\#342](https://github.com/asteris-llc/converge/pull/342) ([stevendborrelli](https://github.com/stevendborrelli))
- update for PR \#307 [\#316](https://github.com/asteris-llc/converge/pull/316) ([stevendborrelli](https://github.com/stevendborrelli))

## [0.3.0-rc1](https://github.com/asteris-llc/converge/tree/0.3.0-rc1) (2016-10-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta3...0.3.0-rc1)

### Closed Issues

- Update documentation regarding conditionals and groups [\#430](https://github.com/asteris-llc/converge/issues/430)

## [0.3.0-beta3](https://github.com/asteris-llc/converge/tree/0.3.0-beta3) (2016-10-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta2...0.3.0-beta3)

### Enhancements

- add installation script [\#416](https://github.com/asteris-llc/converge/pull/416) ([stevendborrelli](https://github.com/stevendborrelli))

### Bugs

- go tests race condition [\#420](https://github.com/asteris-llc/converge/issues/420)

### Closed Pull Requests

- document known bug between conditionals and groups [\#432](https://github.com/asteris-llc/converge/pull/432) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.3.0-beta2](https://github.com/asteris-llc/converge/tree/0.3.0-beta2) (2016-10-24)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta1...0.3.0-beta2)

### Bugs

- Fix container entrypoint diffs [\#412](https://github.com/asteris-llc/converge/pull/412) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- wercker: don't delete removed [\#414](https://github.com/asteris-llc/converge/pull/414) ([BrianHicks](https://github.com/BrianHicks))
- Nightly/test builds [\#413](https://github.com/asteris-llc/converge/pull/413) ([BrianHicks](https://github.com/BrianHicks))
- fix panic [\#408](https://github.com/asteris-llc/converge/pull/408) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.3.0-beta1](https://github.com/asteris-llc/converge/tree/0.3.0-beta1) (2016-10-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/before-moving-resolve-conditionals...0.3.0-beta1)

### Enhancements

- file/dir refactor [\#327](https://github.com/asteris-llc/converge/issues/327)
- Use errgroup in server [\#295](https://github.com/asteris-llc/converge/issues/295)
- Allow ability to indicate group name when adding a user [\#276](https://github.com/asteris-llc/converge/issues/276)
- builtin file group module [\#75](https://github.com/asteris-llc/converge/issues/75)

### Bugs

- Conditional Regression [\#401](https://github.com/asteris-llc/converge/issues/401)
- module dependencies now failing [\#395](https://github.com/asteris-llc/converge/issues/395)
- docker.container regression [\#343](https://github.com/asteris-llc/converge/issues/343)
- samples in the README are formatted incorrectly. [\#104](https://github.com/asteris-llc/converge/issues/104)

### Closed Issues

- Tidy up codebase by fixing `make lint` reports [\#348](https://github.com/asteris-llc/converge/issues/348)
- Can we simplify the visualization of large modules? [\#280](https://github.com/asteris-llc/converge/issues/280)
- How should we manage module docs and examples? [\#173](https://github.com/asteris-llc/converge/issues/173)

### Closed Pull Requests

- Allow lists in parameters when called in a module [\#400](https://github.com/asteris-llc/converge/pull/400) ([BrianHicks](https://github.com/BrianHicks))
- comment typo [\#380](https://github.com/asteris-llc/converge/pull/380) ([rebeccaskinner](https://github.com/rebeccaskinner))
- FIX local command flag added to example [\#370](https://github.com/asteris-llc/converge/pull/370) ([philcryer](https://github.com/philcryer))
- codeclimate: remove markdown linting [\#368](https://github.com/asteris-llc/converge/pull/368) ([BrianHicks](https://github.com/BrianHicks))
- Add testing section to contribution guidelines [\#366](https://github.com/asteris-llc/converge/pull/366) ([arichardet](https://github.com/arichardet))
- remove binary [\#354](https://github.com/asteris-llc/converge/pull/354) ([stevendborrelli](https://github.com/stevendborrelli))
- wait: fix error handling in retrier [\#353](https://github.com/asteris-llc/converge/pull/353) ([ryane](https://github.com/ryane))
- Feature/param name in error [\#352](https://github.com/asteris-llc/converge/pull/352) ([ryane](https://github.com/ryane))
- docs: add section on error values to resource author's guide [\#347](https://github.com/asteris-llc/converge/pull/347) ([BrianHicks](https://github.com/BrianHicks))
- Fix docker regressions [\#345](https://github.com/asteris-llc/converge/pull/345) ([ryane](https://github.com/ryane))
- wait.query and wait.port resources [\#334](https://github.com/asteris-llc/converge/pull/334) ([ryane](https://github.com/ryane))

## [before-moving-resolve-conditionals](https://github.com/asteris-llc/converge/tree/before-moving-resolve-conditionals) (2016-10-05)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0...before-moving-resolve-conditionals)

### Enhancements

- Use github.com/pkg/errors exclusively [\#300](https://github.com/asteris-llc/converge/issues/300)

### Bugs

- handle parameters with valid zero value [\#338](https://github.com/asteris-llc/converge/issues/338)

### Closed Issues

- Universal Benchmarking [\#308](https://github.com/asteris-llc/converge/issues/308)
- Use table tests instead of helpers.PreparerValidator [\#155](https://github.com/asteris-llc/converge/issues/155)

### Closed Pull Requests

- Handle preparer fields where zero is valid [\#344](https://github.com/asteris-llc/converge/pull/344) ([arichardet](https://github.com/arichardet))
- plan: test, which illustrate "return nil, err" problem [\#336](https://github.com/asteris-llc/converge/pull/336) ([avnik](https://github.com/avnik))
- Support for modifying linux groups [\#331](https://github.com/asteris-llc/converge/pull/331) ([arichardet](https://github.com/arichardet))
- Add contributing.md and code of conduct [\#330](https://github.com/asteris-llc/converge/pull/330) ([BrianHicks](https://github.com/BrianHicks))
- shell: non-0 exit code returns StatusWillChange [\#324](https://github.com/asteris-llc/converge/pull/324) ([ryane](https://github.com/ryane))
- map keys are considered "strings" in parse.Node [\#315](https://github.com/asteris-llc/converge/pull/315) ([rebeccaskinner](https://github.com/rebeccaskinner))
- helpers,docker: move AssertDiff to common testhelpers [\#312](https://github.com/asteris-llc/converge/pull/312) ([avnik](https://github.com/avnik))

## [0.2.0](https://github.com/asteris-llc/converge/tree/0.2.0) (2016-09-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-rc1...0.2.0)

### Enhancements

- Calls to `lookup` should be all lower-case [\#223](https://github.com/asteris-llc/converge/issues/223)

### Bugs

- field name conflicts during lookups [\#294](https://github.com/asteris-llc/converge/issues/294)
- panic when trying to lookup against file.directory [\#292](https://github.com/asteris-llc/converge/issues/292)
- Allow adding and deleting a group without specifying gid [\#275](https://github.com/asteris-llc/converge/issues/275)
- param lookup race condition [\#266](https://github.com/asteris-llc/converge/issues/266)
- Node is sometimes unresolvable, depending on ordering within a template [\#254](https://github.com/asteris-llc/converge/issues/254)
- params should handle boolean values [\#251](https://github.com/asteris-llc/converge/issues/251)
- remove token handling logic when only a client and not a server [\#247](https://github.com/asteris-llc/converge/issues/247)
- Check shouldn't be called from Apply [\#240](https://github.com/asteris-llc/converge/issues/240)
- value passing causes dependency resolution failures in child modules [\#239](https://github.com/asteris-llc/converge/issues/239)
- task.Task has extra argument [\#237](https://github.com/asteris-llc/converge/issues/237)
- Lookups fail for file.content [\#236](https://github.com/asteris-llc/converge/issues/236)
- RPC isn't used in healthchecks [\#199](https://github.com/asteris-llc/converge/issues/199)

### Closed Issues

- module author guide [\#291](https://github.com/asteris-llc/converge/issues/291)
- logging is a little too chatty [\#248](https://github.com/asteris-llc/converge/issues/248)
- Add RPC auth documentation [\#246](https://github.com/asteris-llc/converge/issues/246)

### Closed Pull Requests

- Add ability to indicate the group name when adding a user [\#310](https://github.com/asteris-llc/converge/pull/310) ([arichardet](https://github.com/arichardet))
- 0.2.0 release [\#311](https://github.com/asteris-llc/converge/pull/311) ([stevendborrelli](https://github.com/stevendborrelli))
- Add "Basic Usage" section to Server docs [\#287](https://github.com/asteris-llc/converge/pull/287) ([sehqlr](https://github.com/sehqlr))
- Simplify Status implementation [\#284](https://github.com/asteris-llc/converge/pull/284) ([BrianHicks](https://github.com/BrianHicks))
- Fix climate issues [\#282](https://github.com/asteris-llc/converge/pull/282) ([BrianHicks](https://github.com/BrianHicks))
- Add codeclimate checks [\#281](https://github.com/asteris-llc/converge/pull/281) ([BrianHicks](https://github.com/BrianHicks))
- Make the logs quieter [\#274](https://github.com/asteris-llc/converge/pull/274) ([BrianHicks](https://github.com/BrianHicks))
- example: elasticsearch, kibana, and filebeat ~elk [\#272](https://github.com/asteris-llc/converge/pull/272) ([ryane](https://github.com/ryane))
- param: accept boolean values [\#271](https://github.com/asteris-llc/converge/pull/271) ([BrianHicks](https://github.com/BrianHicks))
- Use `MakeLanguage` for dependecy resolution to ensure that no [\#265](https://github.com/asteris-llc/converge/pull/265) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Make notes match sample code [\#262](https://github.com/asteris-llc/converge/pull/262) ([tomduckering](https://github.com/tomduckering))

## [0.2.0-rc1](https://github.com/asteris-llc/converge/tree/0.2.0-rc1) (2016-09-23)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-beta2...0.2.0-rc1)

### Bugs

- docker.image incorrectly reporting changes after apply [\#304](https://github.com/asteris-llc/converge/issues/304)

### Closed Issues

- preprocessor race condition [\#302](https://github.com/asteris-llc/converge/issues/302)

### Closed Pull Requests

- Fixes docker resources and updates examples [\#305](https://github.com/asteris-llc/converge/pull/305) ([ryane](https://github.com/ryane))
- use thread-safe field cache [\#303](https://github.com/asteris-llc/converge/pull/303) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/294 field name conflicts [\#301](https://github.com/asteris-llc/converge/pull/301) ([rebeccaskinner](https://github.com/rebeccaskinner))
- preproc: don't add fields for non-struct anon field [\#293](https://github.com/asteris-llc/converge/pull/293) ([ryane](https://github.com/ryane))
- docs: add draft resource authors guide [\#290](https://github.com/asteris-llc/converge/pull/290) ([BrianHicks](https://github.com/BrianHicks))
- perform healthchecks over RPC [\#289](https://github.com/asteris-llc/converge/pull/289) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/239 [\#288](https://github.com/asteris-llc/converge/pull/288) ([rebeccaskinner](https://github.com/rebeccaskinner))
- cmd: don't set local token if only a client [\#285](https://github.com/asteris-llc/converge/pull/285) ([BrianHicks](https://github.com/BrianHicks))
- Fix/user group - Allow adding/deleting without gid [\#283](https://github.com/asteris-llc/converge/pull/283) ([arichardet](https://github.com/arichardet))

## [0.2.0-beta2](https://github.com/asteris-llc/converge/tree/0.2.0-beta2) (2016-09-20)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-beta1...0.2.0-beta2)

### Enhancements

- module signing [\#16](https://github.com/asteris-llc/converge/issues/16)

### Bugs

- This graph is rendered kinda funky [\#141](https://github.com/asteris-llc/converge/issues/141)

### Closed Issues

- Update all deps to the latest versions [\#171](https://github.com/asteris-llc/converge/issues/171)
- human printer should be profiled and optimized [\#156](https://github.com/asteris-llc/converge/issues/156)

### Closed Pull Requests

- example: docker swarm mode [\#267](https://github.com/asteris-llc/converge/pull/267) ([ryane](https://github.com/ryane))

## [0.2.0-beta1](https://github.com/asteris-llc/converge/tree/0.2.0-beta1) (2016-09-16)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.1.1...0.2.0-beta1)

### Enhancements

- Create a query module [\#227](https://github.com/asteris-llc/converge/issues/227)
- Graphing should work over RPC [\#207](https://github.com/asteris-llc/converge/issues/207)
- Allow Value-passing between modules [\#206](https://github.com/asteris-llc/converge/issues/206)
- logging needs improvement [\#200](https://github.com/asteris-llc/converge/issues/200)
- REST layer on RPC [\#170](https://github.com/asteris-llc/converge/issues/170)
- CLI RPC client [\#169](https://github.com/asteris-llc/converge/issues/169)
- RPC authentication [\#168](https://github.com/asteris-llc/converge/issues/168)
- RPC server [\#167](https://github.com/asteris-llc/converge/issues/167)
- Docker module [\#159](https://github.com/asteris-llc/converge/issues/159)
- Clear/readable errors [\#140](https://github.com/asteris-llc/converge/issues/140)
- Macros? [\#111](https://github.com/asteris-llc/converge/issues/111)
- should template be called file.content? [\#103](https://github.com/asteris-llc/converge/issues/103)
- execution in new graph code [\#92](https://github.com/asteris-llc/converge/issues/92)
- execution planning in new graph code [\#91](https://github.com/asteris-llc/converge/issues/91)
- implement file.group [\#90](https://github.com/asteris-llc/converge/issues/90)
- reimplement file.user [\#89](https://github.com/asteris-llc/converge/issues/89)
- reimplement file.mode [\#88](https://github.com/asteris-llc/converge/issues/88)
- reimplement templates [\#87](https://github.com/asteris-llc/converge/issues/87)
- reimplement shell tasks [\#86](https://github.com/asteris-llc/converge/issues/86)
- implement graph visualization on top of new graph code [\#85](https://github.com/asteris-llc/converge/issues/85)
- resources should keep track of their parents [\#83](https://github.com/asteris-llc/converge/issues/83)
- infer dependencies based on template contents [\#65](https://github.com/asteris-llc/converge/issues/65)
- "native" filemode module [\#40](https://github.com/asteris-llc/converge/issues/40)
- "native" modules addressing scheme [\#39](https://github.com/asteris-llc/converge/issues/39)
- basic diff finding [\#20](https://github.com/asteris-llc/converge/issues/20)
- authenticated module server [\#18](https://github.com/asteris-llc/converge/issues/18)
- module server [\#17](https://github.com/asteris-llc/converge/issues/17)
- HTTP fetching [\#13](https://github.com/asteris-llc/converge/issues/13)
- graph: keep track of parentage inside edges [\#163](https://github.com/asteris-llc/converge/pull/163) ([BrianHicks](https://github.com/BrianHicks))

### Bugs

- lookup should be fully case insensitive [\#232](https://github.com/asteris-llc/converge/issues/232)
- rpc server does not support graph operation anymore [\#205](https://github.com/asteris-llc/converge/issues/205)
- graph is walked out of order [\#160](https://github.com/asteris-llc/converge/issues/160)
- Concurrent Map Read/Write Panic When Running Tests [\#158](https://github.com/asteris-llc/converge/issues/158)
- Shell Doesn't Validate Against User-Defined Interpreter [\#119](https://github.com/asteris-llc/converge/issues/119)
- invalid script syntax doesn't return a good error [\#113](https://github.com/asteris-llc/converge/issues/113)
- Blackbox/test\_apply\_remote.sh Fails [\#102](https://github.com/asteris-llc/converge/issues/102)
- `helpers` is kind of an awful module name [\#96](https://github.com/asteris-llc/converge/issues/96)
- Failing tasks don't block tasks that depend on them [\#82](https://github.com/asteris-llc/converge/issues/82)
- all modules in samples must be valid and useful [\#66](https://github.com/asteris-llc/converge/issues/66)
- params should not have default dependencies [\#62](https://github.com/asteris-llc/converge/issues/62)
- resource names with dashes break graph rendering [\#61](https://github.com/asteris-llc/converge/issues/61)
- Fix error handling when fail to print results during plan or apply [\#183](https://github.com/asteris-llc/converge/pull/183) ([arichardet](https://github.com/arichardet))
- Return better error information for Shell module failures [\#121](https://github.com/asteris-llc/converge/pull/121) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Closed Issues

- errors from check silently swallowed during apply [\#213](https://github.com/asteris-llc/converge/issues/213)
- errors are printed with full stacks [\#208](https://github.com/asteris-llc/converge/issues/208)
- nil pointer panic when file.content resource errors [\#179](https://github.com/asteris-llc/converge/issues/179)
- intermittent concurrent map write panic with dependent params [\#176](https://github.com/asteris-llc/converge/issues/176)
- use standard library net/context library [\#161](https://github.com/asteris-llc/converge/issues/161)
- Subgraph PrettyPrinter should use `nil` as the `SubgraphBottomID` [\#149](https://github.com/asteris-llc/converge/issues/149)
- DigraphPrettyPrinter \(prettyprinters\) should shrink [\#147](https://github.com/asteris-llc/converge/issues/147)
- Renderable \(prettyprinters\) is larger than necessary [\#146](https://github.com/asteris-llc/converge/issues/146)
- Shell command context [\#145](https://github.com/asteris-llc/converge/issues/145)
- Add a 'Health' Module [\#144](https://github.com/asteris-llc/converge/issues/144)
- Builtin file.directory module [\#137](https://github.com/asteris-llc/converge/issues/137)
- Graphviz Generation Fails With Empty Param [\#129](https://github.com/asteris-llc/converge/issues/129)
- parameters should be required if no default is set [\#123](https://github.com/asteris-llc/converge/issues/123)
- Shell Module Error Codes Could Have Better Formatting [\#120](https://github.com/asteris-llc/converge/issues/120)
- params should allow commas [\#114](https://github.com/asteris-llc/converge/issues/114)
- modules and params should not show up by default [\#106](https://github.com/asteris-llc/converge/issues/106)
- use a context in loading [\#94](https://github.com/asteris-llc/converge/issues/94)
- builtin file owner and file permissions module [\#84](https://github.com/asteris-llc/converge/issues/84)

### Closed Pull Requests

- We got an assigned port so here's docs. [\#261](https://github.com/asteris-llc/converge/pull/261) ([BrianHicks](https://github.com/BrianHicks))
- update readme [\#260](https://github.com/asteris-llc/converge/pull/260) ([stevendborrelli](https://github.com/stevendborrelli))
- Add user support [\#259](https://github.com/asteris-llc/converge/pull/259) ([arichardet](https://github.com/arichardet))
- MOAR DOCS PLS [\#258](https://github.com/asteris-llc/converge/pull/258) ([BrianHicks](https://github.com/BrianHicks))
- Makefile: refine xcompile targets [\#257](https://github.com/asteris-llc/converge/pull/257) ([BrianHicks](https://github.com/BrianHicks))
- Fixed issue where apply swallowed error. [\#256](https://github.com/asteris-llc/converge/pull/256) ([Dacode45](https://github.com/Dacode45))
- add mutex to render [\#255](https://github.com/asteris-llc/converge/pull/255) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Add default group implementation for unsupported systems [\#252](https://github.com/asteris-llc/converge/pull/252) ([arichardet](https://github.com/arichardet))
- Feature/module verification [\#245](https://github.com/asteris-llc/converge/pull/245) ([ajhager](https://github.com/ajhager))
- Feature/fully lowercase field names [\#242](https://github.com/asteris-llc/converge/pull/242) ([rebeccaskinner](https://github.com/rebeccaskinner))
- file.directory resource [\#241](https://github.com/asteris-llc/converge/pull/241) ([BrianHicks](https://github.com/BrianHicks))
- Add user group support [\#234](https://github.com/asteris-llc/converge/pull/234) ([arichardet](https://github.com/arichardet))
- support lowercase field names [\#231](https://github.com/asteris-llc/converge/pull/231) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/query refactor [\#230](https://github.com/asteris-llc/converge/pull/230) ([rebeccaskinner](https://github.com/rebeccaskinner))
- handle params that are thunked [\#229](https://github.com/asteris-llc/converge/pull/229) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/query module [\#228](https://github.com/asteris-llc/converge/pull/228) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Refactor helpers into descriptive names [\#222](https://github.com/asteris-llc/converge/pull/222) ([BrianHicks](https://github.com/BrianHicks))
- Feature/value passing [\#220](https://github.com/asteris-llc/converge/pull/220) ([rebeccaskinner](https://github.com/rebeccaskinner))
- rpc: flatten details field [\#218](https://github.com/asteris-llc/converge/pull/218) ([BrianHicks](https://github.com/BrianHicks))
- Feature/formatter \#208 [\#217](https://github.com/asteris-llc/converge/pull/217) ([BrianHicks](https://github.com/BrianHicks))
- docker.container: fix entrypoint diff logic [\#215](https://github.com/asteris-llc/converge/pull/215) ([ryane](https://github.com/ryane))
- Graphing over RPC [\#214](https://github.com/asteris-llc/converge/pull/214) ([BrianHicks](https://github.com/BrianHicks))
- docker.container resource [\#203](https://github.com/asteris-llc/converge/pull/203) ([ryane](https://github.com/ryane))
- Rework logs to use logrus [\#202](https://github.com/asteris-llc/converge/pull/202) ([BrianHicks](https://github.com/BrianHicks))
- Feature/platform refactor [\#197](https://github.com/asteris-llc/converge/pull/197) ([stevendborrelli](https://github.com/stevendborrelli))
- Added host environment variable support [\#195](https://github.com/asteris-llc/converge/pull/195) ([arichardet](https://github.com/arichardet))
- resources\(docker image\): document [\#193](https://github.com/asteris-llc/converge/pull/193) ([BrianHicks](https://github.com/BrianHicks))
- Documentation Site [\#192](https://github.com/asteris-llc/converge/pull/192) ([BrianHicks](https://github.com/BrianHicks))
- Feature/health module [\#189](https://github.com/asteris-llc/converge/pull/189) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docker.image resource redux [\#188](https://github.com/asteris-llc/converge/pull/188) ([ryane](https://github.com/ryane))
- Basic RPC [\#187](https://github.com/asteris-llc/converge/pull/187) ([BrianHicks](https://github.com/BrianHicks))
- shell env and working dir support [\#185](https://github.com/asteris-llc/converge/pull/185) ([ryane](https://github.com/ryane))
- \[GRAPH\] Fix \#158, add triggering files, add tests [\#181](https://github.com/asteris-llc/converge/pull/181) ([sehqlr](https://github.com/sehqlr))
- don't panic when result status is nil [\#180](https://github.com/asteris-llc/converge/pull/180) ([ryane](https://github.com/ryane))
- Feature/optimize human printer [\#175](https://github.com/asteris-llc/converge/pull/175) ([arichardet](https://github.com/arichardet))
- Shrink binaries with some magical ldflags [\#174](https://github.com/asteris-llc/converge/pull/174) ([BrianHicks](https://github.com/BrianHicks))
- Feature/shell refactor [\#172](https://github.com/asteris-llc/converge/pull/172) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Dependency ordered graph walking [\#166](https://github.com/asteris-llc/converge/pull/166) ([BrianHicks](https://github.com/BrianHicks))
- Migrate context for go1.7, refactor fetch/http.go [\#164](https://github.com/asteris-llc/converge/pull/164) ([sehqlr](https://github.com/sehqlr))
- Feature/resource return refactor [\#157](https://github.com/asteris-llc/converge/pull/157) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Move errors into the tree [\#154](https://github.com/asteris-llc/converge/pull/154) ([BrianHicks](https://github.com/BrianHicks))
- Human-readable pretty printer [\#152](https://github.com/asteris-llc/converge/pull/152) ([BrianHicks](https://github.com/BrianHicks))
- make nil a useful default bottom value for subgraphdi [\#150](https://github.com/asteris-llc/converge/pull/150) ([rebeccaskinner](https://github.com/rebeccaskinner))
- JSONL output and interface refactoring [\#148](https://github.com/asteris-llc/converge/pull/148) ([BrianHicks](https://github.com/BrianHicks))
- Feature/template split [\#136](https://github.com/asteris-llc/converge/pull/136) ([rebeccaskinner](https://github.com/rebeccaskinner))
- fetch: add context [\#135](https://github.com/asteris-llc/converge/pull/135) ([BrianHicks](https://github.com/BrianHicks))
- Render merged subtrees [\#133](https://github.com/asteris-llc/converge/pull/133) ([BrianHicks](https://github.com/BrianHicks))
- Fix/119 [\#132](https://github.com/asteris-llc/converge/pull/132) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Trim Duplicate Subtrees [\#131](https://github.com/asteris-llc/converge/pull/131) ([BrianHicks](https://github.com/BrianHicks))
- Appropriately handle required parameters instead of crashing [\#130](https://github.com/asteris-llc/converge/pull/130) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use concurrent map implementation for graph values [\#128](https://github.com/asteris-llc/converge/pull/128) ([BrianHicks](https://github.com/BrianHicks))
- Make params missing defaults required [\#127](https://github.com/asteris-llc/converge/pull/127) ([BrianHicks](https://github.com/BrianHicks))
- makefile: simplify linting [\#126](https://github.com/asteris-llc/converge/pull/126) ([BrianHicks](https://github.com/BrianHicks))
- Rename "template" to "file.content" [\#124](https://github.com/asteris-llc/converge/pull/124) ([BrianHicks](https://github.com/BrianHicks))
- gometalinter fixes [\#122](https://github.com/asteris-llc/converge/pull/122) ([BrianHicks](https://github.com/BrianHicks))
- Graphviz [\#118](https://github.com/asteris-llc/converge/pull/118) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Rewrite graph system [\#116](https://github.com/asteris-llc/converge/pull/116) ([BrianHicks](https://github.com/BrianHicks))
- resource\(shell\): add interpreter [\#109](https://github.com/asteris-llc/converge/pull/109) ([BrianHicks](https://github.com/BrianHicks))
- switch to glide [\#108](https://github.com/asteris-llc/converge/pull/108) ([BrianHicks](https://github.com/BrianHicks))
- Feature/context [\#107](https://github.com/asteris-llc/converge/pull/107) ([BrianHicks](https://github.com/BrianHicks))
- Format samples in readme [\#105](https://github.com/asteris-llc/converge/pull/105) ([ajhager](https://github.com/ajhager))
- Port shell resource and tests [\#101](https://github.com/asteris-llc/converge/pull/101) ([ajhager](https://github.com/ajhager))
- resource\(template\): implement and test template [\#98](https://github.com/asteris-llc/converge/pull/98) ([BrianHicks](https://github.com/BrianHicks))
- Application for preparers [\#97](https://github.com/asteris-llc/converge/pull/97) ([BrianHicks](https://github.com/BrianHicks))
- Plan in new preparers [\#95](https://github.com/asteris-llc/converge/pull/95) ([BrianHicks](https://github.com/BrianHicks))
- dependencies should block execution [\#81](https://github.com/asteris-llc/converge/pull/81) ([siddharthist](https://github.com/siddharthist))
- dependency tracker [\#80](https://github.com/asteris-llc/converge/pull/80) ([BrianHicks](https://github.com/BrianHicks))
- Self Serve [\#79](https://github.com/asteris-llc/converge/pull/79) ([BrianHicks](https://github.com/BrianHicks))
- Module server [\#78](https://github.com/asteris-llc/converge/pull/78) ([BrianHicks](https://github.com/BrianHicks))
- File mode module [\#76](https://github.com/asteris-llc/converge/pull/76) ([BrianHicks](https://github.com/BrianHicks))
- Simplify Application [\#73](https://github.com/asteris-llc/converge/pull/73) ([BrianHicks](https://github.com/BrianHicks))
- Fix graph with dashes [\#71](https://github.com/asteris-llc/converge/pull/71) ([BrianHicks](https://github.com/BrianHicks))
- load: from http [\#23](https://github.com/asteris-llc/converge/pull/23) ([siddharthist](https://github.com/siddharthist))

## [0.1.1](https://github.com/asteris-llc/converge/tree/0.1.1) (2016-06-13)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.1...0.1.1)

### Enhancements

- log lines from DAG walking should not appear [\#67](https://github.com/asteris-llc/converge/issues/67)

### Bugs

- color is not working? [\#64](https://github.com/asteris-llc/converge/issues/64)
- recursive modules are valid [\#63](https://github.com/asteris-llc/converge/issues/63)

### Closed Pull Requests

- Feature/logging [\#70](https://github.com/asteris-llc/converge/pull/70) ([BrianHicks](https://github.com/BrianHicks))
- add color printing again [\#69](https://github.com/asteris-llc/converge/pull/69) ([siddharthist](https://github.com/siddharthist))
- Feature/basic blackbox testing [\#68](https://github.com/asteris-llc/converge/pull/68) ([BrianHicks](https://github.com/BrianHicks))

## [0.1](https://github.com/asteris-llc/converge/tree/0.1) (2016-06-09)
[Full Changelog](https://github.com/asteris-llc/converge/compare/Unreleased...0.1)

### Enhancements

- converge format [\#55](https://github.com/asteris-llc/converge/issues/55)
- colorized summary lines [\#45](https://github.com/asteris-llc/converge/issues/45)
- print a nice summary of changes \(or no changes\) after commands [\#35](https://github.com/asteris-llc/converge/issues/35)
- Progressive output in check/apply [\#34](https://github.com/asteris-llc/converge/issues/34)
- Supplying params is hard [\#33](https://github.com/asteris-llc/converge/issues/33)
- Template should accept file permissions [\#31](https://github.com/asteris-llc/converge/issues/31)
- add net/context to Check [\#30](https://github.com/asteris-llc/converge/issues/30)
- templating in param defaults [\#27](https://github.com/asteris-llc/converge/issues/27)
- passing in parameters to top-level module [\#26](https://github.com/asteris-llc/converge/issues/26)
- Apply [\#15](https://github.com/asteris-llc/converge/issues/15)
- HTTPS fetching [\#14](https://github.com/asteris-llc/converge/issues/14)
- Nicer plan output [\#12](https://github.com/asteris-llc/converge/issues/12)
- Checker [\#11](https://github.com/asteris-llc/converge/issues/11)
- Requirements [\#10](https://github.com/asteris-llc/converge/issues/10)

### Bugs

- duplicate names are lost in the graph [\#56](https://github.com/asteris-llc/converge/issues/56)
- fill in short and long descriptions [\#51](https://github.com/asteris-llc/converge/issues/51)
- don't ignore error when Viper binds pflags [\#48](https://github.com/asteris-llc/converge/issues/48)
- param names should have limited characters [\#41](https://github.com/asteris-llc/converge/issues/41)
- template permissions should be 600 by default [\#38](https://github.com/asteris-llc/converge/issues/38)
- exec.Check tests [\#21](https://github.com/asteris-llc/converge/issues/21)
- Incomplete comment  [\#5](https://github.com/asteris-llc/converge/issues/5)

### Closed Issues

- README is out of date [\#52](https://github.com/asteris-llc/converge/issues/52)

### Closed Pull Requests

- Update 0.1 docs [\#60](https://github.com/asteris-llc/converge/pull/60) ([BrianHicks](https://github.com/BrianHicks))
- fixed souceFile.hcl [\#59](https://github.com/asteris-llc/converge/pull/59) ([Dacode45](https://github.com/Dacode45))
- Bug/names [\#58](https://github.com/asteris-llc/converge/pull/58) ([Dacode45](https://github.com/Dacode45))
- fmt [\#57](https://github.com/asteris-llc/converge/pull/57) ([BrianHicks](https://github.com/BrianHicks))
- Feature/requirements [\#54](https://github.com/asteris-llc/converge/pull/54) ([Dacode45](https://github.com/Dacode45))
- don't ignore the error when Viper binds to PFlags [\#50](https://github.com/asteris-llc/converge/pull/50) ([siddharthist](https://github.com/siddharthist))
- color summaries [\#49](https://github.com/asteris-llc/converge/pull/49) ([siddharthist](https://github.com/siddharthist))
- add key=value parameter parsing [\#47](https://github.com/asteris-llc/converge/pull/47) ([siddharthist](https://github.com/siddharthist))
- Summary lines [\#46](https://github.com/asteris-llc/converge/pull/46) ([BrianHicks](https://github.com/BrianHicks))
- render: fix swallowed errors [\#44](https://github.com/asteris-llc/converge/pull/44) ([BrianHicks](https://github.com/BrianHicks))
- limit identifiers to word chars, dashes, and dots [\#43](https://github.com/asteris-llc/converge/pull/43) ([BrianHicks](https://github.com/BrianHicks))
- resource\(template\): fix permissions [\#42](https://github.com/asteris-llc/converge/pull/42) ([BrianHicks](https://github.com/BrianHicks))
- Lazily load parameters - defaults can be templated [\#37](https://github.com/asteris-llc/converge/pull/37) ([siddharthist](https://github.com/siddharthist))
- commands: add streaming output [\#36](https://github.com/asteris-llc/converge/pull/36) ([BrianHicks](https://github.com/BrianHicks))
- Apply [\#32](https://github.com/asteris-llc/converge/pull/32) ([BrianHicks](https://github.com/BrianHicks))
- load: accept initial parameters [\#28](https://github.com/asteris-llc/converge/pull/28) ([BrianHicks](https://github.com/BrianHicks))
- Add tests for Plan [\#25](https://github.com/asteris-llc/converge/pull/25) ([BrianHicks](https://github.com/BrianHicks))
- plan: pretty-printing terminal output [\#22](https://github.com/asteris-llc/converge/pull/22) ([siddharthist](https://github.com/siddharthist))
- fix incomplete comment [\#19](https://github.com/asteris-llc/converge/pull/19) ([BrianHicks](https://github.com/BrianHicks))
- Checker [\#9](https://github.com/asteris-llc/converge/pull/9) ([BrianHicks](https://github.com/BrianHicks))
- Name\(\) -\> String\(\) [\#8](https://github.com/asteris-llc/converge/pull/8) ([siddharthist](https://github.com/siddharthist))
- \[WIP\] Graphviz Support [\#7](https://github.com/asteris-llc/converge/pull/7) ([BrianHicks](https://github.com/BrianHicks))
- Feature/validation error type [\#3](https://github.com/asteris-llc/converge/pull/3) ([siddharthist](https://github.com/siddharthist))
- shelltask: add 'validate' method [\#2](https://github.com/asteris-llc/converge/pull/2) ([siddharthist](https://github.com/siddharthist))
- readme: status -\> check [\#1](https://github.com/asteris-llc/converge/pull/1) ([siddharthist](https://github.com/siddharthist))

# Change Log

## [Unreleased](https://github.com/asteris-llc/converge/tree/HEAD)

[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0...HEAD)

### Enhancements

- Handle time.Time in preparer [\#490](https://github.com/asteris-llc/converge/issues/490)
- k8s demo [\#480](https://github.com/asteris-llc/converge/issues/480)
- no mo' spaces [\#320](https://github.com/asteris-llc/converge/issues/320)
- We're using too many ports [\#286](https://github.com/asteris-llc/converge/issues/286)
- better status display on plan and execution [\#269](https://github.com/asteris-llc/converge/issues/269)
- builtin file owner module [\#74](https://github.com/asteris-llc/converge/issues/74)
- Add interceptors for RPC auth [\#531](https://github.com/asteris-llc/converge/pull/531) ([BrianHicks](https://github.com/BrianHicks))
- demo: kubernetes - use latest version of converge [\#524](https://github.com/asteris-llc/converge/pull/524) ([ryane](https://github.com/ryane))
- fix shasum filename [\#522](https://github.com/asteris-llc/converge/pull/522) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/improve output [\#516](https://github.com/asteris-llc/converge/pull/516) ([arichardet](https://github.com/arichardet))

### Bugs

- LVM FS module reports failure due to still having changes [\#555](https://github.com/asteris-llc/converge/issues/555)
- handle newlines in docs generation for exported fields [\#548](https://github.com/asteris-llc/converge/issues/548)
- docker.container panics if image not present [\#538](https://github.com/asteris-llc/converge/issues/538)
- Return error instead of message from docker resource plan [\#529](https://github.com/asteris-llc/converge/issues/529)
- misleading error message in converge graph command [\#520](https://github.com/asteris-llc/converge/issues/520)
- exit code is 0 when converge plan has errors [\#492](https://github.com/asteris-llc/converge/issues/492)
- error specifying file mode [\#487](https://github.com/asteris-llc/converge/issues/487)
- docker.network misleading error messages [\#485](https://github.com/asteris-llc/converge/issues/485)
- inconsistent indentation on multi-line diffs [\#478](https://github.com/asteris-llc/converge/issues/478)
- pass typed params through to modules [\#409](https://github.com/asteris-llc/converge/issues/409)
- auth code is repeated [\#273](https://github.com/asteris-llc/converge/issues/273)
- Remove indent for multi-line diffs [\#530](https://github.com/asteris-llc/converge/pull/530) ([arichardet](https://github.com/arichardet))
- status: add MayChange status and Warning display [\#528](https://github.com/asteris-llc/converge/pull/528) ([ryane](https://github.com/ryane))
- Fix/disable docker solaris [\#527](https://github.com/asteris-llc/converge/pull/527) ([ryane](https://github.com/ryane))
- plan cmd: return non-zero exit code when plan has error\(s\) [\#525](https://github.com/asteris-llc/converge/pull/525) ([ryane](https://github.com/ryane))
- graph cmd: fix wrong argument length error [\#521](https://github.com/asteris-llc/converge/pull/521) ([ryane](https://github.com/ryane))
- Fix/file.mode errors [\#519](https://github.com/asteris-llc/converge/pull/519) ([ryane](https://github.com/ryane))
- docker: update docker dependency [\#512](https://github.com/asteris-llc/converge/pull/512) ([ryane](https://github.com/ryane))

### Closed Issues

- Better output for cascading failures [\#367](https://github.com/asteris-llc/converge/issues/367)
- file download/fetch module [\#250](https://github.com/asteris-llc/converge/issues/250)

### Closed Pull Requests

- Minor Fixes After Makefile rewrite [\#558](https://github.com/asteris-llc/converge/pull/558) ([BrianHicks](https://github.com/BrianHicks))
- reset updates needed after apply in lvm [\#557](https://github.com/asteris-llc/converge/pull/557) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Makefile: remove ref to VERSION in PACKAGE\_VERSION [\#556](https://github.com/asteris-llc/converge/pull/556) ([BrianHicks](https://github.com/BrianHicks))
- Better Makefiles [\#553](https://github.com/asteris-llc/converge/pull/553) ([BrianHicks](https://github.com/BrianHicks))
- when apply fails due to changes still being present; use those diffs … [\#551](https://github.com/asteris-llc/converge/pull/551) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docs: use same formatting for \(re\)exported fields as parameters [\#550](https://github.com/asteris-llc/converge/pull/550) ([arichardet](https://github.com/arichardet))
- Feature/74 file owner [\#549](https://github.com/asteris-llc/converge/pull/549) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/file.fetch [\#543](https://github.com/asteris-llc/converge/pull/543) ([arichardet](https://github.com/arichardet))
- Fix/538 docker container panic [\#542](https://github.com/asteris-llc/converge/pull/542) ([rebeccaskinner](https://github.com/rebeccaskinner))
- return a better error for planned network failures [\#541](https://github.com/asteris-llc/converge/pull/541) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Timer status display [\#537](https://github.com/asteris-llc/converge/pull/537) ([BrianHicks](https://github.com/BrianHicks))
- Add fuzzing and benchmarks to CI [\#536](https://github.com/asteris-llc/converge/pull/536) ([BrianHicks](https://github.com/BrianHicks))
- disallow some characters as resource names [\#535](https://github.com/asteris-llc/converge/pull/535) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use a single port [\#534](https://github.com/asteris-llc/converge/pull/534) ([BrianHicks](https://github.com/BrianHicks))
- Fix/374 immutable status [\#533](https://github.com/asteris-llc/converge/pull/533) ([rebeccaskinner](https://github.com/rebeccaskinner))
- add logo to Readme [\#532](https://github.com/asteris-llc/converge/pull/532) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docs [\#526](https://github.com/asteris-llc/converge/pull/526) ([arichardet](https://github.com/arichardet))
- Feature/preparer time.time [\#523](https://github.com/asteris-llc/converge/pull/523) ([arichardet](https://github.com/arichardet))
- update readme with graph and download info [\#518](https://github.com/asteris-llc/converge/pull/518) ([stevendborrelli](https://github.com/stevendborrelli))
- kubernetes example [\#507](https://github.com/asteris-llc/converge/pull/507) ([ryane](https://github.com/ryane))

## [0.4.0](https://github.com/asteris-llc/converge/tree/0.4.0) (2016-11-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-rc1...0.4.0)

### Enhancements

- Add `nonempty` tag to preparer [\#486](https://github.com/asteris-llc/converge/issues/486)
- apt module [\#404](https://github.com/asteris-llc/converge/issues/404)
- Handle time.Duration in Preparer [\#358](https://github.com/asteris-llc/converge/issues/358)
- binary distribution [\#306](https://github.com/asteris-llc/converge/issues/306)
- docker.volume resource [\#299](https://github.com/asteris-llc/converge/issues/299)
- docker.network resource [\#298](https://github.com/asteris-llc/converge/issues/298)
- Add ability to modify user [\#278](https://github.com/asteris-llc/converge/issues/278)
- LVM module [\#270](https://github.com/asteris-llc/converge/issues/270)
- update installer to point to 0.4.0 [\#517](https://github.com/asteris-llc/converge/pull/517) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/preparer time.duration [\#483](https://github.com/asteris-llc/converge/pull/483) ([arichardet](https://github.com/arichardet))
- check rpm for empty packages [\#482](https://github.com/asteris-llc/converge/pull/482) ([stevendborrelli](https://github.com/stevendborrelli))
- only extract binary from tarball [\#481](https://github.com/asteris-llc/converge/pull/481) ([stevendborrelli](https://github.com/stevendborrelli))
- docker.network resource [\#477](https://github.com/asteris-llc/converge/pull/477) ([ryane](https://github.com/ryane))
- Feature/apt package [\#461](https://github.com/asteris-llc/converge/pull/461) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docker.volume [\#453](https://github.com/asteris-llc/converge/pull/453) ([ryane](https://github.com/ryane))
- Add the ability to modify a user [\#434](https://github.com/asteris-llc/converge/pull/434) ([arichardet](https://github.com/arichardet))

### Bugs

- validate non-empty string values on some resources [\#337](https://github.com/asteris-llc/converge/issues/337)
- Intermittent test failure with `make test` [\#476](https://github.com/asteris-llc/converge/issues/476)
- Converge loses dependencies between modules within a branch when  [\#427](https://github.com/asteris-llc/converge/issues/427)
- diff output is partially bold [\#402](https://github.com/asteris-llc/converge/issues/402)
- Apply doesn't show diff output [\#399](https://github.com/asteris-llc/converge/issues/399)
- values that don't implement Stringer log badly [\#398](https://github.com/asteris-llc/converge/issues/398)
- resolving conditional macros line [\#389](https://github.com/asteris-llc/converge/issues/389)
- Case predicates fail when they include `lookup` [\#386](https://github.com/asteris-llc/converge/issues/386)
- converge exit code on a failed run [\#365](https://github.com/asteris-llc/converge/issues/365)
- Address name conflicts [\#361](https://github.com/asteris-llc/converge/issues/361)
- resources should have the ability to gracefully stop on interrupt [\#319](https://github.com/asteris-llc/converge/issues/319)
- converge panics when interrupted [\#318](https://github.com/asteris-llc/converge/issues/318)
- Fix/conditional nodes [\#484](https://github.com/asteris-llc/converge/pull/484) ([rebeccaskinner](https://github.com/rebeccaskinner))
- grouping: fix issues with groups and conditionals [\#460](https://github.com/asteris-llc/converge/pull/460) ([ryane](https://github.com/ryane))
- return non-zero exit code if apply graph has an error [\#436](https://github.com/asteris-llc/converge/pull/436) ([ryane](https://github.com/ryane))
- replace "context" with "golang.org/x/net/context" [\#433](https://github.com/asteris-llc/converge/pull/433) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- update changelog, add script [\#511](https://github.com/asteris-llc/converge/pull/511) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: fix link in resource authors guide [\#479](https://github.com/asteris-llc/converge/pull/479) ([ryane](https://github.com/ryane))
- Fix/resource duplication [\#458](https://github.com/asteris-llc/converge/pull/458) ([arichardet](https://github.com/arichardet))
- Fix/graceful exit [\#447](https://github.com/asteris-llc/converge/pull/447) ([arichardet](https://github.com/arichardet))
- LVM module [\#184](https://github.com/asteris-llc/converge/pull/184) ([avnik](https://github.com/avnik))

## [0.4.0-rc1](https://github.com/asteris-llc/converge/tree/0.4.0-rc1) (2016-11-17)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-beta1...0.4.0-rc1)

### Closed Pull Requests

- update changelog to reflect rc1 status [\#510](https://github.com/asteris-llc/converge/pull/510) ([rebeccaskinner](https://github.com/rebeccaskinner))
- move rendering plant creation inside of the transform block to avoid … [\#509](https://github.com/asteris-llc/converge/pull/509) ([rebeccaskinner](https://github.com/rebeccaskinner))
- set github link to the converge repo [\#508](https://github.com/asteris-llc/converge/pull/508) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.4.0-beta1](https://github.com/asteris-llc/converge/tree/0.4.0-beta1) (2016-11-15)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0...0.4.0-beta1)

### Enhancements

- file.owner resource [\#503](https://github.com/asteris-llc/converge/issues/503)

### Bugs

- package.rpm state does not default to present [\#446](https://github.com/asteris-llc/converge/issues/446)
- Error when adding a user when a group of the same name exists [\#425](https://github.com/asteris-llc/converge/issues/425)

### Closed Issues

- nonempty check in preparer for user/group [\#489](https://github.com/asteris-llc/converge/issues/489)

### Closed Pull Requests

- empty commit for CI [\#506](https://github.com/asteris-llc/converge/pull/506) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/0.4.0 beta1 fix [\#505](https://github.com/asteris-llc/converge/pull/505) ([rebeccaskinner](https://github.com/rebeccaskinner))
- release notes [\#497](https://github.com/asteris-llc/converge/pull/497) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use %v for logging [\#496](https://github.com/asteris-llc/converge/pull/496) ([arichardet](https://github.com/arichardet))
- update docs with conditional changes [\#495](https://github.com/asteris-llc/converge/pull/495) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/nonempty tag [\#494](https://github.com/asteris-llc/converge/pull/494) ([arichardet](https://github.com/arichardet))
- Remove extract executable [\#488](https://github.com/asteris-llc/converge/pull/488) ([arichardet](https://github.com/arichardet))
- Manifest / index page [\#464](https://github.com/asteris-llc/converge/pull/464) ([BrianHicks](https://github.com/BrianHicks))
- default to state 'present' for rpm module [\#463](https://github.com/asteris-llc/converge/pull/463) ([rebeccaskinner](https://github.com/rebeccaskinner))
- vendor: upgrade all deps [\#459](https://github.com/asteris-llc/converge/pull/459) ([BrianHicks](https://github.com/BrianHicks))

## [0.3.0](https://github.com/asteris-llc/converge/tree/0.3.0) (2016-10-27)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-rc1...0.3.0)

### Enhancements

- "param is required" error should include param name [\#335](https://github.com/asteris-llc/converge/issues/335)
- Add ability to modify group [\#279](https://github.com/asteris-llc/converge/issues/279)
- add conditionals to module resource [\#268](https://github.com/asteris-llc/converge/issues/268)
- named locks [\#249](https://github.com/asteris-llc/converge/issues/249)
- pretty printed changes should align values [\#244](https://github.com/asteris-llc/converge/issues/244)
- ability to wait for a condition [\#238](https://github.com/asteris-llc/converge/issues/238)
- proper arrays [\#110](https://github.com/asteris-llc/converge/issues/110)
- makefile: add hashes to files [\#440](https://github.com/asteris-llc/converge/pull/440) ([BrianHicks](https://github.com/BrianHicks))
- \[389\] switch log level from Info to Debug [\#428](https://github.com/asteris-llc/converge/pull/428) ([mason-fish](https://github.com/mason-fish))
- Feature/examples 0.3.0 [\#419](https://github.com/asteris-llc/converge/pull/419) ([ryane](https://github.com/ryane))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- named groups [\#392](https://github.com/asteris-llc/converge/pull/392) ([ryane](https://github.com/ryane))
- Don't render module dependencies [\#387](https://github.com/asteris-llc/converge/pull/387) ([BrianHicks](https://github.com/BrianHicks))
- Added docs, renamed packages to package to  for consistency with other modules [\#384](https://github.com/asteris-llc/converge/pull/384) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use status.RaiseLevel in group [\#377](https://github.com/asteris-llc/converge/pull/377) ([arichardet](https://github.com/arichardet))
- Feature/rpm module [\#373](https://github.com/asteris-llc/converge/pull/373) ([rebeccaskinner](https://github.com/rebeccaskinner))
- use a container for CI [\#372](https://github.com/asteris-llc/converge/pull/372) ([BrianHicks](https://github.com/BrianHicks))
- create a metadata envelope for nodes [\#369](https://github.com/asteris-llc/converge/pull/369) ([BrianHicks](https://github.com/BrianHicks))
- cmd/server.go: use errgroup instead of waitgroup [\#363](https://github.com/asteris-llc/converge/pull/363) ([QuentinPerez](https://github.com/QuentinPerez))
- Feature/conditionals [\#362](https://github.com/asteris-llc/converge/pull/362) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/example swarm wait [\#346](https://github.com/asteris-llc/converge/pull/346) ([ryane](https://github.com/ryane))
- Add Compound Parameters [\#340](https://github.com/asteris-llc/converge/pull/340) ([BrianHicks](https://github.com/BrianHicks))
- load overlay module in elk example [\#326](https://github.com/asteris-llc/converge/pull/326) ([feniix](https://github.com/feniix))
- Refactor deserializers [\#321](https://github.com/asteris-llc/converge/pull/321) ([BrianHicks](https://github.com/BrianHicks))
- Use text/tabwriter to align human output [\#317](https://github.com/asteris-llc/converge/pull/317) ([sehqlr](https://github.com/sehqlr))
- Fix/225 pipeline function refactor [\#307](https://github.com/asteris-llc/converge/pull/307) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Bugs

- Looking up a node from a branch that is also a part of a group introduced a deadlock [\#415](https://github.com/asteris-llc/converge/issues/415)
- user/group not activating changes [\#417](https://github.com/asteris-llc/converge/issues/417)
- docker.container reporting incorrect entrypoint diff [\#410](https://github.com/asteris-llc/converge/issues/410)
- converge panics when encountering an empty conditional [\#407](https://github.com/asteris-llc/converge/issues/407)
- Use lists as params in modules is broken [\#397](https://github.com/asteris-llc/converge/issues/397)
- Explicit dependencies fail inside of case statements [\#385](https://github.com/asteris-llc/converge/issues/385)
- Use `package.rpm` not `rpm.package` for rpm module [\#382](https://github.com/asteris-llc/converge/issues/382)
- shell module should not set StatusLevel based on process exit code [\#323](https://github.com/asteris-llc/converge/issues/323)
- StatusLevel not taken into account during graph execution [\#322](https://github.com/asteris-llc/converge/issues/322)
- param dependency fails in `samples/shellContext.hcl` [\#313](https://github.com/asteris-llc/converge/issues/313)
- Execution Engine Ignoring Warning Levels [\#243](https://github.com/asteris-llc/converge/issues/243)
- Make pipeline functions use mult-return instead of Either [\#225](https://github.com/asteris-llc/converge/issues/225)
- Fix 420 go test race [\#424](https://github.com/asteris-llc/converge/pull/424) ([ryane](https://github.com/ryane))
- Replace `AreSiblingIDs` with `AreSiblings` to prevent deadlocking [\#423](https://github.com/asteris-llc/converge/pull/423) ([rebeccaskinner](https://github.com/rebeccaskinner))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- Handle thunked branches [\#403](https://github.com/asteris-llc/converge/pull/403) ([rebeccaskinner](https://github.com/rebeccaskinner))
- don't exclude modules in getNearestAncestor [\#396](https://github.com/asteris-llc/converge/pull/396) ([ryane](https://github.com/ryane))
- Fix/385 dependencies in conditionals [\#391](https://github.com/asteris-llc/converge/pull/391) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix lookup calls to use os/user in SetAddUserOptions [\#390](https://github.com/asteris-llc/converge/pull/390) ([arichardet](https://github.com/arichardet))
- Change `rpm.package` to `package.rpm` [\#383](https://github.com/asteris-llc/converge/pull/383) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Update user status level and errors, add checks in Apply for group [\#375](https://github.com/asteris-llc/converge/pull/375) ([arichardet](https://github.com/arichardet))
- Ability to use pointers as a preparer value [\#339](https://github.com/asteris-llc/converge/pull/339) ([BrianHicks](https://github.com/BrianHicks))
- Status Error Codes [\#333](https://github.com/asteris-llc/converge/pull/333) ([BrianHicks](https://github.com/BrianHicks))
- fix codeclimate yaml [\#328](https://github.com/asteris-llc/converge/pull/328) ([BrianHicks](https://github.com/BrianHicks))

### Closed Issues

- docs for rpm module [\#381](https://github.com/asteris-llc/converge/issues/381)

### Closed Pull Requests

- update for 0.3.0 [\#443](https://github.com/asteris-llc/converge/pull/443) ([stevendborrelli](https://github.com/stevendborrelli))
- prep for 0.3.0-rc1 [\#437](https://github.com/asteris-llc/converge/pull/437) ([stevendborrelli](https://github.com/stevendborrelli))
- update release notes [\#435](https://github.com/asteris-llc/converge/pull/435) ([stevendborrelli](https://github.com/stevendborrelli))
- beta3 release notes [\#431](https://github.com/asteris-llc/converge/pull/431) ([stevendborrelli](https://github.com/stevendborrelli))
- update docs to include currently identified switch statement limitations [\#388](https://github.com/asteris-llc/converge/pull/388) ([rebeccaskinner](https://github.com/rebeccaskinner))
- update vendored dependencies [\#378](https://github.com/asteris-llc/converge/pull/378) ([BrianHicks](https://github.com/BrianHicks))
- 0.3.0 release documentation [\#376](https://github.com/asteris-llc/converge/pull/376) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: add multi-version control [\#371](https://github.com/asteris-llc/converge/pull/371) ([BrianHicks](https://github.com/BrianHicks))
- add docs for the platform module [\#342](https://github.com/asteris-llc/converge/pull/342) ([stevendborrelli](https://github.com/stevendborrelli))
- update for PR \#307 [\#316](https://github.com/asteris-llc/converge/pull/316) ([stevendborrelli](https://github.com/stevendborrelli))

## [0.3.0-rc1](https://github.com/asteris-llc/converge/tree/0.3.0-rc1) (2016-10-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta3...0.3.0-rc1)

### Closed Issues

- Update documentation regarding conditionals and groups [\#430](https://github.com/asteris-llc/converge/issues/430)

## [0.3.0-beta3](https://github.com/asteris-llc/converge/tree/0.3.0-beta3) (2016-10-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta2...0.3.0-beta3)

### Enhancements

- add installation script [\#416](https://github.com/asteris-llc/converge/pull/416) ([stevendborrelli](https://github.com/stevendborrelli))

### Bugs

- go tests race condition [\#420](https://github.com/asteris-llc/converge/issues/420)

### Closed Pull Requests

- document known bug between conditionals and groups [\#432](https://github.com/asteris-llc/converge/pull/432) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.3.0-beta2](https://github.com/asteris-llc/converge/tree/0.3.0-beta2) (2016-10-24)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta1...0.3.0-beta2)

### Bugs

- Fix container entrypoint diffs [\#412](https://github.com/asteris-llc/converge/pull/412) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- wercker: don't delete removed [\#414](https://github.com/asteris-llc/converge/pull/414) ([BrianHicks](https://github.com/BrianHicks))
- Nightly/test builds [\#413](https://github.com/asteris-llc/converge/pull/413) ([BrianHicks](https://github.com/BrianHicks))
- fix panic [\#408](https://github.com/asteris-llc/converge/pull/408) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.3.0-beta1](https://github.com/asteris-llc/converge/tree/0.3.0-beta1) (2016-10-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/before-moving-resolve-conditionals...0.3.0-beta1)

### Enhancements

- file/dir refactor [\#327](https://github.com/asteris-llc/converge/issues/327)
- Use errgroup in server [\#295](https://github.com/asteris-llc/converge/issues/295)
- Allow ability to indicate group name when adding a user [\#276](https://github.com/asteris-llc/converge/issues/276)
- builtin file group module [\#75](https://github.com/asteris-llc/converge/issues/75)

### Bugs

- Conditional Regression [\#401](https://github.com/asteris-llc/converge/issues/401)
- module dependencies now failing [\#395](https://github.com/asteris-llc/converge/issues/395)
- docker.container regression [\#343](https://github.com/asteris-llc/converge/issues/343)
- samples in the README are formatted incorrectly. [\#104](https://github.com/asteris-llc/converge/issues/104)

### Closed Issues

- Tidy up codebase by fixing `make lint` reports [\#348](https://github.com/asteris-llc/converge/issues/348)
- Can we simplify the visualization of large modules? [\#280](https://github.com/asteris-llc/converge/issues/280)
- How should we manage module docs and examples? [\#173](https://github.com/asteris-llc/converge/issues/173)

### Closed Pull Requests

- Allow lists in parameters when called in a module [\#400](https://github.com/asteris-llc/converge/pull/400) ([BrianHicks](https://github.com/BrianHicks))
- comment typo [\#380](https://github.com/asteris-llc/converge/pull/380) ([rebeccaskinner](https://github.com/rebeccaskinner))
- FIX local command flag added to example [\#370](https://github.com/asteris-llc/converge/pull/370) ([philcryer](https://github.com/philcryer))
- codeclimate: remove markdown linting [\#368](https://github.com/asteris-llc/converge/pull/368) ([BrianHicks](https://github.com/BrianHicks))
- Add testing section to contribution guidelines [\#366](https://github.com/asteris-llc/converge/pull/366) ([arichardet](https://github.com/arichardet))
- remove binary [\#354](https://github.com/asteris-llc/converge/pull/354) ([stevendborrelli](https://github.com/stevendborrelli))
- wait: fix error handling in retrier [\#353](https://github.com/asteris-llc/converge/pull/353) ([ryane](https://github.com/ryane))
- Feature/param name in error [\#352](https://github.com/asteris-llc/converge/pull/352) ([ryane](https://github.com/ryane))
- docs: add section on error values to resource author's guide [\#347](https://github.com/asteris-llc/converge/pull/347) ([BrianHicks](https://github.com/BrianHicks))
- Fix docker regressions [\#345](https://github.com/asteris-llc/converge/pull/345) ([ryane](https://github.com/ryane))
- wait.query and wait.port resources [\#334](https://github.com/asteris-llc/converge/pull/334) ([ryane](https://github.com/ryane))

## [before-moving-resolve-conditionals](https://github.com/asteris-llc/converge/tree/before-moving-resolve-conditionals) (2016-10-05)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0...before-moving-resolve-conditionals)

### Enhancements

- Use github.com/pkg/errors exclusively [\#300](https://github.com/asteris-llc/converge/issues/300)

### Bugs

- handle parameters with valid zero value [\#338](https://github.com/asteris-llc/converge/issues/338)

### Closed Issues

- Universal Benchmarking [\#308](https://github.com/asteris-llc/converge/issues/308)
- Use table tests instead of helpers.PreparerValidator [\#155](https://github.com/asteris-llc/converge/issues/155)

### Closed Pull Requests

- Handle preparer fields where zero is valid [\#344](https://github.com/asteris-llc/converge/pull/344) ([arichardet](https://github.com/arichardet))
- plan: test, which illustrate "return nil, err" problem [\#336](https://github.com/asteris-llc/converge/pull/336) ([avnik](https://github.com/avnik))
- Support for modifying linux groups [\#331](https://github.com/asteris-llc/converge/pull/331) ([arichardet](https://github.com/arichardet))
- Add contributing.md and code of conduct [\#330](https://github.com/asteris-llc/converge/pull/330) ([BrianHicks](https://github.com/BrianHicks))
- shell: non-0 exit code returns StatusWillChange [\#324](https://github.com/asteris-llc/converge/pull/324) ([ryane](https://github.com/ryane))
- map keys are considered "strings" in parse.Node [\#315](https://github.com/asteris-llc/converge/pull/315) ([rebeccaskinner](https://github.com/rebeccaskinner))
- helpers,docker: move AssertDiff to common testhelpers [\#312](https://github.com/asteris-llc/converge/pull/312) ([avnik](https://github.com/avnik))

## [0.2.0](https://github.com/asteris-llc/converge/tree/0.2.0) (2016-09-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-rc1...0.2.0)

### Enhancements

- Calls to `lookup` should be all lower-case [\#223](https://github.com/asteris-llc/converge/issues/223)

### Bugs

- field name conflicts during lookups [\#294](https://github.com/asteris-llc/converge/issues/294)
- panic when trying to lookup against file.directory [\#292](https://github.com/asteris-llc/converge/issues/292)
- Allow adding and deleting a group without specifying gid [\#275](https://github.com/asteris-llc/converge/issues/275)
- param lookup race condition [\#266](https://github.com/asteris-llc/converge/issues/266)
- Node is sometimes unresolvable, depending on ordering within a template [\#254](https://github.com/asteris-llc/converge/issues/254)
- params should handle boolean values [\#251](https://github.com/asteris-llc/converge/issues/251)
- remove token handling logic when only a client and not a server [\#247](https://github.com/asteris-llc/converge/issues/247)
- Check shouldn't be called from Apply [\#240](https://github.com/asteris-llc/converge/issues/240)
- value passing causes dependency resolution failures in child modules [\#239](https://github.com/asteris-llc/converge/issues/239)
- task.Task has extra argument [\#237](https://github.com/asteris-llc/converge/issues/237)
- Lookups fail for file.content [\#236](https://github.com/asteris-llc/converge/issues/236)
- RPC isn't used in healthchecks [\#199](https://github.com/asteris-llc/converge/issues/199)

### Closed Issues

- module author guide [\#291](https://github.com/asteris-llc/converge/issues/291)
- logging is a little too chatty [\#248](https://github.com/asteris-llc/converge/issues/248)
- Add RPC auth documentation [\#246](https://github.com/asteris-llc/converge/issues/246)

### Closed Pull Requests

- Add ability to indicate the group name when adding a user [\#310](https://github.com/asteris-llc/converge/pull/310) ([arichardet](https://github.com/arichardet))
- 0.2.0 release [\#311](https://github.com/asteris-llc/converge/pull/311) ([stevendborrelli](https://github.com/stevendborrelli))
- Add "Basic Usage" section to Server docs [\#287](https://github.com/asteris-llc/converge/pull/287) ([sehqlr](https://github.com/sehqlr))
- Simplify Status implementation [\#284](https://github.com/asteris-llc/converge/pull/284) ([BrianHicks](https://github.com/BrianHicks))
- Fix climate issues [\#282](https://github.com/asteris-llc/converge/pull/282) ([BrianHicks](https://github.com/BrianHicks))
- Add codeclimate checks [\#281](https://github.com/asteris-llc/converge/pull/281) ([BrianHicks](https://github.com/BrianHicks))
- Make the logs quieter [\#274](https://github.com/asteris-llc/converge/pull/274) ([BrianHicks](https://github.com/BrianHicks))
- example: elasticsearch, kibana, and filebeat ~elk [\#272](https://github.com/asteris-llc/converge/pull/272) ([ryane](https://github.com/ryane))
- param: accept boolean values [\#271](https://github.com/asteris-llc/converge/pull/271) ([BrianHicks](https://github.com/BrianHicks))
- Use `MakeLanguage` for dependecy resolution to ensure that no [\#265](https://github.com/asteris-llc/converge/pull/265) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Make notes match sample code [\#262](https://github.com/asteris-llc/converge/pull/262) ([tomduckering](https://github.com/tomduckering))

## [0.2.0-rc1](https://github.com/asteris-llc/converge/tree/0.2.0-rc1) (2016-09-23)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-beta2...0.2.0-rc1)

### Bugs

- docker.image incorrectly reporting changes after apply [\#304](https://github.com/asteris-llc/converge/issues/304)

### Closed Issues

- preprocessor race condition [\#302](https://github.com/asteris-llc/converge/issues/302)

### Closed Pull Requests

- Fixes docker resources and updates examples [\#305](https://github.com/asteris-llc/converge/pull/305) ([ryane](https://github.com/ryane))
- use thread-safe field cache [\#303](https://github.com/asteris-llc/converge/pull/303) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/294 field name conflicts [\#301](https://github.com/asteris-llc/converge/pull/301) ([rebeccaskinner](https://github.com/rebeccaskinner))
- preproc: don't add fields for non-struct anon field [\#293](https://github.com/asteris-llc/converge/pull/293) ([ryane](https://github.com/ryane))
- docs: add draft resource authors guide [\#290](https://github.com/asteris-llc/converge/pull/290) ([BrianHicks](https://github.com/BrianHicks))
- perform healthchecks over RPC [\#289](https://github.com/asteris-llc/converge/pull/289) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/239 [\#288](https://github.com/asteris-llc/converge/pull/288) ([rebeccaskinner](https://github.com/rebeccaskinner))
- cmd: don't set local token if only a client [\#285](https://github.com/asteris-llc/converge/pull/285) ([BrianHicks](https://github.com/BrianHicks))
- Fix/user group - Allow adding/deleting without gid [\#283](https://github.com/asteris-llc/converge/pull/283) ([arichardet](https://github.com/arichardet))

## [0.2.0-beta2](https://github.com/asteris-llc/converge/tree/0.2.0-beta2) (2016-09-20)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0-beta1...0.2.0-beta2)

### Enhancements

- module signing [\#16](https://github.com/asteris-llc/converge/issues/16)

### Bugs

- This graph is rendered kinda funky [\#141](https://github.com/asteris-llc/converge/issues/141)

### Closed Issues

- Update all deps to the latest versions [\#171](https://github.com/asteris-llc/converge/issues/171)
- human printer should be profiled and optimized [\#156](https://github.com/asteris-llc/converge/issues/156)

### Closed Pull Requests

- example: docker swarm mode [\#267](https://github.com/asteris-llc/converge/pull/267) ([ryane](https://github.com/ryane))

## [0.2.0-beta1](https://github.com/asteris-llc/converge/tree/0.2.0-beta1) (2016-09-16)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.1.1...0.2.0-beta1)

### Enhancements

- Create a query module [\#227](https://github.com/asteris-llc/converge/issues/227)
- Graphing should work over RPC [\#207](https://github.com/asteris-llc/converge/issues/207)
- Allow Value-passing between modules [\#206](https://github.com/asteris-llc/converge/issues/206)
- logging needs improvement [\#200](https://github.com/asteris-llc/converge/issues/200)
- REST layer on RPC [\#170](https://github.com/asteris-llc/converge/issues/170)
- CLI RPC client [\#169](https://github.com/asteris-llc/converge/issues/169)
- RPC authentication [\#168](https://github.com/asteris-llc/converge/issues/168)
- RPC server [\#167](https://github.com/asteris-llc/converge/issues/167)
- Docker module [\#159](https://github.com/asteris-llc/converge/issues/159)
- Clear/readable errors [\#140](https://github.com/asteris-llc/converge/issues/140)
- Macros? [\#111](https://github.com/asteris-llc/converge/issues/111)
- should template be called file.content? [\#103](https://github.com/asteris-llc/converge/issues/103)
- execution in new graph code [\#92](https://github.com/asteris-llc/converge/issues/92)
- execution planning in new graph code [\#91](https://github.com/asteris-llc/converge/issues/91)
- implement file.group [\#90](https://github.com/asteris-llc/converge/issues/90)
- reimplement file.user [\#89](https://github.com/asteris-llc/converge/issues/89)
- reimplement file.mode [\#88](https://github.com/asteris-llc/converge/issues/88)
- reimplement templates [\#87](https://github.com/asteris-llc/converge/issues/87)
- reimplement shell tasks [\#86](https://github.com/asteris-llc/converge/issues/86)
- implement graph visualization on top of new graph code [\#85](https://github.com/asteris-llc/converge/issues/85)
- resources should keep track of their parents [\#83](https://github.com/asteris-llc/converge/issues/83)
- infer dependencies based on template contents [\#65](https://github.com/asteris-llc/converge/issues/65)
- "native" filemode module [\#40](https://github.com/asteris-llc/converge/issues/40)
- "native" modules addressing scheme [\#39](https://github.com/asteris-llc/converge/issues/39)
- basic diff finding [\#20](https://github.com/asteris-llc/converge/issues/20)
- authenticated module server [\#18](https://github.com/asteris-llc/converge/issues/18)
- module server [\#17](https://github.com/asteris-llc/converge/issues/17)
- HTTP fetching [\#13](https://github.com/asteris-llc/converge/issues/13)
- graph: keep track of parentage inside edges [\#163](https://github.com/asteris-llc/converge/pull/163) ([BrianHicks](https://github.com/BrianHicks))

### Bugs

- lookup should be fully case insensitive [\#232](https://github.com/asteris-llc/converge/issues/232)
- rpc server does not support graph operation anymore [\#205](https://github.com/asteris-llc/converge/issues/205)
- graph is walked out of order [\#160](https://github.com/asteris-llc/converge/issues/160)
- Concurrent Map Read/Write Panic When Running Tests [\#158](https://github.com/asteris-llc/converge/issues/158)
- Shell Doesn't Validate Against User-Defined Interpreter [\#119](https://github.com/asteris-llc/converge/issues/119)
- invalid script syntax doesn't return a good error [\#113](https://github.com/asteris-llc/converge/issues/113)
- Blackbox/test\_apply\_remote.sh Fails [\#102](https://github.com/asteris-llc/converge/issues/102)
- `helpers` is kind of an awful module name [\#96](https://github.com/asteris-llc/converge/issues/96)
- Failing tasks don't block tasks that depend on them [\#82](https://github.com/asteris-llc/converge/issues/82)
- all modules in samples must be valid and useful [\#66](https://github.com/asteris-llc/converge/issues/66)
- params should not have default dependencies [\#62](https://github.com/asteris-llc/converge/issues/62)
- resource names with dashes break graph rendering [\#61](https://github.com/asteris-llc/converge/issues/61)
- Fix error handling when fail to print results during plan or apply [\#183](https://github.com/asteris-llc/converge/pull/183) ([arichardet](https://github.com/arichardet))
- Return better error information for Shell module failures [\#121](https://github.com/asteris-llc/converge/pull/121) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Closed Issues

- errors from check silently swallowed during apply [\#213](https://github.com/asteris-llc/converge/issues/213)
- errors are printed with full stacks [\#208](https://github.com/asteris-llc/converge/issues/208)
- nil pointer panic when file.content resource errors [\#179](https://github.com/asteris-llc/converge/issues/179)
- intermittent concurrent map write panic with dependent params [\#176](https://github.com/asteris-llc/converge/issues/176)
- use standard library net/context library [\#161](https://github.com/asteris-llc/converge/issues/161)
- Subgraph PrettyPrinter should use `nil` as the `SubgraphBottomID` [\#149](https://github.com/asteris-llc/converge/issues/149)
- DigraphPrettyPrinter \(prettyprinters\) should shrink [\#147](https://github.com/asteris-llc/converge/issues/147)
- Renderable \(prettyprinters\) is larger than necessary [\#146](https://github.com/asteris-llc/converge/issues/146)
- Shell command context [\#145](https://github.com/asteris-llc/converge/issues/145)
- Add a 'Health' Module [\#144](https://github.com/asteris-llc/converge/issues/144)
- Builtin file.directory module [\#137](https://github.com/asteris-llc/converge/issues/137)
- Graphviz Generation Fails With Empty Param [\#129](https://github.com/asteris-llc/converge/issues/129)
- parameters should be required if no default is set [\#123](https://github.com/asteris-llc/converge/issues/123)
- Shell Module Error Codes Could Have Better Formatting [\#120](https://github.com/asteris-llc/converge/issues/120)
- params should allow commas [\#114](https://github.com/asteris-llc/converge/issues/114)
- modules and params should not show up by default [\#106](https://github.com/asteris-llc/converge/issues/106)
- use a context in loading [\#94](https://github.com/asteris-llc/converge/issues/94)
- builtin file owner and file permissions module [\#84](https://github.com/asteris-llc/converge/issues/84)

### Closed Pull Requests

- We got an assigned port so here's docs. [\#261](https://github.com/asteris-llc/converge/pull/261) ([BrianHicks](https://github.com/BrianHicks))
- update readme [\#260](https://github.com/asteris-llc/converge/pull/260) ([stevendborrelli](https://github.com/stevendborrelli))
- Add user support [\#259](https://github.com/asteris-llc/converge/pull/259) ([arichardet](https://github.com/arichardet))
- MOAR DOCS PLS [\#258](https://github.com/asteris-llc/converge/pull/258) ([BrianHicks](https://github.com/BrianHicks))
- Makefile: refine xcompile targets [\#257](https://github.com/asteris-llc/converge/pull/257) ([BrianHicks](https://github.com/BrianHicks))
- Fixed issue where apply swallowed error. [\#256](https://github.com/asteris-llc/converge/pull/256) ([Dacode45](https://github.com/Dacode45))
- add mutex to render [\#255](https://github.com/asteris-llc/converge/pull/255) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Add default group implementation for unsupported systems [\#252](https://github.com/asteris-llc/converge/pull/252) ([arichardet](https://github.com/arichardet))
- Feature/module verification [\#245](https://github.com/asteris-llc/converge/pull/245) ([ajhager](https://github.com/ajhager))
- Feature/fully lowercase field names [\#242](https://github.com/asteris-llc/converge/pull/242) ([rebeccaskinner](https://github.com/rebeccaskinner))
- file.directory resource [\#241](https://github.com/asteris-llc/converge/pull/241) ([BrianHicks](https://github.com/BrianHicks))
- Add user group support [\#234](https://github.com/asteris-llc/converge/pull/234) ([arichardet](https://github.com/arichardet))
- support lowercase field names [\#231](https://github.com/asteris-llc/converge/pull/231) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/query refactor [\#230](https://github.com/asteris-llc/converge/pull/230) ([rebeccaskinner](https://github.com/rebeccaskinner))
- handle params that are thunked [\#229](https://github.com/asteris-llc/converge/pull/229) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/query module [\#228](https://github.com/asteris-llc/converge/pull/228) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Refactor helpers into descriptive names [\#222](https://github.com/asteris-llc/converge/pull/222) ([BrianHicks](https://github.com/BrianHicks))
- Feature/value passing [\#220](https://github.com/asteris-llc/converge/pull/220) ([rebeccaskinner](https://github.com/rebeccaskinner))
- rpc: flatten details field [\#218](https://github.com/asteris-llc/converge/pull/218) ([BrianHicks](https://github.com/BrianHicks))
- Feature/formatter \#208 [\#217](https://github.com/asteris-llc/converge/pull/217) ([BrianHicks](https://github.com/BrianHicks))
- docker.container: fix entrypoint diff logic [\#215](https://github.com/asteris-llc/converge/pull/215) ([ryane](https://github.com/ryane))
- Graphing over RPC [\#214](https://github.com/asteris-llc/converge/pull/214) ([BrianHicks](https://github.com/BrianHicks))
- docker.container resource [\#203](https://github.com/asteris-llc/converge/pull/203) ([ryane](https://github.com/ryane))
- Rework logs to use logrus [\#202](https://github.com/asteris-llc/converge/pull/202) ([BrianHicks](https://github.com/BrianHicks))
- Feature/platform refactor [\#197](https://github.com/asteris-llc/converge/pull/197) ([stevendborrelli](https://github.com/stevendborrelli))
- Added host environment variable support [\#195](https://github.com/asteris-llc/converge/pull/195) ([arichardet](https://github.com/arichardet))
- resources\(docker image\): document [\#193](https://github.com/asteris-llc/converge/pull/193) ([BrianHicks](https://github.com/BrianHicks))
- Documentation Site [\#192](https://github.com/asteris-llc/converge/pull/192) ([BrianHicks](https://github.com/BrianHicks))
- Feature/health module [\#189](https://github.com/asteris-llc/converge/pull/189) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docker.image resource redux [\#188](https://github.com/asteris-llc/converge/pull/188) ([ryane](https://github.com/ryane))
- Basic RPC [\#187](https://github.com/asteris-llc/converge/pull/187) ([BrianHicks](https://github.com/BrianHicks))
- shell env and working dir support [\#185](https://github.com/asteris-llc/converge/pull/185) ([ryane](https://github.com/ryane))
- \[GRAPH\] Fix \#158, add triggering files, add tests [\#181](https://github.com/asteris-llc/converge/pull/181) ([sehqlr](https://github.com/sehqlr))
- don't panic when result status is nil [\#180](https://github.com/asteris-llc/converge/pull/180) ([ryane](https://github.com/ryane))
- Feature/optimize human printer [\#175](https://github.com/asteris-llc/converge/pull/175) ([arichardet](https://github.com/arichardet))
- Shrink binaries with some magical ldflags [\#174](https://github.com/asteris-llc/converge/pull/174) ([BrianHicks](https://github.com/BrianHicks))
- Feature/shell refactor [\#172](https://github.com/asteris-llc/converge/pull/172) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Dependency ordered graph walking [\#166](https://github.com/asteris-llc/converge/pull/166) ([BrianHicks](https://github.com/BrianHicks))
- Migrate context for go1.7, refactor fetch/http.go [\#164](https://github.com/asteris-llc/converge/pull/164) ([sehqlr](https://github.com/sehqlr))
- Feature/resource return refactor [\#157](https://github.com/asteris-llc/converge/pull/157) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Move errors into the tree [\#154](https://github.com/asteris-llc/converge/pull/154) ([BrianHicks](https://github.com/BrianHicks))
- Human-readable pretty printer [\#152](https://github.com/asteris-llc/converge/pull/152) ([BrianHicks](https://github.com/BrianHicks))
- make nil a useful default bottom value for subgraphdi [\#150](https://github.com/asteris-llc/converge/pull/150) ([rebeccaskinner](https://github.com/rebeccaskinner))
- JSONL output and interface refactoring [\#148](https://github.com/asteris-llc/converge/pull/148) ([BrianHicks](https://github.com/BrianHicks))
- Feature/template split [\#136](https://github.com/asteris-llc/converge/pull/136) ([rebeccaskinner](https://github.com/rebeccaskinner))
- fetch: add context [\#135](https://github.com/asteris-llc/converge/pull/135) ([BrianHicks](https://github.com/BrianHicks))
- Render merged subtrees [\#133](https://github.com/asteris-llc/converge/pull/133) ([BrianHicks](https://github.com/BrianHicks))
- Fix/119 [\#132](https://github.com/asteris-llc/converge/pull/132) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Trim Duplicate Subtrees [\#131](https://github.com/asteris-llc/converge/pull/131) ([BrianHicks](https://github.com/BrianHicks))
- Appropriately handle required parameters instead of crashing [\#130](https://github.com/asteris-llc/converge/pull/130) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use concurrent map implementation for graph values [\#128](https://github.com/asteris-llc/converge/pull/128) ([BrianHicks](https://github.com/BrianHicks))
- Make params missing defaults required [\#127](https://github.com/asteris-llc/converge/pull/127) ([BrianHicks](https://github.com/BrianHicks))
- makefile: simplify linting [\#126](https://github.com/asteris-llc/converge/pull/126) ([BrianHicks](https://github.com/BrianHicks))
- Rename "template" to "file.content" [\#124](https://github.com/asteris-llc/converge/pull/124) ([BrianHicks](https://github.com/BrianHicks))
- gometalinter fixes [\#122](https://github.com/asteris-llc/converge/pull/122) ([BrianHicks](https://github.com/BrianHicks))
- Graphviz [\#118](https://github.com/asteris-llc/converge/pull/118) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Rewrite graph system [\#116](https://github.com/asteris-llc/converge/pull/116) ([BrianHicks](https://github.com/BrianHicks))
- resource\(shell\): add interpreter [\#109](https://github.com/asteris-llc/converge/pull/109) ([BrianHicks](https://github.com/BrianHicks))
- switch to glide [\#108](https://github.com/asteris-llc/converge/pull/108) ([BrianHicks](https://github.com/BrianHicks))
- Feature/context [\#107](https://github.com/asteris-llc/converge/pull/107) ([BrianHicks](https://github.com/BrianHicks))
- Format samples in readme [\#105](https://github.com/asteris-llc/converge/pull/105) ([ajhager](https://github.com/ajhager))
- Port shell resource and tests [\#101](https://github.com/asteris-llc/converge/pull/101) ([ajhager](https://github.com/ajhager))
- resource\(template\): implement and test template [\#98](https://github.com/asteris-llc/converge/pull/98) ([BrianHicks](https://github.com/BrianHicks))
- Application for preparers [\#97](https://github.com/asteris-llc/converge/pull/97) ([BrianHicks](https://github.com/BrianHicks))
- Plan in new preparers [\#95](https://github.com/asteris-llc/converge/pull/95) ([BrianHicks](https://github.com/BrianHicks))
- dependencies should block execution [\#81](https://github.com/asteris-llc/converge/pull/81) ([siddharthist](https://github.com/siddharthist))
- dependency tracker [\#80](https://github.com/asteris-llc/converge/pull/80) ([BrianHicks](https://github.com/BrianHicks))
- Self Serve [\#79](https://github.com/asteris-llc/converge/pull/79) ([BrianHicks](https://github.com/BrianHicks))
- Module server [\#78](https://github.com/asteris-llc/converge/pull/78) ([BrianHicks](https://github.com/BrianHicks))
- File mode module [\#76](https://github.com/asteris-llc/converge/pull/76) ([BrianHicks](https://github.com/BrianHicks))
- Simplify Application [\#73](https://github.com/asteris-llc/converge/pull/73) ([BrianHicks](https://github.com/BrianHicks))
- Fix graph with dashes [\#71](https://github.com/asteris-llc/converge/pull/71) ([BrianHicks](https://github.com/BrianHicks))
- load: from http [\#23](https://github.com/asteris-llc/converge/pull/23) ([siddharthist](https://github.com/siddharthist))

## [0.1.1](https://github.com/asteris-llc/converge/tree/0.1.1) (2016-06-13)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.1...0.1.1)

### Enhancements

- log lines from DAG walking should not appear [\#67](https://github.com/asteris-llc/converge/issues/67)

### Bugs

- color is not working? [\#64](https://github.com/asteris-llc/converge/issues/64)
- recursive modules are valid [\#63](https://github.com/asteris-llc/converge/issues/63)

### Closed Pull Requests

- Feature/logging [\#70](https://github.com/asteris-llc/converge/pull/70) ([BrianHicks](https://github.com/BrianHicks))
- add color printing again [\#69](https://github.com/asteris-llc/converge/pull/69) ([siddharthist](https://github.com/siddharthist))
- Feature/basic blackbox testing [\#68](https://github.com/asteris-llc/converge/pull/68) ([BrianHicks](https://github.com/BrianHicks))

## [0.1](https://github.com/asteris-llc/converge/tree/0.1) (2016-06-09)
[Full Changelog](https://github.com/asteris-llc/converge/compare/Unreleased...0.1)

### Enhancements

- converge format [\#55](https://github.com/asteris-llc/converge/issues/55)
- colorized summary lines [\#45](https://github.com/asteris-llc/converge/issues/45)
- print a nice summary of changes \(or no changes\) after commands [\#35](https://github.com/asteris-llc/converge/issues/35)
- Progressive output in check/apply [\#34](https://github.com/asteris-llc/converge/issues/34)
- Supplying params is hard [\#33](https://github.com/asteris-llc/converge/issues/33)
- Template should accept file permissions [\#31](https://github.com/asteris-llc/converge/issues/31)
- add net/context to Check [\#30](https://github.com/asteris-llc/converge/issues/30)
- templating in param defaults [\#27](https://github.com/asteris-llc/converge/issues/27)
- passing in parameters to top-level module [\#26](https://github.com/asteris-llc/converge/issues/26)
- Apply [\#15](https://github.com/asteris-llc/converge/issues/15)
- HTTPS fetching [\#14](https://github.com/asteris-llc/converge/issues/14)
- Nicer plan output [\#12](https://github.com/asteris-llc/converge/issues/12)
- Checker [\#11](https://github.com/asteris-llc/converge/issues/11)
- Requirements [\#10](https://github.com/asteris-llc/converge/issues/10)

### Bugs

- duplicate names are lost in the graph [\#56](https://github.com/asteris-llc/converge/issues/56)
- fill in short and long descriptions [\#51](https://github.com/asteris-llc/converge/issues/51)
- don't ignore error when Viper binds pflags [\#48](https://github.com/asteris-llc/converge/issues/48)
- param names should have limited characters [\#41](https://github.com/asteris-llc/converge/issues/41)
- template permissions should be 600 by default [\#38](https://github.com/asteris-llc/converge/issues/38)
- exec.Check tests [\#21](https://github.com/asteris-llc/converge/issues/21)
- Incomplete comment  [\#5](https://github.com/asteris-llc/converge/issues/5)

### Closed Issues

- README is out of date [\#52](https://github.com/asteris-llc/converge/issues/52)

### Closed Pull Requests

- Update 0.1 docs [\#60](https://github.com/asteris-llc/converge/pull/60) ([BrianHicks](https://github.com/BrianHicks))
- fixed souceFile.hcl [\#59](https://github.com/asteris-llc/converge/pull/59) ([Dacode45](https://github.com/Dacode45))
- Bug/names [\#58](https://github.com/asteris-llc/converge/pull/58) ([Dacode45](https://github.com/Dacode45))
- fmt [\#57](https://github.com/asteris-llc/converge/pull/57) ([BrianHicks](https://github.com/BrianHicks))
- Feature/requirements [\#54](https://github.com/asteris-llc/converge/pull/54) ([Dacode45](https://github.com/Dacode45))
- don't ignore the error when Viper binds to PFlags [\#50](https://github.com/asteris-llc/converge/pull/50) ([siddharthist](https://github.com/siddharthist))
- color summaries [\#49](https://github.com/asteris-llc/converge/pull/49) ([siddharthist](https://github.com/siddharthist))
- add key=value parameter parsing [\#47](https://github.com/asteris-llc/converge/pull/47) ([siddharthist](https://github.com/siddharthist))
- Summary lines [\#46](https://github.com/asteris-llc/converge/pull/46) ([BrianHicks](https://github.com/BrianHicks))
- render: fix swallowed errors [\#44](https://github.com/asteris-llc/converge/pull/44) ([BrianHicks](https://github.com/BrianHicks))
- limit identifiers to word chars, dashes, and dots [\#43](https://github.com/asteris-llc/converge/pull/43) ([BrianHicks](https://github.com/BrianHicks))
- resource\(template\): fix permissions [\#42](https://github.com/asteris-llc/converge/pull/42) ([BrianHicks](https://github.com/BrianHicks))
- Lazily load parameters - defaults can be templated [\#37](https://github.com/asteris-llc/converge/pull/37) ([siddharthist](https://github.com/siddharthist))
- commands: add streaming output [\#36](https://github.com/asteris-llc/converge/pull/36) ([BrianHicks](https://github.com/BrianHicks))
- Apply [\#32](https://github.com/asteris-llc/converge/pull/32) ([BrianHicks](https://github.com/BrianHicks))
- load: accept initial parameters [\#28](https://github.com/asteris-llc/converge/pull/28) ([BrianHicks](https://github.com/BrianHicks))
- Add tests for Plan [\#25](https://github.com/asteris-llc/converge/pull/25) ([BrianHicks](https://github.com/BrianHicks))
- plan: pretty-printing terminal output [\#22](https://github.com/asteris-llc/converge/pull/22) ([siddharthist](https://github.com/siddharthist))
- fix incomplete comment [\#19](https://github.com/asteris-llc/converge/pull/19) ([BrianHicks](https://github.com/BrianHicks))
- Checker [\#9](https://github.com/asteris-llc/converge/pull/9) ([BrianHicks](https://github.com/BrianHicks))
- Name\(\) -\> String\(\) [\#8](https://github.com/asteris-llc/converge/pull/8) ([siddharthist](https://github.com/siddharthist))
- \[WIP\] Graphviz Support [\#7](https://github.com/asteris-llc/converge/pull/7) ([BrianHicks](https://github.com/BrianHicks))
- Feature/validation error type [\#3](https://github.com/asteris-llc/converge/pull/3) ([siddharthist](https://github.com/siddharthist))
- shelltask: add 'validate' method [\#2](https://github.com/asteris-llc/converge/pull/2) ([siddharthist](https://github.com/siddharthist))
- readme: status -\> check [\#1](https://github.com/asteris-llc/converge/pull/1) ([siddharthist](https://github.com/siddharthist))

# Change Log

## [Unreleased](https://github.com/asteris-llc/converge/tree/HEAD)

[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0...HEAD)

### Enhancements

- Handle time.Time in preparer [\#490](https://github.com/asteris-llc/converge/issues/490)
- k8s demo [\#480](https://github.com/asteris-llc/converge/issues/480)
- no mo' spaces [\#320](https://github.com/asteris-llc/converge/issues/320)
- We're using too many ports [\#286](https://github.com/asteris-llc/converge/issues/286)
- better status display on plan and execution [\#269](https://github.com/asteris-llc/converge/issues/269)
- builtin file owner module [\#74](https://github.com/asteris-llc/converge/issues/74)
- Add interceptors for RPC auth [\#531](https://github.com/asteris-llc/converge/pull/531) ([BrianHicks](https://github.com/BrianHicks))
- demo: kubernetes - use latest version of converge [\#524](https://github.com/asteris-llc/converge/pull/524) ([ryane](https://github.com/ryane))
- fix shasum filename [\#522](https://github.com/asteris-llc/converge/pull/522) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/improve output [\#516](https://github.com/asteris-llc/converge/pull/516) ([arichardet](https://github.com/arichardet))

### Bugs

- LVM FS module reports failure due to still having changes [\#555](https://github.com/asteris-llc/converge/issues/555)
- handle newlines in docs generation for exported fields [\#548](https://github.com/asteris-llc/converge/issues/548)
- docker.container panics if image not present [\#538](https://github.com/asteris-llc/converge/issues/538)
- Return error instead of message from docker resource plan [\#529](https://github.com/asteris-llc/converge/issues/529)
- misleading error message in converge graph command [\#520](https://github.com/asteris-llc/converge/issues/520)
- exit code is 0 when converge plan has errors [\#492](https://github.com/asteris-llc/converge/issues/492)
- error specifying file mode [\#487](https://github.com/asteris-llc/converge/issues/487)
- docker.network misleading error messages [\#485](https://github.com/asteris-llc/converge/issues/485)
- inconsistent indentation on multi-line diffs [\#478](https://github.com/asteris-llc/converge/issues/478)
- pass typed params through to modules [\#409](https://github.com/asteris-llc/converge/issues/409)
- auth code is repeated [\#273](https://github.com/asteris-llc/converge/issues/273)
- Remove indent for multi-line diffs [\#530](https://github.com/asteris-llc/converge/pull/530) ([arichardet](https://github.com/arichardet))
- status: add MayChange status and Warning display [\#528](https://github.com/asteris-llc/converge/pull/528) ([ryane](https://github.com/ryane))
- Fix/disable docker solaris [\#527](https://github.com/asteris-llc/converge/pull/527) ([ryane](https://github.com/ryane))
- plan cmd: return non-zero exit code when plan has error\(s\) [\#525](https://github.com/asteris-llc/converge/pull/525) ([ryane](https://github.com/ryane))
- graph cmd: fix wrong argument length error [\#521](https://github.com/asteris-llc/converge/pull/521) ([ryane](https://github.com/ryane))
- Fix/file.mode errors [\#519](https://github.com/asteris-llc/converge/pull/519) ([ryane](https://github.com/ryane))
- docker: update docker dependency [\#512](https://github.com/asteris-llc/converge/pull/512) ([ryane](https://github.com/ryane))

### Closed Issues

- Better output for cascading failures [\#367](https://github.com/asteris-llc/converge/issues/367)
- file download/fetch module [\#250](https://github.com/asteris-llc/converge/issues/250)

### Closed Pull Requests

- reset updates needed after apply in lvm [\#557](https://github.com/asteris-llc/converge/pull/557) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Makefile: remove ref to VERSION in PACKAGE\_VERSION [\#556](https://github.com/asteris-llc/converge/pull/556) ([BrianHicks](https://github.com/BrianHicks))
- Better Makefiles [\#553](https://github.com/asteris-llc/converge/pull/553) ([BrianHicks](https://github.com/BrianHicks))
- when apply fails due to changes still being present; use those diffs … [\#551](https://github.com/asteris-llc/converge/pull/551) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docs: use same formatting for \(re\)exported fields as parameters [\#550](https://github.com/asteris-llc/converge/pull/550) ([arichardet](https://github.com/arichardet))
- Feature/74 file owner [\#549](https://github.com/asteris-llc/converge/pull/549) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/file.fetch [\#543](https://github.com/asteris-llc/converge/pull/543) ([arichardet](https://github.com/arichardet))
- Fix/538 docker container panic [\#542](https://github.com/asteris-llc/converge/pull/542) ([rebeccaskinner](https://github.com/rebeccaskinner))
- return a better error for planned network failures [\#541](https://github.com/asteris-llc/converge/pull/541) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Timer status display [\#537](https://github.com/asteris-llc/converge/pull/537) ([BrianHicks](https://github.com/BrianHicks))
- Add fuzzing and benchmarks to CI [\#536](https://github.com/asteris-llc/converge/pull/536) ([BrianHicks](https://github.com/BrianHicks))
- disallow some characters as resource names [\#535](https://github.com/asteris-llc/converge/pull/535) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use a single port [\#534](https://github.com/asteris-llc/converge/pull/534) ([BrianHicks](https://github.com/BrianHicks))
- Fix/374 immutable status [\#533](https://github.com/asteris-llc/converge/pull/533) ([rebeccaskinner](https://github.com/rebeccaskinner))
- add logo to Readme [\#532](https://github.com/asteris-llc/converge/pull/532) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docs [\#526](https://github.com/asteris-llc/converge/pull/526) ([arichardet](https://github.com/arichardet))
- Feature/preparer time.time [\#523](https://github.com/asteris-llc/converge/pull/523) ([arichardet](https://github.com/arichardet))
- update readme with graph and download info [\#518](https://github.com/asteris-llc/converge/pull/518) ([stevendborrelli](https://github.com/stevendborrelli))
- kubernetes example [\#507](https://github.com/asteris-llc/converge/pull/507) ([ryane](https://github.com/ryane))

## [0.4.0](https://github.com/asteris-llc/converge/tree/0.4.0) (2016-11-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-rc1...0.4.0)

### Enhancements

- Add `nonempty` tag to preparer [\#486](https://github.com/asteris-llc/converge/issues/486)
- apt module [\#404](https://github.com/asteris-llc/converge/issues/404)
- Handle time.Duration in Preparer [\#358](https://github.com/asteris-llc/converge/issues/358)
- binary distribution [\#306](https://github.com/asteris-llc/converge/issues/306)
- docker.volume resource [\#299](https://github.com/asteris-llc/converge/issues/299)
- docker.network resource [\#298](https://github.com/asteris-llc/converge/issues/298)
- Add ability to modify user [\#278](https://github.com/asteris-llc/converge/issues/278)
- LVM module [\#270](https://github.com/asteris-llc/converge/issues/270)
- update installer to point to 0.4.0 [\#517](https://github.com/asteris-llc/converge/pull/517) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/preparer time.duration [\#483](https://github.com/asteris-llc/converge/pull/483) ([arichardet](https://github.com/arichardet))
- check rpm for empty packages [\#482](https://github.com/asteris-llc/converge/pull/482) ([stevendborrelli](https://github.com/stevendborrelli))
- only extract binary from tarball [\#481](https://github.com/asteris-llc/converge/pull/481) ([stevendborrelli](https://github.com/stevendborrelli))
- docker.network resource [\#477](https://github.com/asteris-llc/converge/pull/477) ([ryane](https://github.com/ryane))
- Feature/apt package [\#461](https://github.com/asteris-llc/converge/pull/461) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docker.volume [\#453](https://github.com/asteris-llc/converge/pull/453) ([ryane](https://github.com/ryane))
- Add the ability to modify a user [\#434](https://github.com/asteris-llc/converge/pull/434) ([arichardet](https://github.com/arichardet))

### Bugs

- validate non-empty string values on some resources [\#337](https://github.com/asteris-llc/converge/issues/337)
- Intermittent test failure with `make test` [\#476](https://github.com/asteris-llc/converge/issues/476)
- Converge loses dependencies between modules within a branch when  [\#427](https://github.com/asteris-llc/converge/issues/427)
- diff output is partially bold [\#402](https://github.com/asteris-llc/converge/issues/402)
- Apply doesn't show diff output [\#399](https://github.com/asteris-llc/converge/issues/399)
- values that don't implement Stringer log badly [\#398](https://github.com/asteris-llc/converge/issues/398)
- resolving conditional macros line [\#389](https://github.com/asteris-llc/converge/issues/389)
- Case predicates fail when they include `lookup` [\#386](https://github.com/asteris-llc/converge/issues/386)
- converge exit code on a failed run [\#365](https://github.com/asteris-llc/converge/issues/365)
- Address name conflicts [\#361](https://github.com/asteris-llc/converge/issues/361)
- resources should have the ability to gracefully stop on interrupt [\#319](https://github.com/asteris-llc/converge/issues/319)
- converge panics when interrupted [\#318](https://github.com/asteris-llc/converge/issues/318)
- Fix/conditional nodes [\#484](https://github.com/asteris-llc/converge/pull/484) ([rebeccaskinner](https://github.com/rebeccaskinner))
- grouping: fix issues with groups and conditionals [\#460](https://github.com/asteris-llc/converge/pull/460) ([ryane](https://github.com/ryane))
- return non-zero exit code if apply graph has an error [\#436](https://github.com/asteris-llc/converge/pull/436) ([ryane](https://github.com/ryane))
- replace "context" with "golang.org/x/net/context" [\#433](https://github.com/asteris-llc/converge/pull/433) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- update changelog, add script [\#511](https://github.com/asteris-llc/converge/pull/511) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: fix link in resource authors guide [\#479](https://github.com/asteris-llc/converge/pull/479) ([ryane](https://github.com/ryane))
- Fix/resource duplication [\#458](https://github.com/asteris-llc/converge/pull/458) ([arichardet](https://github.com/arichardet))
- Fix/graceful exit [\#447](https://github.com/asteris-llc/converge/pull/447) ([arichardet](https://github.com/arichardet))
- LVM module [\#184](https://github.com/asteris-llc/converge/pull/184) ([avnik](https://github.com/avnik))

# Change Log

## [0.4.0-rc1](https://github.com/asteris-llc/converge/tree/0.4.0-rc1) (2016-11-17)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.4.0-beta1...0.4.0-rc1)

### Closed Pull Requests

- update changelog to reflect rc1 status [\#510](https://github.com/asteris-llc/converge/pull/510) ([rebeccaskinner](https://github.com/rebeccaskinner))
- move rendering plant creation inside of the transform block to avoid … [\#509](https://github.com/asteris-llc/converge/pull/509) ([rebeccaskinner](https://github.com/rebeccaskinner))
- set github link to the converge repo [\#508](https://github.com/asteris-llc/converge/pull/508) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.4.0-beta1](https://github.com/asteris-llc/converge/tree/0.4.0-beta1) (2016-11-15)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0...0.4.0-beta1)

### Enhancements

- file.owner resource [\#503](https://github.com/asteris-llc/converge/issues/503)
- Add `nonempty` tag to preparer [\#486](https://github.com/asteris-llc/converge/issues/486)
- apt module [\#404](https://github.com/asteris-llc/converge/issues/404)
- Handle time.Duration in Preparer [\#358](https://github.com/asteris-llc/converge/issues/358)
- binary distribution [\#306](https://github.com/asteris-llc/converge/issues/306)
- docker.volume resource [\#299](https://github.com/asteris-llc/converge/issues/299)
- docker.network resource [\#298](https://github.com/asteris-llc/converge/issues/298)
- Add ability to modify user [\#278](https://github.com/asteris-llc/converge/issues/278)
- LVM module [\#270](https://github.com/asteris-llc/converge/issues/270)
- Feature/preparer time.duration [\#483](https://github.com/asteris-llc/converge/pull/483) ([arichardet](https://github.com/arichardet))
- check rpm for empty packages [\#482](https://github.com/asteris-llc/converge/pull/482) ([stevendborrelli](https://github.com/stevendborrelli))
- only extract binary from tarball [\#481](https://github.com/asteris-llc/converge/pull/481) ([stevendborrelli](https://github.com/stevendborrelli))
- docker.network resource [\#477](https://github.com/asteris-llc/converge/pull/477) ([ryane](https://github.com/ryane))
- Feature/apt package [\#461](https://github.com/asteris-llc/converge/pull/461) ([stevendborrelli](https://github.com/stevendborrelli))
- Feature/docker.volume [\#453](https://github.com/asteris-llc/converge/pull/453) ([ryane](https://github.com/ryane))
- Add the ability to modify a user [\#434](https://github.com/asteris-llc/converge/pull/434) ([arichardet](https://github.com/arichardet))

### Bugs

- Intermittent test failure with `make test` [\#476](https://github.com/asteris-llc/converge/issues/476)
- package.rpm state does not default to present [\#446](https://github.com/asteris-llc/converge/issues/446)
- Converge loses dependencies between modules within a branch when  [\#427](https://github.com/asteris-llc/converge/issues/427)
- Error when adding a user when a group of the same name exists [\#425](https://github.com/asteris-llc/converge/issues/425)
- values that don't implement Stringer log badly [\#398](https://github.com/asteris-llc/converge/issues/398)
- Case predicates fail when they include `lookup` [\#386](https://github.com/asteris-llc/converge/issues/386)
- converge exit code on a failed run [\#365](https://github.com/asteris-llc/converge/issues/365)
- Address name conflicts [\#361](https://github.com/asteris-llc/converge/issues/361)
- resources should have the ability to gracefully stop on interrupt [\#319](https://github.com/asteris-llc/converge/issues/319)
- converge panics when interrupted [\#318](https://github.com/asteris-llc/converge/issues/318)
- Fix/conditional nodes [\#484](https://github.com/asteris-llc/converge/pull/484) ([rebeccaskinner](https://github.com/rebeccaskinner))
- grouping: fix issues with groups and conditionals [\#460](https://github.com/asteris-llc/converge/pull/460) ([ryane](https://github.com/ryane))
- return non-zero exit code if apply graph has an error [\#436](https://github.com/asteris-llc/converge/pull/436) ([ryane](https://github.com/ryane))
- replace "context" with "golang.org/x/net/context" [\#433](https://github.com/asteris-llc/converge/pull/433) ([ryane](https://github.com/ryane))

### Closed Issues

- nonempty check in preparer for user/group [\#489](https://github.com/asteris-llc/converge/issues/489)

### Closed Pull Requests

- empty commit for CI [\#506](https://github.com/asteris-llc/converge/pull/506) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/0.4.0 beta1 fix [\#505](https://github.com/asteris-llc/converge/pull/505) ([rebeccaskinner](https://github.com/rebeccaskinner))
- release notes [\#497](https://github.com/asteris-llc/converge/pull/497) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use %v for logging [\#496](https://github.com/asteris-llc/converge/pull/496) ([arichardet](https://github.com/arichardet))
- update docs with conditional changes [\#495](https://github.com/asteris-llc/converge/pull/495) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/nonempty tag [\#494](https://github.com/asteris-llc/converge/pull/494) ([arichardet](https://github.com/arichardet))
- Remove extract executable [\#488](https://github.com/asteris-llc/converge/pull/488) ([arichardet](https://github.com/arichardet))
- docs: fix link in resource authors guide [\#479](https://github.com/asteris-llc/converge/pull/479) ([ryane](https://github.com/ryane))
- Manifest / index page [\#464](https://github.com/asteris-llc/converge/pull/464) ([BrianHicks](https://github.com/BrianHicks))
- default to state 'present' for rpm module [\#463](https://github.com/asteris-llc/converge/pull/463) ([rebeccaskinner](https://github.com/rebeccaskinner))
- vendor: upgrade all deps [\#459](https://github.com/asteris-llc/converge/pull/459) ([BrianHicks](https://github.com/BrianHicks))
- Fix/resource duplication [\#458](https://github.com/asteris-llc/converge/pull/458) ([arichardet](https://github.com/arichardet))
- Fix/graceful exit [\#447](https://github.com/asteris-llc/converge/pull/447) ([arichardet](https://github.com/arichardet))
- LVM module [\#184](https://github.com/asteris-llc/converge/pull/184) ([avnik](https://github.com/avnik))

## [0.3.0](https://github.com/asteris-llc/converge/tree/0.3.0) (2016-10-27)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-rc1...0.3.0)

### Enhancements

- "param is required" error should include param name [\#335](https://github.com/asteris-llc/converge/issues/335)
- Add ability to modify group [\#279](https://github.com/asteris-llc/converge/issues/279)
- add conditionals to module resource [\#268](https://github.com/asteris-llc/converge/issues/268)
- named locks [\#249](https://github.com/asteris-llc/converge/issues/249)
- pretty printed changes should align values [\#244](https://github.com/asteris-llc/converge/issues/244)
- ability to wait for a condition [\#238](https://github.com/asteris-llc/converge/issues/238)
- proper arrays [\#110](https://github.com/asteris-llc/converge/issues/110)
- makefile: add hashes to files [\#440](https://github.com/asteris-llc/converge/pull/440) ([BrianHicks](https://github.com/BrianHicks))
- \[389\] switch log level from Info to Debug [\#428](https://github.com/asteris-llc/converge/pull/428) ([mason-fish](https://github.com/mason-fish))
- Feature/examples 0.3.0 [\#419](https://github.com/asteris-llc/converge/pull/419) ([ryane](https://github.com/ryane))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- named groups [\#392](https://github.com/asteris-llc/converge/pull/392) ([ryane](https://github.com/ryane))
- Don't render module dependencies [\#387](https://github.com/asteris-llc/converge/pull/387) ([BrianHicks](https://github.com/BrianHicks))
- Added docs, renamed packages to package to  for consistency with other modules [\#384](https://github.com/asteris-llc/converge/pull/384) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use status.RaiseLevel in group [\#377](https://github.com/asteris-llc/converge/pull/377) ([arichardet](https://github.com/arichardet))
- Feature/rpm module [\#373](https://github.com/asteris-llc/converge/pull/373) ([rebeccaskinner](https://github.com/rebeccaskinner))
- use a container for CI [\#372](https://github.com/asteris-llc/converge/pull/372) ([BrianHicks](https://github.com/BrianHicks))
- create a metadata envelope for nodes [\#369](https://github.com/asteris-llc/converge/pull/369) ([BrianHicks](https://github.com/BrianHicks))
- cmd/server.go: use errgroup instead of waitgroup [\#363](https://github.com/asteris-llc/converge/pull/363) ([QuentinPerez](https://github.com/QuentinPerez))
- Feature/conditionals [\#362](https://github.com/asteris-llc/converge/pull/362) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/example swarm wait [\#346](https://github.com/asteris-llc/converge/pull/346) ([ryane](https://github.com/ryane))
- Add Compound Parameters [\#340](https://github.com/asteris-llc/converge/pull/340) ([BrianHicks](https://github.com/BrianHicks))
- load overlay module in elk example [\#326](https://github.com/asteris-llc/converge/pull/326) ([feniix](https://github.com/feniix))
- Refactor deserializers [\#321](https://github.com/asteris-llc/converge/pull/321) ([BrianHicks](https://github.com/BrianHicks))
- Use text/tabwriter to align human output [\#317](https://github.com/asteris-llc/converge/pull/317) ([sehqlr](https://github.com/sehqlr))
- Fix/225 pipeline function refactor [\#307](https://github.com/asteris-llc/converge/pull/307) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Bugs

- Looking up a node from a branch that is also a part of a group introduced a deadlock [\#415](https://github.com/asteris-llc/converge/issues/415)
- user/group not activating changes [\#417](https://github.com/asteris-llc/converge/issues/417)
- docker.container reporting incorrect entrypoint diff [\#410](https://github.com/asteris-llc/converge/issues/410)
- converge panics when encountering an empty conditional [\#407](https://github.com/asteris-llc/converge/issues/407)
- Use lists as params in modules is broken [\#397](https://github.com/asteris-llc/converge/issues/397)
- Explicit dependencies fail inside of case statements [\#385](https://github.com/asteris-llc/converge/issues/385)
- Use `package.rpm` not `rpm.package` for rpm module [\#382](https://github.com/asteris-llc/converge/issues/382)
- shell module should not set StatusLevel based on process exit code [\#323](https://github.com/asteris-llc/converge/issues/323)
- StatusLevel not taken into account during graph execution [\#322](https://github.com/asteris-llc/converge/issues/322)
- param dependency fails in `samples/shellContext.hcl` [\#313](https://github.com/asteris-llc/converge/issues/313)
- Execution Engine Ignoring Warning Levels [\#243](https://github.com/asteris-llc/converge/issues/243)
- Make pipeline functions use mult-return instead of Either [\#225](https://github.com/asteris-llc/converge/issues/225)
- Fix 420 go test race [\#424](https://github.com/asteris-llc/converge/pull/424) ([ryane](https://github.com/ryane))
- Replace `AreSiblingIDs` with `AreSiblings` to prevent deadlocking [\#423](https://github.com/asteris-llc/converge/pull/423) ([rebeccaskinner](https://github.com/rebeccaskinner))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- Handle thunked branches [\#403](https://github.com/asteris-llc/converge/pull/403) ([rebeccaskinner](https://github.com/rebeccaskinner))
- don't exclude modules in getNearestAncestor [\#396](https://github.com/asteris-llc/converge/pull/396) ([ryane](https://github.com/ryane))
- Fix/385 dependencies in conditionals [\#391](https://github.com/asteris-llc/converge/pull/391) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix lookup calls to use os/user in SetAddUserOptions [\#390](https://github.com/asteris-llc/converge/pull/390) ([arichardet](https://github.com/arichardet))
- Change `rpm.package` to `package.rpm` [\#383](https://github.com/asteris-llc/converge/pull/383) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Update user status level and errors, add checks in Apply for group [\#375](https://github.com/asteris-llc/converge/pull/375) ([arichardet](https://github.com/arichardet))
- Ability to use pointers as a preparer value [\#339](https://github.com/asteris-llc/converge/pull/339) ([BrianHicks](https://github.com/BrianHicks))
- Status Error Codes [\#333](https://github.com/asteris-llc/converge/pull/333) ([BrianHicks](https://github.com/BrianHicks))
- fix codeclimate yaml [\#328](https://github.com/asteris-llc/converge/pull/328) ([BrianHicks](https://github.com/BrianHicks))

### Closed Issues

- docs for rpm module [\#381](https://github.com/asteris-llc/converge/issues/381)

### Closed Pull Requests

- update for 0.3.0 [\#443](https://github.com/asteris-llc/converge/pull/443) ([stevendborrelli](https://github.com/stevendborrelli))
- prep for 0.3.0-rc1 [\#437](https://github.com/asteris-llc/converge/pull/437) ([stevendborrelli](https://github.com/stevendborrelli))
- update release notes [\#435](https://github.com/asteris-llc/converge/pull/435) ([stevendborrelli](https://github.com/stevendborrelli))
- beta3 release notes [\#431](https://github.com/asteris-llc/converge/pull/431) ([stevendborrelli](https://github.com/stevendborrelli))
- update docs to include currently identified switch statement limitations [\#388](https://github.com/asteris-llc/converge/pull/388) ([rebeccaskinner](https://github.com/rebeccaskinner))
- update vendored dependencies [\#378](https://github.com/asteris-llc/converge/pull/378) ([BrianHicks](https://github.com/BrianHicks))
- 0.3.0 release documentation [\#376](https://github.com/asteris-llc/converge/pull/376) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: add multi-version control [\#371](https://github.com/asteris-llc/converge/pull/371) ([BrianHicks](https://github.com/BrianHicks))
- add docs for the platform module [\#342](https://github.com/asteris-llc/converge/pull/342) ([stevendborrelli](https://github.com/stevendborrelli))
- update for PR \#307 [\#316](https://github.com/asteris-llc/converge/pull/316) ([stevendborrelli](https://github.com/stevendborrelli))

## [0.3.0-rc1](https://github.com/asteris-llc/converge/tree/0.3.0-rc1) (2016-10-26)
[Full Changelog](https://github.com/asteris-llc/converge/compare/before-moving-resolve-conditionals...0.3.0-rc1)

### Enhancements

- file/dir refactor [\#327](https://github.com/asteris-llc/converge/issues/327)
- Use errgroup in server [\#295](https://github.com/asteris-llc/converge/issues/295)
- Allow ability to indicate group name when adding a user [\#276](https://github.com/asteris-llc/converge/issues/276)
- builtin file group module [\#75](https://github.com/asteris-llc/converge/issues/75)
- add installation script [\#416](https://github.com/asteris-llc/converge/pull/416) ([stevendborrelli](https://github.com/stevendborrelli))

### Bugs

- go tests race condition [\#420](https://github.com/asteris-llc/converge/issues/420)
- diff output is partially bold [\#402](https://github.com/asteris-llc/converge/issues/402)
- Conditional Regression [\#401](https://github.com/asteris-llc/converge/issues/401)
- Apply doesn't show diff output [\#399](https://github.com/asteris-llc/converge/issues/399)
- module dependencies now failing [\#395](https://github.com/asteris-llc/converge/issues/395)
- resolving conditional macros line [\#389](https://github.com/asteris-llc/converge/issues/389)
- docker.container regression [\#343](https://github.com/asteris-llc/converge/issues/343)
- samples in the README are formatted incorrectly. [\#104](https://github.com/asteris-llc/converge/issues/104)
- Fix container entrypoint diffs [\#412](https://github.com/asteris-llc/converge/pull/412) ([ryane](https://github.com/ryane))

### Closed Issues

- Update documentation regarding conditionals and groups [\#430](https://github.com/asteris-llc/converge/issues/430)
- Tidy up codebase by fixing `make lint` reports [\#348](https://github.com/asteris-llc/converge/issues/348)
- Can we simplify the visualization of large modules? [\#280](https://github.com/asteris-llc/converge/issues/280)
- How should we manage module docs and examples? [\#173](https://github.com/asteris-llc/converge/issues/173)

### Closed Pull Requests

- document known bug between conditionals and groups [\#432](https://github.com/asteris-llc/converge/pull/432) ([rebeccaskinner](https://github.com/rebeccaskinner))
- wercker: don't delete removed [\#414](https://github.com/asteris-llc/converge/pull/414) ([BrianHicks](https://github.com/BrianHicks))
- Nightly/test builds [\#413](https://github.com/asteris-llc/converge/pull/413) ([BrianHicks](https://github.com/BrianHicks))
- fix panic [\#408](https://github.com/asteris-llc/converge/pull/408) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Allow lists in parameters when called in a module [\#400](https://github.com/asteris-llc/converge/pull/400) ([BrianHicks](https://github.com/BrianHicks))
- comment typo [\#380](https://github.com/asteris-llc/converge/pull/380) ([rebeccaskinner](https://github.com/rebeccaskinner))
- FIX local command flag added to example [\#370](https://github.com/asteris-llc/converge/pull/370) ([philcryer](https://github.com/philcryer))
- codeclimate: remove markdown linting [\#368](https://github.com/asteris-llc/converge/pull/368) ([BrianHicks](https://github.com/BrianHicks))
- Add testing section to contribution guidelines [\#366](https://github.com/asteris-llc/converge/pull/366) ([arichardet](https://github.com/arichardet))
- remove binary [\#354](https://github.com/asteris-llc/converge/pull/354) ([stevendborrelli](https://github.com/stevendborrelli))
- wait: fix error handling in retrier [\#353](https://github.com/asteris-llc/converge/pull/353) ([ryane](https://github.com/ryane))
- Feature/param name in error [\#352](https://github.com/asteris-llc/converge/pull/352) ([ryane](https://github.com/ryane))
- docs: add section on error values to resource author's guide [\#347](https://github.com/asteris-llc/converge/pull/347) ([BrianHicks](https://github.com/BrianHicks))
- Fix docker regressions [\#345](https://github.com/asteris-llc/converge/pull/345) ([ryane](https://github.com/ryane))
- wait.query and wait.port resources [\#334](https://github.com/asteris-llc/converge/pull/334) ([ryane](https://github.com/ryane))

## [before-moving-resolve-conditionals](https://github.com/asteris-llc/converge/tree/before-moving-resolve-conditionals) (2016-10-05)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta3...before-moving-resolve-conditionals)

### Enhancements

- Use github.com/pkg/errors exclusively [\#300](https://github.com/asteris-llc/converge/issues/300)
- Create a query module [\#227](https://github.com/asteris-llc/converge/issues/227)
- Calls to `lookup` should be all lower-case [\#223](https://github.com/asteris-llc/converge/issues/223)
- Graphing should work over RPC [\#207](https://github.com/asteris-llc/converge/issues/207)
- Allow Value-passing between modules [\#206](https://github.com/asteris-llc/converge/issues/206)
- logging needs improvement [\#200](https://github.com/asteris-llc/converge/issues/200)
- REST layer on RPC [\#170](https://github.com/asteris-llc/converge/issues/170)
- CLI RPC client [\#169](https://github.com/asteris-llc/converge/issues/169)
- RPC authentication [\#168](https://github.com/asteris-llc/converge/issues/168)
- RPC server [\#167](https://github.com/asteris-llc/converge/issues/167)
- Docker module [\#159](https://github.com/asteris-llc/converge/issues/159)
- Clear/readable errors [\#140](https://github.com/asteris-llc/converge/issues/140)
- Macros? [\#111](https://github.com/asteris-llc/converge/issues/111)
- should template be called file.content? [\#103](https://github.com/asteris-llc/converge/issues/103)
- execution in new graph code [\#92](https://github.com/asteris-llc/converge/issues/92)
- execution planning in new graph code [\#91](https://github.com/asteris-llc/converge/issues/91)
- implement file.group [\#90](https://github.com/asteris-llc/converge/issues/90)
- reimplement file.user [\#89](https://github.com/asteris-llc/converge/issues/89)
- reimplement file.mode [\#88](https://github.com/asteris-llc/converge/issues/88)
- reimplement templates [\#87](https://github.com/asteris-llc/converge/issues/87)
- reimplement shell tasks [\#86](https://github.com/asteris-llc/converge/issues/86)
- implement graph visualization on top of new graph code [\#85](https://github.com/asteris-llc/converge/issues/85)
- resources should keep track of their parents [\#83](https://github.com/asteris-llc/converge/issues/83)
- log lines from DAG walking should not appear [\#67](https://github.com/asteris-llc/converge/issues/67)
- infer dependencies based on template contents [\#65](https://github.com/asteris-llc/converge/issues/65)
- converge format [\#55](https://github.com/asteris-llc/converge/issues/55)
- colorized summary lines [\#45](https://github.com/asteris-llc/converge/issues/45)
- "native" filemode module [\#40](https://github.com/asteris-llc/converge/issues/40)
- "native" modules addressing scheme [\#39](https://github.com/asteris-llc/converge/issues/39)
- print a nice summary of changes \(or no changes\) after commands [\#35](https://github.com/asteris-llc/converge/issues/35)
- Progressive output in check/apply [\#34](https://github.com/asteris-llc/converge/issues/34)
- Supplying params is hard [\#33](https://github.com/asteris-llc/converge/issues/33)
- Template should accept file permissions [\#31](https://github.com/asteris-llc/converge/issues/31)
- add net/context to Check [\#30](https://github.com/asteris-llc/converge/issues/30)
- templating in param defaults [\#27](https://github.com/asteris-llc/converge/issues/27)
- passing in parameters to top-level module [\#26](https://github.com/asteris-llc/converge/issues/26)
- basic diff finding [\#20](https://github.com/asteris-llc/converge/issues/20)
- authenticated module server [\#18](https://github.com/asteris-llc/converge/issues/18)
- module server [\#17](https://github.com/asteris-llc/converge/issues/17)
- module signing [\#16](https://github.com/asteris-llc/converge/issues/16)
- Apply [\#15](https://github.com/asteris-llc/converge/issues/15)
- HTTPS fetching [\#14](https://github.com/asteris-llc/converge/issues/14)
- HTTP fetching [\#13](https://github.com/asteris-llc/converge/issues/13)
- Nicer plan output [\#12](https://github.com/asteris-llc/converge/issues/12)
- Checker [\#11](https://github.com/asteris-llc/converge/issues/11)
- Requirements [\#10](https://github.com/asteris-llc/converge/issues/10)
- graph: keep track of parentage inside edges [\#163](https://github.com/asteris-llc/converge/pull/163) ([BrianHicks](https://github.com/BrianHicks))

### Bugs

- handle parameters with valid zero value [\#338](https://github.com/asteris-llc/converge/issues/338)
- docker.image incorrectly reporting changes after apply [\#304](https://github.com/asteris-llc/converge/issues/304)
- field name conflicts during lookups [\#294](https://github.com/asteris-llc/converge/issues/294)
- panic when trying to lookup against file.directory [\#292](https://github.com/asteris-llc/converge/issues/292)
- Allow adding and deleting a group without specifying gid [\#275](https://github.com/asteris-llc/converge/issues/275)
- param lookup race condition [\#266](https://github.com/asteris-llc/converge/issues/266)
- Node is sometimes unresolvable, depending on ordering within a template [\#254](https://github.com/asteris-llc/converge/issues/254)
- params should handle boolean values [\#251](https://github.com/asteris-llc/converge/issues/251)
- remove token handling logic when only a client and not a server [\#247](https://github.com/asteris-llc/converge/issues/247)
- Check shouldn't be called from Apply [\#240](https://github.com/asteris-llc/converge/issues/240)
- value passing causes dependency resolution failures in child modules [\#239](https://github.com/asteris-llc/converge/issues/239)
- task.Task has extra argument [\#237](https://github.com/asteris-llc/converge/issues/237)
- Lookups fail for file.content [\#236](https://github.com/asteris-llc/converge/issues/236)
- lookup should be fully case insensitive [\#232](https://github.com/asteris-llc/converge/issues/232)
- rpc server does not support graph operation anymore [\#205](https://github.com/asteris-llc/converge/issues/205)
- RPC isn't used in healthchecks [\#199](https://github.com/asteris-llc/converge/issues/199)
- graph is walked out of order [\#160](https://github.com/asteris-llc/converge/issues/160)
- Concurrent Map Read/Write Panic When Running Tests [\#158](https://github.com/asteris-llc/converge/issues/158)
- This graph is rendered kinda funky [\#141](https://github.com/asteris-llc/converge/issues/141)
- Shell Doesn't Validate Against User-Defined Interpreter [\#119](https://github.com/asteris-llc/converge/issues/119)
- invalid script syntax doesn't return a good error [\#113](https://github.com/asteris-llc/converge/issues/113)
- Blackbox/test\_apply\_remote.sh Fails [\#102](https://github.com/asteris-llc/converge/issues/102)
- `helpers` is kind of an awful module name [\#96](https://github.com/asteris-llc/converge/issues/96)
- Failing tasks don't block tasks that depend on them [\#82](https://github.com/asteris-llc/converge/issues/82)
- all modules in samples must be valid and useful [\#66](https://github.com/asteris-llc/converge/issues/66)
- color is not working? [\#64](https://github.com/asteris-llc/converge/issues/64)
- recursive modules are valid [\#63](https://github.com/asteris-llc/converge/issues/63)
- params should not have default dependencies [\#62](https://github.com/asteris-llc/converge/issues/62)
- resource names with dashes break graph rendering [\#61](https://github.com/asteris-llc/converge/issues/61)
- duplicate names are lost in the graph [\#56](https://github.com/asteris-llc/converge/issues/56)
- fill in short and long descriptions [\#51](https://github.com/asteris-llc/converge/issues/51)
- don't ignore error when Viper binds pflags [\#48](https://github.com/asteris-llc/converge/issues/48)
- param names should have limited characters [\#41](https://github.com/asteris-llc/converge/issues/41)
- template permissions should be 600 by default [\#38](https://github.com/asteris-llc/converge/issues/38)
- exec.Check tests [\#21](https://github.com/asteris-llc/converge/issues/21)
- Incomplete comment  [\#5](https://github.com/asteris-llc/converge/issues/5)
- Fix error handling when fail to print results during plan or apply [\#183](https://github.com/asteris-llc/converge/pull/183) ([arichardet](https://github.com/arichardet))
- Return better error information for Shell module failures [\#121](https://github.com/asteris-llc/converge/pull/121) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Closed Issues

- Universal Benchmarking [\#308](https://github.com/asteris-llc/converge/issues/308)
- preprocessor race condition [\#302](https://github.com/asteris-llc/converge/issues/302)
- module author guide [\#291](https://github.com/asteris-llc/converge/issues/291)
- logging is a little too chatty [\#248](https://github.com/asteris-llc/converge/issues/248)
- Add RPC auth documentation [\#246](https://github.com/asteris-llc/converge/issues/246)
- errors from check silently swallowed during apply [\#213](https://github.com/asteris-llc/converge/issues/213)
- errors are printed with full stacks [\#208](https://github.com/asteris-llc/converge/issues/208)
- nil pointer panic when file.content resource errors [\#179](https://github.com/asteris-llc/converge/issues/179)
- intermittent concurrent map write panic with dependent params [\#176](https://github.com/asteris-llc/converge/issues/176)
- Update all deps to the latest versions [\#171](https://github.com/asteris-llc/converge/issues/171)
- use standard library net/context library [\#161](https://github.com/asteris-llc/converge/issues/161)
- human printer should be profiled and optimized [\#156](https://github.com/asteris-llc/converge/issues/156)
- Use table tests instead of helpers.PreparerValidator [\#155](https://github.com/asteris-llc/converge/issues/155)
- Subgraph PrettyPrinter should use `nil` as the `SubgraphBottomID` [\#149](https://github.com/asteris-llc/converge/issues/149)
- DigraphPrettyPrinter \(prettyprinters\) should shrink [\#147](https://github.com/asteris-llc/converge/issues/147)
- Renderable \(prettyprinters\) is larger than necessary [\#146](https://github.com/asteris-llc/converge/issues/146)
- Shell command context [\#145](https://github.com/asteris-llc/converge/issues/145)
- Add a 'Health' Module [\#144](https://github.com/asteris-llc/converge/issues/144)
- Builtin file.directory module [\#137](https://github.com/asteris-llc/converge/issues/137)
- Graphviz Generation Fails With Empty Param [\#129](https://github.com/asteris-llc/converge/issues/129)
- parameters should be required if no default is set [\#123](https://github.com/asteris-llc/converge/issues/123)
- Shell Module Error Codes Could Have Better Formatting [\#120](https://github.com/asteris-llc/converge/issues/120)
- params should allow commas [\#114](https://github.com/asteris-llc/converge/issues/114)
- modules and params should not show up by default [\#106](https://github.com/asteris-llc/converge/issues/106)
- use a context in loading [\#94](https://github.com/asteris-llc/converge/issues/94)
- builtin file owner and file permissions module [\#84](https://github.com/asteris-llc/converge/issues/84)
- README is out of date [\#52](https://github.com/asteris-llc/converge/issues/52)

### Closed Pull Requests

- Handle preparer fields where zero is valid [\#344](https://github.com/asteris-llc/converge/pull/344) ([arichardet](https://github.com/arichardet))
- plan: test, which illustrate "return nil, err" problem [\#336](https://github.com/asteris-llc/converge/pull/336) ([avnik](https://github.com/avnik))
- Support for modifying linux groups [\#331](https://github.com/asteris-llc/converge/pull/331) ([arichardet](https://github.com/arichardet))
- Add contributing.md and code of conduct [\#330](https://github.com/asteris-llc/converge/pull/330) ([BrianHicks](https://github.com/BrianHicks))
- shell: non-0 exit code returns StatusWillChange [\#324](https://github.com/asteris-llc/converge/pull/324) ([ryane](https://github.com/ryane))
- map keys are considered "strings" in parse.Node [\#315](https://github.com/asteris-llc/converge/pull/315) ([rebeccaskinner](https://github.com/rebeccaskinner))
- helpers,docker: move AssertDiff to common testhelpers [\#312](https://github.com/asteris-llc/converge/pull/312) ([avnik](https://github.com/avnik))
- 0.2.0 release [\#311](https://github.com/asteris-llc/converge/pull/311) ([stevendborrelli](https://github.com/stevendborrelli))
- Add ability to indicate the group name when adding a user [\#310](https://github.com/asteris-llc/converge/pull/310) ([arichardet](https://github.com/arichardet))
- Fixes docker resources and updates examples [\#305](https://github.com/asteris-llc/converge/pull/305) ([ryane](https://github.com/ryane))
- use thread-safe field cache [\#303](https://github.com/asteris-llc/converge/pull/303) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/294 field name conflicts [\#301](https://github.com/asteris-llc/converge/pull/301) ([rebeccaskinner](https://github.com/rebeccaskinner))
- preproc: don't add fields for non-struct anon field [\#293](https://github.com/asteris-llc/converge/pull/293) ([ryane](https://github.com/ryane))
- docs: add draft resource authors guide [\#290](https://github.com/asteris-llc/converge/pull/290) ([BrianHicks](https://github.com/BrianHicks))
- perform healthchecks over RPC [\#289](https://github.com/asteris-llc/converge/pull/289) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/239 [\#288](https://github.com/asteris-llc/converge/pull/288) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Add "Basic Usage" section to Server docs [\#287](https://github.com/asteris-llc/converge/pull/287) ([sehqlr](https://github.com/sehqlr))
- cmd: don't set local token if only a client [\#285](https://github.com/asteris-llc/converge/pull/285) ([BrianHicks](https://github.com/BrianHicks))
- Simplify Status implementation [\#284](https://github.com/asteris-llc/converge/pull/284) ([BrianHicks](https://github.com/BrianHicks))
- Fix/user group - Allow adding/deleting without gid [\#283](https://github.com/asteris-llc/converge/pull/283) ([arichardet](https://github.com/arichardet))
- Fix climate issues [\#282](https://github.com/asteris-llc/converge/pull/282) ([BrianHicks](https://github.com/BrianHicks))
- Add codeclimate checks [\#281](https://github.com/asteris-llc/converge/pull/281) ([BrianHicks](https://github.com/BrianHicks))
- Make the logs quieter [\#274](https://github.com/asteris-llc/converge/pull/274) ([BrianHicks](https://github.com/BrianHicks))
- example: elasticsearch, kibana, and filebeat ~elk [\#272](https://github.com/asteris-llc/converge/pull/272) ([ryane](https://github.com/ryane))
- param: accept boolean values [\#271](https://github.com/asteris-llc/converge/pull/271) ([BrianHicks](https://github.com/BrianHicks))
- example: docker swarm mode [\#267](https://github.com/asteris-llc/converge/pull/267) ([ryane](https://github.com/ryane))
- Use `MakeLanguage` for dependecy resolution to ensure that no [\#265](https://github.com/asteris-llc/converge/pull/265) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Make notes match sample code [\#262](https://github.com/asteris-llc/converge/pull/262) ([tomduckering](https://github.com/tomduckering))
- We got an assigned port so here's docs. [\#261](https://github.com/asteris-llc/converge/pull/261) ([BrianHicks](https://github.com/BrianHicks))
- update readme [\#260](https://github.com/asteris-llc/converge/pull/260) ([stevendborrelli](https://github.com/stevendborrelli))
- Add user support [\#259](https://github.com/asteris-llc/converge/pull/259) ([arichardet](https://github.com/arichardet))
- MOAR DOCS PLS [\#258](https://github.com/asteris-llc/converge/pull/258) ([BrianHicks](https://github.com/BrianHicks))
- Makefile: refine xcompile targets [\#257](https://github.com/asteris-llc/converge/pull/257) ([BrianHicks](https://github.com/BrianHicks))
- Fixed issue where apply swallowed error. [\#256](https://github.com/asteris-llc/converge/pull/256) ([Dacode45](https://github.com/Dacode45))
- add mutex to render [\#255](https://github.com/asteris-llc/converge/pull/255) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Add default group implementation for unsupported systems [\#252](https://github.com/asteris-llc/converge/pull/252) ([arichardet](https://github.com/arichardet))
- Feature/module verification [\#245](https://github.com/asteris-llc/converge/pull/245) ([ajhager](https://github.com/ajhager))
- Feature/fully lowercase field names [\#242](https://github.com/asteris-llc/converge/pull/242) ([rebeccaskinner](https://github.com/rebeccaskinner))
- file.directory resource [\#241](https://github.com/asteris-llc/converge/pull/241) ([BrianHicks](https://github.com/BrianHicks))
- Add user group support [\#234](https://github.com/asteris-llc/converge/pull/234) ([arichardet](https://github.com/arichardet))
- support lowercase field names [\#231](https://github.com/asteris-llc/converge/pull/231) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/query refactor [\#230](https://github.com/asteris-llc/converge/pull/230) ([rebeccaskinner](https://github.com/rebeccaskinner))
- handle params that are thunked [\#229](https://github.com/asteris-llc/converge/pull/229) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/query module [\#228](https://github.com/asteris-llc/converge/pull/228) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Refactor helpers into descriptive names [\#222](https://github.com/asteris-llc/converge/pull/222) ([BrianHicks](https://github.com/BrianHicks))
- Feature/value passing [\#220](https://github.com/asteris-llc/converge/pull/220) ([rebeccaskinner](https://github.com/rebeccaskinner))
- rpc: flatten details field [\#218](https://github.com/asteris-llc/converge/pull/218) ([BrianHicks](https://github.com/BrianHicks))
- Feature/formatter \#208 [\#217](https://github.com/asteris-llc/converge/pull/217) ([BrianHicks](https://github.com/BrianHicks))
- docker.container: fix entrypoint diff logic [\#215](https://github.com/asteris-llc/converge/pull/215) ([ryane](https://github.com/ryane))
- Graphing over RPC [\#214](https://github.com/asteris-llc/converge/pull/214) ([BrianHicks](https://github.com/BrianHicks))
- docker.container resource [\#203](https://github.com/asteris-llc/converge/pull/203) ([ryane](https://github.com/ryane))
- Rework logs to use logrus [\#202](https://github.com/asteris-llc/converge/pull/202) ([BrianHicks](https://github.com/BrianHicks))
- Feature/platform refactor [\#197](https://github.com/asteris-llc/converge/pull/197) ([stevendborrelli](https://github.com/stevendborrelli))
- Added host environment variable support [\#195](https://github.com/asteris-llc/converge/pull/195) ([arichardet](https://github.com/arichardet))
- resources\(docker image\): document [\#193](https://github.com/asteris-llc/converge/pull/193) ([BrianHicks](https://github.com/BrianHicks))
- Documentation Site [\#192](https://github.com/asteris-llc/converge/pull/192) ([BrianHicks](https://github.com/BrianHicks))
- Feature/health module [\#189](https://github.com/asteris-llc/converge/pull/189) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docker.image resource redux [\#188](https://github.com/asteris-llc/converge/pull/188) ([ryane](https://github.com/ryane))
- Basic RPC [\#187](https://github.com/asteris-llc/converge/pull/187) ([BrianHicks](https://github.com/BrianHicks))
- shell env and working dir support [\#185](https://github.com/asteris-llc/converge/pull/185) ([ryane](https://github.com/ryane))
- \[GRAPH\] Fix \#158, add triggering files, add tests [\#181](https://github.com/asteris-llc/converge/pull/181) ([sehqlr](https://github.com/sehqlr))
- don't panic when result status is nil [\#180](https://github.com/asteris-llc/converge/pull/180) ([ryane](https://github.com/ryane))
- Feature/optimize human printer [\#175](https://github.com/asteris-llc/converge/pull/175) ([arichardet](https://github.com/arichardet))
- Shrink binaries with some magical ldflags [\#174](https://github.com/asteris-llc/converge/pull/174) ([BrianHicks](https://github.com/BrianHicks))
- Feature/shell refactor [\#172](https://github.com/asteris-llc/converge/pull/172) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Dependency ordered graph walking [\#166](https://github.com/asteris-llc/converge/pull/166) ([BrianHicks](https://github.com/BrianHicks))
- Migrate context for go1.7, refactor fetch/http.go [\#164](https://github.com/asteris-llc/converge/pull/164) ([sehqlr](https://github.com/sehqlr))
- Feature/resource return refactor [\#157](https://github.com/asteris-llc/converge/pull/157) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Move errors into the tree [\#154](https://github.com/asteris-llc/converge/pull/154) ([BrianHicks](https://github.com/BrianHicks))
- Human-readable pretty printer [\#152](https://github.com/asteris-llc/converge/pull/152) ([BrianHicks](https://github.com/BrianHicks))
- make nil a useful default bottom value for subgraphdi [\#150](https://github.com/asteris-llc/converge/pull/150) ([rebeccaskinner](https://github.com/rebeccaskinner))
- JSONL output and interface refactoring [\#148](https://github.com/asteris-llc/converge/pull/148) ([BrianHicks](https://github.com/BrianHicks))
- Feature/template split [\#136](https://github.com/asteris-llc/converge/pull/136) ([rebeccaskinner](https://github.com/rebeccaskinner))
- fetch: add context [\#135](https://github.com/asteris-llc/converge/pull/135) ([BrianHicks](https://github.com/BrianHicks))
- Render merged subtrees [\#133](https://github.com/asteris-llc/converge/pull/133) ([BrianHicks](https://github.com/BrianHicks))
- Fix/119 [\#132](https://github.com/asteris-llc/converge/pull/132) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Trim Duplicate Subtrees [\#131](https://github.com/asteris-llc/converge/pull/131) ([BrianHicks](https://github.com/BrianHicks))
- Appropriately handle required parameters instead of crashing [\#130](https://github.com/asteris-llc/converge/pull/130) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use concurrent map implementation for graph values [\#128](https://github.com/asteris-llc/converge/pull/128) ([BrianHicks](https://github.com/BrianHicks))
- Make params missing defaults required [\#127](https://github.com/asteris-llc/converge/pull/127) ([BrianHicks](https://github.com/BrianHicks))
- makefile: simplify linting [\#126](https://github.com/asteris-llc/converge/pull/126) ([BrianHicks](https://github.com/BrianHicks))
- Rename "template" to "file.content" [\#124](https://github.com/asteris-llc/converge/pull/124) ([BrianHicks](https://github.com/BrianHicks))
- gometalinter fixes [\#122](https://github.com/asteris-llc/converge/pull/122) ([BrianHicks](https://github.com/BrianHicks))
- Graphviz [\#118](https://github.com/asteris-llc/converge/pull/118) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Rewrite graph system [\#116](https://github.com/asteris-llc/converge/pull/116) ([BrianHicks](https://github.com/BrianHicks))
- resource\(shell\): add interpreter [\#109](https://github.com/asteris-llc/converge/pull/109) ([BrianHicks](https://github.com/BrianHicks))
- switch to glide [\#108](https://github.com/asteris-llc/converge/pull/108) ([BrianHicks](https://github.com/BrianHicks))
- Feature/context [\#107](https://github.com/asteris-llc/converge/pull/107) ([BrianHicks](https://github.com/BrianHicks))
- Format samples in readme [\#105](https://github.com/asteris-llc/converge/pull/105) ([ajhager](https://github.com/ajhager))
- Port shell resource and tests [\#101](https://github.com/asteris-llc/converge/pull/101) ([ajhager](https://github.com/ajhager))
- resource\(template\): implement and test template [\#98](https://github.com/asteris-llc/converge/pull/98) ([BrianHicks](https://github.com/BrianHicks))
- Application for preparers [\#97](https://github.com/asteris-llc/converge/pull/97) ([BrianHicks](https://github.com/BrianHicks))
- Plan in new preparers [\#95](https://github.com/asteris-llc/converge/pull/95) ([BrianHicks](https://github.com/BrianHicks))
- dependencies should block execution [\#81](https://github.com/asteris-llc/converge/pull/81) ([siddharthist](https://github.com/siddharthist))
- dependency tracker [\#80](https://github.com/asteris-llc/converge/pull/80) ([BrianHicks](https://github.com/BrianHicks))
- Self Serve [\#79](https://github.com/asteris-llc/converge/pull/79) ([BrianHicks](https://github.com/BrianHicks))
- Module server [\#78](https://github.com/asteris-llc/converge/pull/78) ([BrianHicks](https://github.com/BrianHicks))
- File mode module [\#76](https://github.com/asteris-llc/converge/pull/76) ([BrianHicks](https://github.com/BrianHicks))
- Simplify Application [\#73](https://github.com/asteris-llc/converge/pull/73) ([BrianHicks](https://github.com/BrianHicks))
- Fix graph with dashes [\#71](https://github.com/asteris-llc/converge/pull/71) ([BrianHicks](https://github.com/BrianHicks))
- Feature/logging [\#70](https://github.com/asteris-llc/converge/pull/70) ([BrianHicks](https://github.com/BrianHicks))
- add color printing again [\#69](https://github.com/asteris-llc/converge/pull/69) ([siddharthist](https://github.com/siddharthist))
- Feature/basic blackbox testing [\#68](https://github.com/asteris-llc/converge/pull/68) ([BrianHicks](https://github.com/BrianHicks))
- Update 0.1 docs [\#60](https://github.com/asteris-llc/converge/pull/60) ([BrianHicks](https://github.com/BrianHicks))
- fixed souceFile.hcl [\#59](https://github.com/asteris-llc/converge/pull/59) ([Dacode45](https://github.com/Dacode45))
- Bug/names [\#58](https://github.com/asteris-llc/converge/pull/58) ([Dacode45](https://github.com/Dacode45))
- fmt [\#57](https://github.com/asteris-llc/converge/pull/57) ([BrianHicks](https://github.com/BrianHicks))
- Feature/requirements [\#54](https://github.com/asteris-llc/converge/pull/54) ([Dacode45](https://github.com/Dacode45))
- don't ignore the error when Viper binds to PFlags [\#50](https://github.com/asteris-llc/converge/pull/50) ([siddharthist](https://github.com/siddharthist))
- color summaries [\#49](https://github.com/asteris-llc/converge/pull/49) ([siddharthist](https://github.com/siddharthist))
- add key=value parameter parsing [\#47](https://github.com/asteris-llc/converge/pull/47) ([siddharthist](https://github.com/siddharthist))
- Summary lines [\#46](https://github.com/asteris-llc/converge/pull/46) ([BrianHicks](https://github.com/BrianHicks))
- render: fix swallowed errors [\#44](https://github.com/asteris-llc/converge/pull/44) ([BrianHicks](https://github.com/BrianHicks))
- limit identifiers to word chars, dashes, and dots [\#43](https://github.com/asteris-llc/converge/pull/43) ([BrianHicks](https://github.com/BrianHicks))
- resource\(template\): fix permissions [\#42](https://github.com/asteris-llc/converge/pull/42) ([BrianHicks](https://github.com/BrianHicks))
- Lazily load parameters - defaults can be templated [\#37](https://github.com/asteris-llc/converge/pull/37) ([siddharthist](https://github.com/siddharthist))
- commands: add streaming output [\#36](https://github.com/asteris-llc/converge/pull/36) ([BrianHicks](https://github.com/BrianHicks))
- Apply [\#32](https://github.com/asteris-llc/converge/pull/32) ([BrianHicks](https://github.com/BrianHicks))
- load: accept initial parameters [\#28](https://github.com/asteris-llc/converge/pull/28) ([BrianHicks](https://github.com/BrianHicks))
- Add tests for Plan [\#25](https://github.com/asteris-llc/converge/pull/25) ([BrianHicks](https://github.com/BrianHicks))
- load: from http [\#23](https://github.com/asteris-llc/converge/pull/23) ([siddharthist](https://github.com/siddharthist))
- plan: pretty-printing terminal output [\#22](https://github.com/asteris-llc/converge/pull/22) ([siddharthist](https://github.com/siddharthist))
- fix incomplete comment [\#19](https://github.com/asteris-llc/converge/pull/19) ([BrianHicks](https://github.com/BrianHicks))
- Checker [\#9](https://github.com/asteris-llc/converge/pull/9) ([BrianHicks](https://github.com/BrianHicks))
- Name\(\) -\> String\(\) [\#8](https://github.com/asteris-llc/converge/pull/8) ([siddharthist](https://github.com/siddharthist))
- \[WIP\] Graphviz Support [\#7](https://github.com/asteris-llc/converge/pull/7) ([BrianHicks](https://github.com/BrianHicks))
- Feature/validation error type [\#3](https://github.com/asteris-llc/converge/pull/3) ([siddharthist](https://github.com/siddharthist))
- shelltask: add 'validate' method [\#2](https://github.com/asteris-llc/converge/pull/2) ([siddharthist](https://github.com/siddharthist))
- readme: status -\> check [\#1](https://github.com/asteris-llc/converge/pull/1) ([siddharthist](https://github.com/siddharthist))

# Change Log

## [0.3.0-beta3](https://github.com/asteris-llc/converge/tree/0.3.0-beta3) (2016-10-26)

[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta2...0.3.0-beta3)

[Summarized Release Notes](http://converge.aster.is/0.3.0/release-notes/0-3-0/)

### Enhancements

- \[389\] switch log level from Info to Debug [\#428](https://github.com/asteris-llc/converge/pull/428) ([mason-fish](https://github.com/mason-fish))
- Feature/examples 0.3.0 [\#419](https://github.com/asteris-llc/converge/pull/419) ([ryane](https://github.com/ryane))
- use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))
- add installation script [\#416](https://github.com/asteris-llc/converge/pull/416) ([stevendborrelli](https://github.com/stevendborrelli))

### Bugs

- go tests race condition [\#420](https://github.com/asteris-llc/converge/issues/420)
- user/group not activating changes [\#417](https://github.com/asteris-llc/converge/issues/417)
- diff output is partially bold [\#402](https://github.com/asteris-llc/converge/issues/402)
- resolving conditional macros line [\#389](https://github.com/asteris-llc/converge/issues/389)
- Fix 420 go test race [\#424](https://github.com/asteris-llc/converge/pull/424) ([ryane](https://github.com/ryane))
- Replace `AreSiblingIDs` with `AreSiblings` to prevent deadlocking [\#423](https://github.com/asteris-llc/converge/pull/423) ([rebeccaskinner](https://github.com/rebeccaskinner))
-  use ldflags to set version [\#418](https://github.com/asteris-llc/converge/pull/418) ([stevendborrelli](https://github.com/stevendborrelli))

## [0.3.0-beta2](https://github.com/asteris-llc/converge/tree/0.3.0-beta2) (2016-10-24)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.3.0-beta1...0.3.0-beta2)

### Bugs

- docker.container reporting incorrect entrypoint diff [\#410](https://github.com/asteris-llc/converge/issues/410)
- converge panics when encountering an empty conditional [\#407](https://github.com/asteris-llc/converge/issues/407)
- Fix container entrypoint diffs [\#412](https://github.com/asteris-llc/converge/pull/412) ([ryane](https://github.com/ryane))

### Closed Pull Requests

- wercker: don't delete removed [\#414](https://github.com/asteris-llc/converge/pull/414) ([BrianHicks](https://github.com/BrianHicks))
- Nightly/test builds [\#413](https://github.com/asteris-llc/converge/pull/413) ([BrianHicks](https://github.com/BrianHicks))
- fix panic [\#408](https://github.com/asteris-llc/converge/pull/408) ([rebeccaskinner](https://github.com/rebeccaskinner))

## [0.3.0-beta1](https://github.com/asteris-llc/converge/tree/0.3.0-beta1) (2016-10-21)
[Full Changelog](https://github.com/asteris-llc/converge/compare/0.2.0...0.3.0-beta1)

### Enhancements

- "param is required" error should include param name [\#335](https://github.com/asteris-llc/converge/issues/335)
- file/dir refactor [\#327](https://github.com/asteris-llc/converge/issues/327)
- Use github.com/pkg/errors exclusively [\#300](https://github.com/asteris-llc/converge/issues/300)
- Use errgroup in server [\#295](https://github.com/asteris-llc/converge/issues/295)
- Add ability to modify group [\#279](https://github.com/asteris-llc/converge/issues/279)
- Allow ability to indicate group name when adding a user [\#276](https://github.com/asteris-llc/converge/issues/276)
- add conditionals to module resource [\#268](https://github.com/asteris-llc/converge/issues/268)
- named locks [\#249](https://github.com/asteris-llc/converge/issues/249)
- pretty printed changes should align values [\#244](https://github.com/asteris-llc/converge/issues/244)
- ability to wait for a condition [\#238](https://github.com/asteris-llc/converge/issues/238)
- Create a query module [\#227](https://github.com/asteris-llc/converge/issues/227)
- Calls to `lookup` should be all lower-case [\#223](https://github.com/asteris-llc/converge/issues/223)
- Graphing should work over RPC [\#207](https://github.com/asteris-llc/converge/issues/207)
- Allow Value-passing between modules [\#206](https://github.com/asteris-llc/converge/issues/206)
- logging needs improvement [\#200](https://github.com/asteris-llc/converge/issues/200)
- REST layer on RPC [\#170](https://github.com/asteris-llc/converge/issues/170)
- CLI RPC client [\#169](https://github.com/asteris-llc/converge/issues/169)
- RPC authentication [\#168](https://github.com/asteris-llc/converge/issues/168)
- RPC server [\#167](https://github.com/asteris-llc/converge/issues/167)
- Docker module [\#159](https://github.com/asteris-llc/converge/issues/159)
- Clear/readable errors [\#140](https://github.com/asteris-llc/converge/issues/140)
- Macros? [\#111](https://github.com/asteris-llc/converge/issues/111)
- proper arrays [\#110](https://github.com/asteris-llc/converge/issues/110)
- should template be called file.content? [\#103](https://github.com/asteris-llc/converge/issues/103)
- execution in new graph code [\#92](https://github.com/asteris-llc/converge/issues/92)
- execution planning in new graph code [\#91](https://github.com/asteris-llc/converge/issues/91)
- implement file.group [\#90](https://github.com/asteris-llc/converge/issues/90)
- reimplement file.user [\#89](https://github.com/asteris-llc/converge/issues/89)
- reimplement file.mode [\#88](https://github.com/asteris-llc/converge/issues/88)
- reimplement templates [\#87](https://github.com/asteris-llc/converge/issues/87)
- reimplement shell tasks [\#86](https://github.com/asteris-llc/converge/issues/86)
- implement graph visualization on top of new graph code [\#85](https://github.com/asteris-llc/converge/issues/85)
- resources should keep track of their parents [\#83](https://github.com/asteris-llc/converge/issues/83)
- builtin file group module [\#75](https://github.com/asteris-llc/converge/issues/75)
- log lines from DAG walking should not appear [\#67](https://github.com/asteris-llc/converge/issues/67)
- infer dependencies based on template contents [\#65](https://github.com/asteris-llc/converge/issues/65)
- converge format [\#55](https://github.com/asteris-llc/converge/issues/55)
- colorized summary lines [\#45](https://github.com/asteris-llc/converge/issues/45)
- "native" filemode module [\#40](https://github.com/asteris-llc/converge/issues/40)
- "native" modules addressing scheme [\#39](https://github.com/asteris-llc/converge/issues/39)
- print a nice summary of changes \(or no changes\) after commands [\#35](https://github.com/asteris-llc/converge/issues/35)
- Progressive output in check/apply [\#34](https://github.com/asteris-llc/converge/issues/34)
- Supplying params is hard [\#33](https://github.com/asteris-llc/converge/issues/33)
- Template should accept file permissions [\#31](https://github.com/asteris-llc/converge/issues/31)
- add net/context to Check [\#30](https://github.com/asteris-llc/converge/issues/30)
- templating in param defaults [\#27](https://github.com/asteris-llc/converge/issues/27)
- passing in parameters to top-level module [\#26](https://github.com/asteris-llc/converge/issues/26)
- basic diff finding [\#20](https://github.com/asteris-llc/converge/issues/20)
- authenticated module server [\#18](https://github.com/asteris-llc/converge/issues/18)
- module server [\#17](https://github.com/asteris-llc/converge/issues/17)
- module signing [\#16](https://github.com/asteris-llc/converge/issues/16)
- Apply [\#15](https://github.com/asteris-llc/converge/issues/15)
- HTTPS fetching [\#14](https://github.com/asteris-llc/converge/issues/14)
- HTTP fetching [\#13](https://github.com/asteris-llc/converge/issues/13)
- Nicer plan output [\#12](https://github.com/asteris-llc/converge/issues/12)
- Checker [\#11](https://github.com/asteris-llc/converge/issues/11)
- Requirements [\#10](https://github.com/asteris-llc/converge/issues/10)
- named groups [\#392](https://github.com/asteris-llc/converge/pull/392) ([ryane](https://github.com/ryane))
- Don't render module dependencies [\#387](https://github.com/asteris-llc/converge/pull/387) ([BrianHicks](https://github.com/BrianHicks))
- Added docs, renamed packages to package to  for consistency with other modules [\#384](https://github.com/asteris-llc/converge/pull/384) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use status.RaiseLevel in group [\#377](https://github.com/asteris-llc/converge/pull/377) ([arichardet](https://github.com/arichardet))
- Feature/rpm module [\#373](https://github.com/asteris-llc/converge/pull/373) ([rebeccaskinner](https://github.com/rebeccaskinner))
- use a container for CI [\#372](https://github.com/asteris-llc/converge/pull/372) ([BrianHicks](https://github.com/BrianHicks))
- create a metadata envelope for nodes [\#369](https://github.com/asteris-llc/converge/pull/369) ([BrianHicks](https://github.com/BrianHicks))
- cmd/server.go: use errgroup instead of waitgroup [\#363](https://github.com/asteris-llc/converge/pull/363) ([QuentinPerez](https://github.com/QuentinPerez))
- Feature/conditionals [\#362](https://github.com/asteris-llc/converge/pull/362) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/example swarm wait [\#346](https://github.com/asteris-llc/converge/pull/346) ([ryane](https://github.com/ryane))
- Add Compound Parameters [\#340](https://github.com/asteris-llc/converge/pull/340) ([BrianHicks](https://github.com/BrianHicks))
- load overlay module in elk example [\#326](https://github.com/asteris-llc/converge/pull/326) ([feniix](https://github.com/feniix))
- Refactor deserializers [\#321](https://github.com/asteris-llc/converge/pull/321) ([BrianHicks](https://github.com/BrianHicks))
- Use text/tabwriter to align human output [\#317](https://github.com/asteris-llc/converge/pull/317) ([sehqlr](https://github.com/sehqlr))
- Fix/225 pipeline function refactor [\#307](https://github.com/asteris-llc/converge/pull/307) ([rebeccaskinner](https://github.com/rebeccaskinner))
- graph: keep track of parentage inside edges [\#163](https://github.com/asteris-llc/converge/pull/163) ([BrianHicks](https://github.com/BrianHicks))

### Bugs

- Conditional Regression [\#401](https://github.com/asteris-llc/converge/issues/401)
- Apply doesn't show diff output [\#399](https://github.com/asteris-llc/converge/issues/399)
- Use lists as params in modules is broken [\#397](https://github.com/asteris-llc/converge/issues/397)
- module dependencies now failing [\#395](https://github.com/asteris-llc/converge/issues/395)
- Explicit dependencies fail inside of case statements [\#385](https://github.com/asteris-llc/converge/issues/385)
- Use `package.rpm` not `rpm.package` for rpm module [\#382](https://github.com/asteris-llc/converge/issues/382)
- docker.container regression [\#343](https://github.com/asteris-llc/converge/issues/343)
- handle parameters with valid zero value [\#338](https://github.com/asteris-llc/converge/issues/338)
- shell module should not set StatusLevel based on process exit code [\#323](https://github.com/asteris-llc/converge/issues/323)
- StatusLevel not taken into account during graph execution [\#322](https://github.com/asteris-llc/converge/issues/322)
- param dependency fails in `samples/shellContext.hcl` [\#313](https://github.com/asteris-llc/converge/issues/313)
- docker.image incorrectly reporting changes after apply [\#304](https://github.com/asteris-llc/converge/issues/304)
- field name conflicts during lookups [\#294](https://github.com/asteris-llc/converge/issues/294)
- panic when trying to lookup against file.directory [\#292](https://github.com/asteris-llc/converge/issues/292)
- Allow adding and deleting a group without specifying gid [\#275](https://github.com/asteris-llc/converge/issues/275)
- param lookup race condition [\#266](https://github.com/asteris-llc/converge/issues/266)
- Node is sometimes unresolvable, depending on ordering within a template [\#254](https://github.com/asteris-llc/converge/issues/254)
- params should handle boolean values [\#251](https://github.com/asteris-llc/converge/issues/251)
- remove token handling logic when only a client and not a server [\#247](https://github.com/asteris-llc/converge/issues/247)
- Execution Engine Ignoring Warning Levels [\#243](https://github.com/asteris-llc/converge/issues/243)
- Check shouldn't be called from Apply [\#240](https://github.com/asteris-llc/converge/issues/240)
- value passing causes dependency resolution failures in child modules [\#239](https://github.com/asteris-llc/converge/issues/239)
- task.Task has extra argument [\#237](https://github.com/asteris-llc/converge/issues/237)
- Lookups fail for file.content [\#236](https://github.com/asteris-llc/converge/issues/236)
- lookup should be fully case insensitive [\#232](https://github.com/asteris-llc/converge/issues/232)
- Make pipeline functions use mult-return instead of Either [\#225](https://github.com/asteris-llc/converge/issues/225)
- rpc server does not support graph operation anymore [\#205](https://github.com/asteris-llc/converge/issues/205)
- RPC isn't used in healthchecks [\#199](https://github.com/asteris-llc/converge/issues/199)
- graph is walked out of order [\#160](https://github.com/asteris-llc/converge/issues/160)
- Concurrent Map Read/Write Panic When Running Tests [\#158](https://github.com/asteris-llc/converge/issues/158)
- This graph is rendered kinda funky [\#141](https://github.com/asteris-llc/converge/issues/141)
- Shell Doesn't Validate Against User-Defined Interpreter [\#119](https://github.com/asteris-llc/converge/issues/119)
- invalid script syntax doesn't return a good error [\#113](https://github.com/asteris-llc/converge/issues/113)
- samples in the README are formatted incorrectly. [\#104](https://github.com/asteris-llc/converge/issues/104)
- Blackbox/test\_apply\_remote.sh Fails [\#102](https://github.com/asteris-llc/converge/issues/102)
- `helpers` is kind of an awful module name [\#96](https://github.com/asteris-llc/converge/issues/96)
- Failing tasks don't block tasks that depend on them [\#82](https://github.com/asteris-llc/converge/issues/82)
- all modules in samples must be valid and useful [\#66](https://github.com/asteris-llc/converge/issues/66)
- color is not working? [\#64](https://github.com/asteris-llc/converge/issues/64)
- recursive modules are valid [\#63](https://github.com/asteris-llc/converge/issues/63)
- params should not have default dependencies [\#62](https://github.com/asteris-llc/converge/issues/62)
- resource names with dashes break graph rendering [\#61](https://github.com/asteris-llc/converge/issues/61)
- duplicate names are lost in the graph [\#56](https://github.com/asteris-llc/converge/issues/56)
- fill in short and long descriptions [\#51](https://github.com/asteris-llc/converge/issues/51)
- don't ignore error when Viper binds pflags [\#48](https://github.com/asteris-llc/converge/issues/48)
- param names should have limited characters [\#41](https://github.com/asteris-llc/converge/issues/41)
- template permissions should be 600 by default [\#38](https://github.com/asteris-llc/converge/issues/38)
- exec.Check tests [\#21](https://github.com/asteris-llc/converge/issues/21)
- Incomplete comment  [\#5](https://github.com/asteris-llc/converge/issues/5)
- Handle thunked branches [\#403](https://github.com/asteris-llc/converge/pull/403) ([rebeccaskinner](https://github.com/rebeccaskinner))
- don't exclude modules in getNearestAncestor [\#396](https://github.com/asteris-llc/converge/pull/396) ([ryane](https://github.com/ryane))
- Fix/385 dependencies in conditionals [\#391](https://github.com/asteris-llc/converge/pull/391) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix lookup calls to use os/user in SetAddUserOptions [\#390](https://github.com/asteris-llc/converge/pull/390) ([arichardet](https://github.com/arichardet))
- Change `rpm.package` to `package.rpm` [\#383](https://github.com/asteris-llc/converge/pull/383) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Update user status level and errors, add checks in Apply for group [\#375](https://github.com/asteris-llc/converge/pull/375) ([arichardet](https://github.com/arichardet))
- Ability to use pointers as a preparer value [\#339](https://github.com/asteris-llc/converge/pull/339) ([BrianHicks](https://github.com/BrianHicks))
- Status Error Codes [\#333](https://github.com/asteris-llc/converge/pull/333) ([BrianHicks](https://github.com/BrianHicks))
- fix codeclimate yaml [\#328](https://github.com/asteris-llc/converge/pull/328) ([BrianHicks](https://github.com/BrianHicks))
- Fix error handling when fail to print results during plan or apply [\#183](https://github.com/asteris-llc/converge/pull/183) ([arichardet](https://github.com/arichardet))
- Return better error information for Shell module failures [\#121](https://github.com/asteris-llc/converge/pull/121) ([rebeccaskinner](https://github.com/rebeccaskinner))

### Closed Issues

- docs for rpm module [\#381](https://github.com/asteris-llc/converge/issues/381)
- Tidy up codebase by fixing `make lint` reports [\#348](https://github.com/asteris-llc/converge/issues/348)
- Universal Benchmarking [\#308](https://github.com/asteris-llc/converge/issues/308)
- preprocessor race condition [\#302](https://github.com/asteris-llc/converge/issues/302)
- module author guide [\#291](https://github.com/asteris-llc/converge/issues/291)
- Can we simplify the visualization of large modules? [\#280](https://github.com/asteris-llc/converge/issues/280)
- logging is a little too chatty [\#248](https://github.com/asteris-llc/converge/issues/248)
- Add RPC auth documentation [\#246](https://github.com/asteris-llc/converge/issues/246)
- errors from check silently swallowed during apply [\#213](https://github.com/asteris-llc/converge/issues/213)
- errors are printed with full stacks [\#208](https://github.com/asteris-llc/converge/issues/208)
- nil pointer panic when file.content resource errors [\#179](https://github.com/asteris-llc/converge/issues/179)
- intermittent concurrent map write panic with dependent params [\#176](https://github.com/asteris-llc/converge/issues/176)
- How should we manage module docs and examples? [\#173](https://github.com/asteris-llc/converge/issues/173)
- Update all deps to the latest versions [\#171](https://github.com/asteris-llc/converge/issues/171)
- use standard library net/context library [\#161](https://github.com/asteris-llc/converge/issues/161)
- human printer should be profiled and optimized [\#156](https://github.com/asteris-llc/converge/issues/156)
- Use table tests instead of helpers.PreparerValidator [\#155](https://github.com/asteris-llc/converge/issues/155)
- Subgraph PrettyPrinter should use `nil` as the `SubgraphBottomID` [\#149](https://github.com/asteris-llc/converge/issues/149)
- DigraphPrettyPrinter \(prettyprinters\) should shrink [\#147](https://github.com/asteris-llc/converge/issues/147)
- Renderable \(prettyprinters\) is larger than necessary [\#146](https://github.com/asteris-llc/converge/issues/146)
- Shell command context [\#145](https://github.com/asteris-llc/converge/issues/145)
- Add a 'Health' Module [\#144](https://github.com/asteris-llc/converge/issues/144)
- Builtin file.directory module [\#137](https://github.com/asteris-llc/converge/issues/137)
- Graphviz Generation Fails With Empty Param [\#129](https://github.com/asteris-llc/converge/issues/129)
- parameters should be required if no default is set [\#123](https://github.com/asteris-llc/converge/issues/123)
- Shell Module Error Codes Could Have Better Formatting [\#120](https://github.com/asteris-llc/converge/issues/120)
- params should allow commas [\#114](https://github.com/asteris-llc/converge/issues/114)
- modules and params should not show up by default [\#106](https://github.com/asteris-llc/converge/issues/106)
- use a context in loading [\#94](https://github.com/asteris-llc/converge/issues/94)
- builtin file owner and file permissions module [\#84](https://github.com/asteris-llc/converge/issues/84)
- README is out of date [\#52](https://github.com/asteris-llc/converge/issues/52)

### Closed Pull Requests

- Allow lists in parameters when called in a module [\#400](https://github.com/asteris-llc/converge/pull/400) ([BrianHicks](https://github.com/BrianHicks))
- update docs to include currently identified switch statement limitations [\#388](https://github.com/asteris-llc/converge/pull/388) ([rebeccaskinner](https://github.com/rebeccaskinner))
- comment typo [\#380](https://github.com/asteris-llc/converge/pull/380) ([rebeccaskinner](https://github.com/rebeccaskinner))
- update vendored dependencies [\#378](https://github.com/asteris-llc/converge/pull/378) ([BrianHicks](https://github.com/BrianHicks))
- 0.3.0 release documentation [\#376](https://github.com/asteris-llc/converge/pull/376) ([stevendborrelli](https://github.com/stevendborrelli))
- docs: add multi-version control [\#371](https://github.com/asteris-llc/converge/pull/371) ([BrianHicks](https://github.com/BrianHicks))
- FIX local command flag added to example [\#370](https://github.com/asteris-llc/converge/pull/370) ([philcryer](https://github.com/philcryer))
- codeclimate: remove markdown linting [\#368](https://github.com/asteris-llc/converge/pull/368) ([BrianHicks](https://github.com/BrianHicks))
- Add testing section to contribution guidelines [\#366](https://github.com/asteris-llc/converge/pull/366) ([arichardet](https://github.com/arichardet))
- remove binary [\#354](https://github.com/asteris-llc/converge/pull/354) ([stevendborrelli](https://github.com/stevendborrelli))
- wait: fix error handling in retrier [\#353](https://github.com/asteris-llc/converge/pull/353) ([ryane](https://github.com/ryane))
- Feature/param name in error [\#352](https://github.com/asteris-llc/converge/pull/352) ([ryane](https://github.com/ryane))
- docs: add section on error values to resource author's guide [\#347](https://github.com/asteris-llc/converge/pull/347) ([BrianHicks](https://github.com/BrianHicks))
- Fix docker regressions [\#345](https://github.com/asteris-llc/converge/pull/345) ([ryane](https://github.com/ryane))
- Handle preparer fields where zero is valid [\#344](https://github.com/asteris-llc/converge/pull/344) ([arichardet](https://github.com/arichardet))
- add docs for the platform module [\#342](https://github.com/asteris-llc/converge/pull/342) ([stevendborrelli](https://github.com/stevendborrelli))
- plan: test, which illustrate "return nil, err" problem [\#336](https://github.com/asteris-llc/converge/pull/336) ([avnik](https://github.com/avnik))
- wait.query and wait.port resources [\#334](https://github.com/asteris-llc/converge/pull/334) ([ryane](https://github.com/ryane))
- Support for modifying linux groups [\#331](https://github.com/asteris-llc/converge/pull/331) ([arichardet](https://github.com/arichardet))
- Add contributing.md and code of conduct [\#330](https://github.com/asteris-llc/converge/pull/330) ([BrianHicks](https://github.com/BrianHicks))
- shell: non-0 exit code returns StatusWillChange [\#324](https://github.com/asteris-llc/converge/pull/324) ([ryane](https://github.com/ryane))
- update for PR \#307 [\#316](https://github.com/asteris-llc/converge/pull/316) ([stevendborrelli](https://github.com/stevendborrelli))
- map keys are considered "strings" in parse.Node [\#315](https://github.com/asteris-llc/converge/pull/315) ([rebeccaskinner](https://github.com/rebeccaskinner))
- helpers,docker: move AssertDiff to common testhelpers [\#312](https://github.com/asteris-llc/converge/pull/312) ([avnik](https://github.com/avnik))
- 0.2.0 release [\#311](https://github.com/asteris-llc/converge/pull/311) ([stevendborrelli](https://github.com/stevendborrelli))
- Add ability to indicate the group name when adding a user [\#310](https://github.com/asteris-llc/converge/pull/310) ([arichardet](https://github.com/arichardet))
- Fixes docker resources and updates examples [\#305](https://github.com/asteris-llc/converge/pull/305) ([ryane](https://github.com/ryane))
- use thread-safe field cache [\#303](https://github.com/asteris-llc/converge/pull/303) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/294 field name conflicts [\#301](https://github.com/asteris-llc/converge/pull/301) ([rebeccaskinner](https://github.com/rebeccaskinner))
- preproc: don't add fields for non-struct anon field [\#293](https://github.com/asteris-llc/converge/pull/293) ([ryane](https://github.com/ryane))
- docs: add draft resource authors guide [\#290](https://github.com/asteris-llc/converge/pull/290) ([BrianHicks](https://github.com/BrianHicks))
- perform healthchecks over RPC [\#289](https://github.com/asteris-llc/converge/pull/289) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/239 [\#288](https://github.com/asteris-llc/converge/pull/288) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Add "Basic Usage" section to Server docs [\#287](https://github.com/asteris-llc/converge/pull/287) ([sehqlr](https://github.com/sehqlr))
- cmd: don't set local token if only a client [\#285](https://github.com/asteris-llc/converge/pull/285) ([BrianHicks](https://github.com/BrianHicks))
- Simplify Status implementation [\#284](https://github.com/asteris-llc/converge/pull/284) ([BrianHicks](https://github.com/BrianHicks))
- Fix/user group - Allow adding/deleting without gid [\#283](https://github.com/asteris-llc/converge/pull/283) ([arichardet](https://github.com/arichardet))
- Fix climate issues [\#282](https://github.com/asteris-llc/converge/pull/282) ([BrianHicks](https://github.com/BrianHicks))
- Add codeclimate checks [\#281](https://github.com/asteris-llc/converge/pull/281) ([BrianHicks](https://github.com/BrianHicks))
- Make the logs quieter [\#274](https://github.com/asteris-llc/converge/pull/274) ([BrianHicks](https://github.com/BrianHicks))
- example: elasticsearch, kibana, and filebeat ~elk [\#272](https://github.com/asteris-llc/converge/pull/272) ([ryane](https://github.com/ryane))
- param: accept boolean values [\#271](https://github.com/asteris-llc/converge/pull/271) ([BrianHicks](https://github.com/BrianHicks))
- example: docker swarm mode [\#267](https://github.com/asteris-llc/converge/pull/267) ([ryane](https://github.com/ryane))
- Use `MakeLanguage` for dependecy resolution to ensure that no [\#265](https://github.com/asteris-llc/converge/pull/265) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Make notes match sample code [\#262](https://github.com/asteris-llc/converge/pull/262) ([tomduckering](https://github.com/tomduckering))
- We got an assigned port so here's docs. [\#261](https://github.com/asteris-llc/converge/pull/261) ([BrianHicks](https://github.com/BrianHicks))
- update readme [\#260](https://github.com/asteris-llc/converge/pull/260) ([stevendborrelli](https://github.com/stevendborrelli))
- Add user support [\#259](https://github.com/asteris-llc/converge/pull/259) ([arichardet](https://github.com/arichardet))
- MOAR DOCS PLS [\#258](https://github.com/asteris-llc/converge/pull/258) ([BrianHicks](https://github.com/BrianHicks))
- Makefile: refine xcompile targets [\#257](https://github.com/asteris-llc/converge/pull/257) ([BrianHicks](https://github.com/BrianHicks))
- Fixed issue where apply swallowed error. [\#256](https://github.com/asteris-llc/converge/pull/256) ([Dacode45](https://github.com/Dacode45))
- add mutex to render [\#255](https://github.com/asteris-llc/converge/pull/255) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Add default group implementation for unsupported systems [\#252](https://github.com/asteris-llc/converge/pull/252) ([arichardet](https://github.com/arichardet))
- Feature/module verification [\#245](https://github.com/asteris-llc/converge/pull/245) ([ajhager](https://github.com/ajhager))
- Feature/fully lowercase field names [\#242](https://github.com/asteris-llc/converge/pull/242) ([rebeccaskinner](https://github.com/rebeccaskinner))
- file.directory resource [\#241](https://github.com/asteris-llc/converge/pull/241) ([BrianHicks](https://github.com/BrianHicks))
- Add user group support [\#234](https://github.com/asteris-llc/converge/pull/234) ([arichardet](https://github.com/arichardet))
- support lowercase field names [\#231](https://github.com/asteris-llc/converge/pull/231) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Fix/query refactor [\#230](https://github.com/asteris-llc/converge/pull/230) ([rebeccaskinner](https://github.com/rebeccaskinner))
- handle params that are thunked [\#229](https://github.com/asteris-llc/converge/pull/229) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Feature/query module [\#228](https://github.com/asteris-llc/converge/pull/228) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Refactor helpers into descriptive names [\#222](https://github.com/asteris-llc/converge/pull/222) ([BrianHicks](https://github.com/BrianHicks))
- Feature/value passing [\#220](https://github.com/asteris-llc/converge/pull/220) ([rebeccaskinner](https://github.com/rebeccaskinner))
- rpc: flatten details field [\#218](https://github.com/asteris-llc/converge/pull/218) ([BrianHicks](https://github.com/BrianHicks))
- Feature/formatter \#208 [\#217](https://github.com/asteris-llc/converge/pull/217) ([BrianHicks](https://github.com/BrianHicks))
- docker.container: fix entrypoint diff logic [\#215](https://github.com/asteris-llc/converge/pull/215) ([ryane](https://github.com/ryane))
- Graphing over RPC [\#214](https://github.com/asteris-llc/converge/pull/214) ([BrianHicks](https://github.com/BrianHicks))
- docker.container resource [\#203](https://github.com/asteris-llc/converge/pull/203) ([ryane](https://github.com/ryane))
- Rework logs to use logrus [\#202](https://github.com/asteris-llc/converge/pull/202) ([BrianHicks](https://github.com/BrianHicks))
- Feature/platform refactor [\#197](https://github.com/asteris-llc/converge/pull/197) ([stevendborrelli](https://github.com/stevendborrelli))
- Added host environment variable support [\#195](https://github.com/asteris-llc/converge/pull/195) ([arichardet](https://github.com/arichardet))
- resources\(docker image\): document [\#193](https://github.com/asteris-llc/converge/pull/193) ([BrianHicks](https://github.com/BrianHicks))
- Documentation Site [\#192](https://github.com/asteris-llc/converge/pull/192) ([BrianHicks](https://github.com/BrianHicks))
- Feature/health module [\#189](https://github.com/asteris-llc/converge/pull/189) ([rebeccaskinner](https://github.com/rebeccaskinner))
- docker.image resource redux [\#188](https://github.com/asteris-llc/converge/pull/188) ([ryane](https://github.com/ryane))
- Basic RPC [\#187](https://github.com/asteris-llc/converge/pull/187) ([BrianHicks](https://github.com/BrianHicks))
- shell env and working dir support [\#185](https://github.com/asteris-llc/converge/pull/185) ([ryane](https://github.com/ryane))
- \[GRAPH\] Fix \#158, add triggering files, add tests [\#181](https://github.com/asteris-llc/converge/pull/181) ([sehqlr](https://github.com/sehqlr))
- don't panic when result status is nil [\#180](https://github.com/asteris-llc/converge/pull/180) ([ryane](https://github.com/ryane))
- Feature/optimize human printer [\#175](https://github.com/asteris-llc/converge/pull/175) ([arichardet](https://github.com/arichardet))
- Shrink binaries with some magical ldflags [\#174](https://github.com/asteris-llc/converge/pull/174) ([BrianHicks](https://github.com/BrianHicks))
- Feature/shell refactor [\#172](https://github.com/asteris-llc/converge/pull/172) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Dependency ordered graph walking [\#166](https://github.com/asteris-llc/converge/pull/166) ([BrianHicks](https://github.com/BrianHicks))
- Migrate context for go1.7, refactor fetch/http.go [\#164](https://github.com/asteris-llc/converge/pull/164) ([sehqlr](https://github.com/sehqlr))
- Feature/resource return refactor [\#157](https://github.com/asteris-llc/converge/pull/157) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Move errors into the tree [\#154](https://github.com/asteris-llc/converge/pull/154) ([BrianHicks](https://github.com/BrianHicks))
- Human-readable pretty printer [\#152](https://github.com/asteris-llc/converge/pull/152) ([BrianHicks](https://github.com/BrianHicks))
- make nil a useful default bottom value for subgraphdi [\#150](https://github.com/asteris-llc/converge/pull/150) ([rebeccaskinner](https://github.com/rebeccaskinner))
- JSONL output and interface refactoring [\#148](https://github.com/asteris-llc/converge/pull/148) ([BrianHicks](https://github.com/BrianHicks))
- Feature/template split [\#136](https://github.com/asteris-llc/converge/pull/136) ([rebeccaskinner](https://github.com/rebeccaskinner))
- fetch: add context [\#135](https://github.com/asteris-llc/converge/pull/135) ([BrianHicks](https://github.com/BrianHicks))
- Render merged subtrees [\#133](https://github.com/asteris-llc/converge/pull/133) ([BrianHicks](https://github.com/BrianHicks))
- Fix/119 [\#132](https://github.com/asteris-llc/converge/pull/132) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Trim Duplicate Subtrees [\#131](https://github.com/asteris-llc/converge/pull/131) ([BrianHicks](https://github.com/BrianHicks))
- Appropriately handle required parameters instead of crashing [\#130](https://github.com/asteris-llc/converge/pull/130) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Use concurrent map implementation for graph values [\#128](https://github.com/asteris-llc/converge/pull/128) ([BrianHicks](https://github.com/BrianHicks))
- Make params missing defaults required [\#127](https://github.com/asteris-llc/converge/pull/127) ([BrianHicks](https://github.com/BrianHicks))
- makefile: simplify linting [\#126](https://github.com/asteris-llc/converge/pull/126) ([BrianHicks](https://github.com/BrianHicks))
- Rename "template" to "file.content" [\#124](https://github.com/asteris-llc/converge/pull/124) ([BrianHicks](https://github.com/BrianHicks))
- gometalinter fixes [\#122](https://github.com/asteris-llc/converge/pull/122) ([BrianHicks](https://github.com/BrianHicks))
- Graphviz [\#118](https://github.com/asteris-llc/converge/pull/118) ([rebeccaskinner](https://github.com/rebeccaskinner))
- Rewrite graph system [\#116](https://github.com/asteris-llc/converge/pull/116) ([BrianHicks](https://github.com/BrianHicks))
- resource\(shell\): add interpreter [\#109](https://github.com/asteris-llc/converge/pull/109) ([BrianHicks](https://github.com/BrianHicks))
- switch to glide [\#108](https://github.com/asteris-llc/converge/pull/108) ([BrianHicks](https://github.com/BrianHicks))
- Feature/context [\#107](https://github.com/asteris-llc/converge/pull/107) ([BrianHicks](https://github.com/BrianHicks))
- Format samples in readme [\#105](https://github.com/asteris-llc/converge/pull/105) ([ajhager](https://github.com/ajhager))
- Port shell resource and tests [\#101](https://github.com/asteris-llc/converge/pull/101) ([ajhager](https://github.com/ajhager))
- resource\(template\): implement and test template [\#98](https://github.com/asteris-llc/converge/pull/98) ([BrianHicks](https://github.com/BrianHicks))
- Application for preparers [\#97](https://github.com/asteris-llc/converge/pull/97) ([BrianHicks](https://github.com/BrianHicks))
- Plan in new preparers [\#95](https://github.com/asteris-llc/converge/pull/95) ([BrianHicks](https://github.com/BrianHicks))
- dependencies should block execution [\#81](https://github.com/asteris-llc/converge/pull/81) ([siddharthist](https://github.com/siddharthist))
- dependency tracker [\#80](https://github.com/asteris-llc/converge/pull/80) ([BrianHicks](https://github.com/BrianHicks))
- Self Serve [\#79](https://github.com/asteris-llc/converge/pull/79) ([BrianHicks](https://github.com/BrianHicks))
- Module server [\#78](https://github.com/asteris-llc/converge/pull/78) ([BrianHicks](https://github.com/BrianHicks))
- File mode module [\#76](https://github.com/asteris-llc/converge/pull/76) ([BrianHicks](https://github.com/BrianHicks))
- Simplify Application [\#73](https://github.com/asteris-llc/converge/pull/73) ([BrianHicks](https://github.com/BrianHicks))
- Fix graph with dashes [\#71](https://github.com/asteris-llc/converge/pull/71) ([BrianHicks](https://github.com/BrianHicks))
- Feature/logging [\#70](https://github.com/asteris-llc/converge/pull/70) ([BrianHicks](https://github.com/BrianHicks))
- add color printing again [\#69](https://github.com/asteris-llc/converge/pull/69) ([siddharthist](https://github.com/siddharthist))
- Feature/basic blackbox testing [\#68](https://github.com/asteris-llc/converge/pull/68) ([BrianHicks](https://github.com/BrianHicks))
- Update 0.1 docs [\#60](https://github.com/asteris-llc/converge/pull/60) ([BrianHicks](https://github.com/BrianHicks))
- fixed souceFile.hcl [\#59](https://github.com/asteris-llc/converge/pull/59) ([Dacode45](https://github.com/Dacode45))
- Bug/names [\#58](https://github.com/asteris-llc/converge/pull/58) ([Dacode45](https://github.com/Dacode45))
- fmt [\#57](https://github.com/asteris-llc/converge/pull/57) ([BrianHicks](https://github.com/BrianHicks))
- Feature/requirements [\#54](https://github.com/asteris-llc/converge/pull/54) ([Dacode45](https://github.com/Dacode45))
- don't ignore the error when Viper binds to PFlags [\#50](https://github.com/asteris-llc/converge/pull/50) ([siddharthist](https://github.com/siddharthist))
- color summaries [\#49](https://github.com/asteris-llc/converge/pull/49) ([siddharthist](https://github.com/siddharthist))
- add key=value parameter parsing [\#47](https://github.com/asteris-llc/converge/pull/47) ([siddharthist](https://github.com/siddharthist))
- Summary lines [\#46](https://github.com/asteris-llc/converge/pull/46) ([BrianHicks](https://github.com/BrianHicks))
- render: fix swallowed errors [\#44](https://github.com/asteris-llc/converge/pull/44) ([BrianHicks](https://github.com/BrianHicks))
- limit identifiers to word chars, dashes, and dots [\#43](https://github.com/asteris-llc/converge/pull/43) ([BrianHicks](https://github.com/BrianHicks))
- resource\(template\): fix permissions [\#42](https://github.com/asteris-llc/converge/pull/42) ([BrianHicks](https://github.com/BrianHicks))
- Lazily load parameters - defaults can be templated [\#37](https://github.com/asteris-llc/converge/pull/37) ([siddharthist](https://github.com/siddharthist))
- commands: add streaming output [\#36](https://github.com/asteris-llc/converge/pull/36) ([BrianHicks](https://github.com/BrianHicks))
- Apply [\#32](https://github.com/asteris-llc/converge/pull/32) ([BrianHicks](https://github.com/BrianHicks))
- load: accept initial parameters [\#28](https://github.com/asteris-llc/converge/pull/28) ([BrianHicks](https://github.com/BrianHicks))
- Add tests for Plan [\#25](https://github.com/asteris-llc/converge/pull/25) ([BrianHicks](https://github.com/BrianHicks))
- load: from http [\#23](https://github.com/asteris-llc/converge/pull/23) ([siddharthist](https://github.com/siddharthist))
- plan: pretty-printing terminal output [\#22](https://github.com/asteris-llc/converge/pull/22) ([siddharthist](https://github.com/siddharthist))
- fix incomplete comment [\#19](https://github.com/asteris-llc/converge/pull/19) ([BrianHicks](https://github.com/BrianHicks))
- Checker [\#9](https://github.com/asteris-llc/converge/pull/9) ([BrianHicks](https://github.com/BrianHicks))
- Name\(\) -\> String\(\) [\#8](https://github.com/asteris-llc/converge/pull/8) ([siddharthist](https://github.com/siddharthist))
- \[WIP\] Graphviz Support [\#7](https://github.com/asteris-llc/converge/pull/7) ([BrianHicks](https://github.com/BrianHicks))
- Feature/validation error type [\#3](https://github.com/asteris-llc/converge/pull/3) ([siddharthist](https://github.com/siddharthist))
- shelltask: add 'validate' method [\#2](https://github.com/asteris-llc/converge/pull/2) ([siddharthist](https://github.com/siddharthist))
- readme: status -\> check [\#1](https://github.com/asteris-llc/converge/pull/1) ([siddharthist](https://github.com/siddharthist))

## 0.2.0 (September 26, 2016)

### Enhancements

- shell env and working dir support (#185)
- Docker image resource (#188)
- RPC support (#187)
- Expose basic platform information (#194)
- Docker container resource (#203)
- Support for create/delete linux groups (#234)
- Refactor Status Interface (#237)
- Refactor `check` and `apply` behavior (#240)
- Feature/module verification (#245)
- Reduce logging verbosity (#248)
- Boolean support in parameters (#251)
- Change Port to 4774 after IANA approval (#261)
- Support for create/delete linux users (#259)

### Bug fixes

- don't panic when result status is nil (#180)
- Fix error handling when fail to print results during plan or apply (#183)
- Order fixes (#254)
- Race condition fixes (#266)
- Fix/user group - Allow adding/deleting without gid (#283)
- Perform healthchecks over RPC (#289)
- Fix/294 field name conflicts (#301)
- Use thread-safe field cache (#303)

### Documentation/Examples

- Documentation Site (#192)
- Docker Swarm mode (#267)
- ELK (Elasticsearch, Logstash, and Kibana) stack (#272)
- Add CodeClimate Checks to build (#281)
- docs: add draft resource authors guide (#290)


\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*

\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*

\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*

\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*

\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*