# Goleet

Leetcode solutions with Go!

How does everything work?

Devcontainers is used for development setup. [A script](./.devcontainer/post-install.sh) installs necessary dependencies after container is built.

Tests are generated with [gotests](https://github.com/cweill/gotests) & ran in watch mode with [Go Watch (gow)](https://github.com/mitranim/gow)

```bash
gow test two-sum.go two-sum_test.go
```

## Want to try it out?

Fork the repository and use on [Github Codespaces](https://github.com/features/codespaces) or your local machine with Docker & Visual Studio Code

| ID  | Name                                                                                              | Difficulty |           Tags            |
| --- | ------------------------------------------------------------------------------------------------- | :--------: | :-----------------------: |
| 1   | [Two Sum](https://leetcode.com/problems/two-sum/)                                                 |    easy    |           array           |
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) |    easy    |           array           |
| 509 | [Fibonacchi Number](https://leetcode.com/problems/fibonacci-number/)                              |    easy    | recursion, implementation |
| 55  | [Jump Game](https://leetcode.com/problems/jump-game/)                                             |   medium   |        dp, greedy         |
