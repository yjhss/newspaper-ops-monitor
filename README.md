# newspaper-ops-monitor

### Branches
`<type>/<issue-number>-<description>`

| Prefix | 용도 | 예시 |
|---|---|---|
| `feat/` | 새로운 기능 | `feat/12-file-stability` |
| `fix/` | 버그 수정 | `fix/18-tabloid-pair-parser` |
| `refactor/` | 기능 변화 없는 구조 개선 | `refactor/25-monitor-service` |
| `docs/` | 문서 | `docs/30-architecture` |
| `chore/` | 설정, CI, 의존성 등 | `chore/5-github-templates` |
| `test/` | 테스트 추가 | `test/22-reissue-scenarios` |

### Commit
`<type>: <summary> (#<issue-number>)`

| Type | 용도 | 예시 |
|---|---|---|
| `feat` | 새로운 기능 추가 | `feat: add file stability detection (#2)` |
| `fix` | 버그 수정 | `fix: handle invalid media config (#3)` |
| `refactor` | 기능 변화 없는 코드 구조 개선 | `refactor: simplify monitor service (#4)` |
| `test` | 테스트 추가/수정 | `test: add snapshot detection tests (#5)` |
| `docs` | 문서 작성/수정 | `docs: add architecture guide (#6)` |
| `chore` | 설정, 환경, 의존성, 템플릿 | `chore: add GitHub templates (#7)` |
| `ci` | CI/CD 설정 | `ci: add GitHub Actions workflow (#8)` |
| `perf` | 성능 개선 | `perf: reduce repeated directory scans (#9)` |

### Issue
`[Type] Short description`

### PR
`[Type] Short description (#issue-number)`
