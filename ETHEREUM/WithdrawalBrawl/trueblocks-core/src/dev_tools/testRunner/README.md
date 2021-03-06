## testRunner

The `testRunner` tool runs the TrueBlocks tests. The test cases are stored in CSV files a folder called `./testCases`.

Test cases contain the following fields:

| field   | meaning                                                                                  |
| ------- | ---------------------------------------------------------------------------------------- |
| onOff   | either `on` or `off` to enable/disable the test                                          |
| mode    | one of `cmd`, `api` or `both` (runs against command line, API server or both)            |
| speed   | one of `slow`, `medium` or `fast` allows test runner to skip tests that take a long time |
| route   | the API route of the tool                                                                |
| tool    | the command line location of the tool                                                    |
| name    | the name of the resulting test case file                                                 |
| post    | one of `y` or `n` determines if the resulting output is run through `jq` JSON formatter  |
| options | the command line / API url options (in API format)                                       |
| extra   | extra data needed for some tests                                                         |

You may run all tests from the `./build` folder by entering `make test-all` (or `make tests` to run only the `cmd/fast` cases).

### Usage

`Usage:`    testRunner [-m|-f|-c|-r|-v|-h]  
`Purpose:`  Run TrueBlocks' test cases with options.

`Where:`

|          | Option                           | Description                                                                  |
| -------- | -------------------------------- | ---------------------------------------------------------------------------- |
| &#8208;m | &#8208;&#8208;mode &lt;val&gt;   | determine which set of tests to run, one of [cmd, api,<br/>both]             |
| &#8208;f | &#8208;&#8208;filter &lt;val&gt; | determine how long it takes to run tests, one of [fast,<br/>medi, slow, all] |
| &#8208;c | &#8208;&#8208;clean              | clean working folder before running tests                                    |
| &#8208;r | &#8208;&#8208;report             | display performance report to screen                                         |
| &#8208;x | &#8208;&#8208;fmt &lt;val&gt;    | export format, one of [none, json, txt, csv, api]                            |
| &#8208;v | &#8208;&#8208;verbose            | set verbose level (optional level defaults to 1)                             |
| &#8208;h | &#8208;&#8208;help               | display this help screen                                                     |

#### Other Options

All tools accept the following additional flags, although in some cases, they have no meaning.

| Command     | Description                                                   |
| ----------- | ------------------------------------------------------------- |
| --version   | display the current version of the tool                       |
| --wei       | export values in wei (the default)                            |
| --ether     | export values in ether                                        |
| --dollars   | export values in US dollars                                   |
| --raw       | pass raw RPC data directly from the node with no processing   |
| --to_file   | write the results to a temporary file and return the filename |
| --output fn | write the results to file 'fn' and return the filename        |
| --file fn   | specify multiple sets of command line options in a file       |

<small>*For the `--file fn` option, you may place a series of valid command lines in a file using any of the above flags. In some cases, this may significantly improve performance. A semi-colon at the start of any line makes that line a comment.*</small>
