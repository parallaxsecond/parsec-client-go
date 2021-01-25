# Parsec Go Client End to End  (Continuous Integration) tests

Currently only the CI test with all providers enabled is working at all.  To run it, you will need docker.

```bash
./ci-all.sh
```

The ci-*.sh scripts all build docker images defined in the provider_cfg folder, which has subfolders for each of the provider configurations.

The docker containers run the ci.sh script in the top level folder.  This script is in this folder so it can simply access the whole of the repository, but may
be refactored into this folder at some point in the future.