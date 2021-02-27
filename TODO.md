### Features

- Validation config: `allowMissingStage`, etc.

- Try to infer the `stage` of a KEP

  ```golang
  func (p *Proposal) InferStage() error {
    stage := p.Stage
    milestone := p.LatestMilestone

    releaseSemVer, err := semver.ParseTolerant(ActiveRelease)
    if err != nil {
      return errors.Wrap(
        err,
        "creating a SemVer object for the active release milestone",
      )
    }

    if stage != "" {
      return nil
    } else {
      if p.Milestone.Alpha != "" {
        alphaSemVer, err := semver.ParseTolerant(p.Milestone.Alpha)
        if err != nil {
          return errors.Wrap(
            err,
            "creating a SemVer object for the alpha milestone",
          )
        }

        if releaseSemVer.EQ(alphaSemVer) {
          
        }
      }
    }

    return nil
  }
  ```

- Set `stage` to required
- `RepoClient` to hold all the things

### Done

#### First pass

- [x] enable strict checks
- [x] assert to require
- [x] yaml to v3
- [x] go mod to go1.15 and bump dependencies
- [x] import k/release/pkg/log + logrus
- [x] need OWNERS
- [x] add badges
- [x] why is cmd/kepval/main_test.go in cmd/ without a command
- [x] verify-go-mod is too strict
- [x] proper cobra structure
  - [x] kepctl
  - [x] ~~kepval (does this actually exist anymore?)~~
- [x] pull PRR structs into API
- [x] cleanup makefile and hack directory

### Follow-ups

- [ ] stop calling directories `hack/` (use `scripts/`)
- [ ] split tests into:
  - go
  - content verification
  - build verification
- [ ] migrate kepify to cobra
- [ ] plumb logrus (replace `fmt`)
- [ ] refactor out `Person`/`Reviewer` field
- [ ] no references to API in commands (means we haven't encapsulated properly)
- [ ] use k/release github token package, if exists
- [ ] inconsistent flag choices

#### CI

- [ ] use Makefile instead of hack/verify.sh in CI
- [ ] need py3 for boilerplate script
- [ ] need docker for shellcheck

#### k/release

- [ ] update k/repo-infra version
- [ ] go1.16 go.mod

### Handoff (@jeremyrickard)

- [ ] sweep TODOs
- [ ] generate godoc?
- [ ] no type assertions
- [ ] no "helpers"
- [ ] no util package
