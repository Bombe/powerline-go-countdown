# powerline-go-countdown

**powerline-go-countdown** is a segment for [powerline-go](https://github.com/justjanne/powerline-go) that shows the number of remaining days until a date. It can be used to count down the days to your next vacation, your wedding, or your retirement, or whatever else you are looking forward to.

The configuration is done via JSON, in `$XDG_CONFIG_HOME/powerline-go/countdown.json`. Here’s an example:

```json
{
  "deadlines": [
    {
      "date": "2026-06-21",
      "occasion": "Summer Solstice",
      "symbol": "🌞"
    },
    {
      "date": "2026-03-07",
      "occasion": "Local Election",
      "symbol": "🗳️",
      "color": "#ff0",
      "bgcolor": "#f00"
    }
  ]
}
```

The only real important piece of configuration is **date**, everything else is optional. **occasion** is a value that only exists to make it easier for you to navigate your configuration file. The **symbol** is prepended to the number of remaining days. **color** and **bgcolor** control the foreground and background color of the segment; when omitted, **powerline-go** will use its defaults.
