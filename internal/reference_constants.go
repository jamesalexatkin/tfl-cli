package internal

// ASCII references

//nolint:gochecknoglobals,unused
var minisculeRoundel = `
 ╭───╮
───────
 ╰───╯
`

//nolint:gochecknoglobals,unused
var exampleDepartureBoard = `
   ╭────────────────────────────────╮
┌──┤ %s                             ├──┐
|  └────────────────────────────────┘  |
│  ╭───╮                               |
| ───────  Platform 3 (Elizabeth Line) |
|  ╰───╯                               │
├──────────────────────────────────────|
| 1 Heathrow Terminal 4 - 5mins        |
| 2 Heathrow Terminal 4 - 14mins       |
| 3 Heathrow Terminal 4 - 29mins       |
| 4 Heathrow Terminal 4 - 44mins       |
|                                      |
├──────────────────────────────────────|
│  ╭───╮                               |
| ───────  Platform 4 (Elizabeth Line) |
|  ╰───╯                               │
├──────────────────────────────────────|
| 1 Heathrow Terminal 4 - 5mins        |
| 2 Heathrow Terminal 4 - 14mins       |
| 3 Heathrow Terminal 4 - 29mins       |
| 4 Heathrow Terminal 4 - 44mins       |
└──────────────────────────────────────┘
`

//nolint:gochecknoglobals,unused
var exampleStatusBoard = `
╭───────────────────────────╮             ╭───────────────────────────╮             ╭───────────────────────────╮             ╭───────────────────────────╮
│ London Underground        │             │ London Overground         │             │ ELIZABETH LINE            │             │ DLR                       │
├───────────────────────────┴─────────────┼───────────────────────────┴─────────────┼───────────────────────────┴─────────────┼───────────────────────────┴─────────────┐
│  ╭───╮                                  │  ╭───╮                                  │  ╭───╮                                  │  ╭───╮                                  │
│ ───────                                 │ ───────                                 │ ───────                                 │ ───────                                 │
│  ╰───╯                                  │  ╰───╯                                  │  ╰───╯                                  │  ╰───╯                                  │
│ █ Bakerloo: Minor Delays                │ █ Liberty: Good Service                 │ █ Elizabeth Line: Good Service          │ █ DLR: Good Service                     │
│ █ Central: Good Service                 │ █ Lioness: Good Service                 │                                         │                                         │
│ █ Circle: Good Service                  │ █ Mildmay: Good Service                 │                                         │                                         │
│ █ District: Good Service                │ █ Suffragette: Good Service             │                                         │                                         │
│ █ Hammersmith & City: Good Service      │ █ Weaver: Minor Delays                  │                                         │                                         │
│ █ Jubilee: Good Service                 │ █ Windrush: Good Service                │                                         │                                         │
│ █ Metropolitan: Good Service            │                                         │                                         │                                         │
│ █ Northern: Good Service                │                                         │                                         │                                         │
│ █ Piccadilly: Good Service              │                                         │                                         │                                         │
│ █ Victoria: Minor Delays                │                                         │                                         │                                         │
│ █ Waterloo & City: Minor Delays         │                                         │                                         │                                         │
└─────────────────────────────────────────┴─────────────────────────────────────────┴─────────────────────────────────────────┴─────────────────────────────────────────┘
`
