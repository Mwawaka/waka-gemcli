For GemCLI, I'll implement Continuous Delivery -- every new version tag automatically builds binaries and pushes the Docker image, but I decide when to tag a release

Developer pushes code
        ↓
CI — run fmt, vet, test, build    ← catches bugs early
        ↓
CD — build binaries, push Docker image, create GitHub Release    ← ships automatically
