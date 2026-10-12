# Reproducing the assignment stages

Use `git log --oneline` to select each commit. In a separate checkout of that
commit, enter `grade-calculator` and run:

    go test -v -count=1 -coverprofile=coverage.out ./...
    go tool cover -func=coverage.out

Part 5 copied the Part 4 module into a scratch folder, commented out every
`t.Errorf` call while preserving setup, comparisons and production calls, and ran
the same command. Coverage profiles were identical with and without these failure
checks. The final code retains assertions. Profiles refer to source locations at
their own stage; use the matching checkout to interpret an earlier profile.
