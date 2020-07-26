package cameo

const Favicon = `
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
`
