# message code

### message code components
message code is a 6-digit number, use `ABCDEF` to present each digit
- `A`: the log level, 1-debug, 2-info, 3-warn, 4-error
- `BC`: the module number
- `D`: the submodule
- `EF`: the sequence number

### relations between code and module

| BC  | module      | D   | submodule          |
|-----|-------------|-----|--------------------|
| 00  | message     | 0   | general            |
| 01  | metadata    | 0   | app                |
| 01  | metadata    | 1   | db                 |
| 01  | metadata    | 2   | env                |
| 01  | metadata    | 3   | middleware cluster |
| 01  | metadata    | 4   | middleware server  |
| 01  | metadata    | 5   | monitor system     |
| 01  | metadata    | 6   | mysql cluster      |
| 01  | metadata    | 7   | mysql server       |
| 01  | metadata    | 8   | table              |
| 01  | metadata    | 9   | user               |
| 02  | metadata    | 0   | resource group     |
| 03  | healthcheck | 0   | default engine     |
| 03  | healthcheck | 1   | service            |
| 04  | query       | 0   | query              |
| 05  | sqladvisor  | 0   | service            |
| 06  | alert       | 0   | http               |
| 06  | alert       | 1   | service            |
| 07  | privilege   | 0   | service            |
| 08  | router      | 0   | token              |
| 09  | health      | 0   | health             |
