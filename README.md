# EORM

[![codecov](https://codecov.io/gh/gotomicro/eorm/branch/main/graph/badge.svg?token=vc0BDor3Lk)](https://codecov.io/gh/gotomicro/eorm)

简单的 ORM 框架。

## 使用注意事项（技术选型前必读）

丑话说前头。

- [eorm 支持的字段类型](https://github.com/gotomicro/eorm/discussions/71)：这是指在 Go 语言层面上支持的类型，eorm 本身并不关心数据库表里面定义的类型；

### Go 版本

请使用 Go 1.18 以上版本。

### SQL 2003 标准
理论上来说，我们计划支持 [SQL 2003 standard](https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#query%20specification). 不过据我们所知，并不是所有的数据库都支持全部的 SQL 2003 标准，所以用户还是需要进一步检查目标数据库的语法。

### 全中文仓库

这是一个全中文的仓库。这意味着注释、文档和错误信息，都会是中文。介意的用户可以选择别的 ORM 仓库。

但是不必来反馈说希望提供英文版本，我们是不会提供的，因为：
- 这个仓库目前以及可预测的将来，都不会有外国人使用。很多国内开发者的开原仓库提供了英文选项但是实际上英文用户数量感人，我就不想做这种性价比低的事情；
- 双语言会导致英文用户有一些不切实际的期望，比如说在 issues 和 discussions 里面都使用英文交流。但是这对于我的目标来说，也是不现实的。因为很多国内开发者不具备这种英文能力，而且这些开发者非常倔强，即便我们一再强调请使用英文，他们也会固执使用中文。因此我索性将这个搞成全中文仓库，一了百了；
- 中文终究是我的母语，所以使用中文的表达能力更加准确；
- 翻译软件非常发达，真有英语用户要用，可以自己找翻译软件；

### 社区组织和讨论

短时间内，我不会组建任何的微信群，QQ 群或者钉钉群之类的即时通讯群。

我注意到很多开源仓库都会组织类似的群，但是这种群有利有弊，而且对于项目本身来说是弊大于利的。组建了不同的群之后会导致问题讨论被切割到不同的群里面。例如某个群讨论了 A 问题，其它群完全看不到。

另外一个原因就是，因为即时通讯过于便捷，会给维护者带来庞大的维护压力。用户可能期望自己的所有答案都能从群里得到解答，因此不愿意自己花时间去读文档，读注释，读例子。在这种情况下，他们会频繁艾特维护者，并希望维护者能够实时给出详细回答。而实际情况是，一般的小项目维护者可能只有两三个人，所以没有足够的精力来维护这种群。

我想要的社区是大家统一在 github 下，利用 issue 和 discussion 来讨论问题。这样别的用户都可以搜索到所有的讨论。

## 加入我们
- [贡献指南](.docs/contribution.md)

我们欢迎任何人给我们提合并请求，但是我们希望合并请求能够做到：
- 一个合并请求一个 Commit ID
- 自己要先确保合并请求能够通过 CI
- 我们使用 uber 的[代码风格](https://github.com/uber-go/guide/blob/master/style.md)

我们尤其欢迎两类人：
- 新特性试用者
- 长期稳定贡献者

### 设置开发环境

如果你是 Windows 用户，那么我们建议你使用 WSL，因为这个仓库会使用到一个 Unix 命令来帮助构建、测试等。

#### 安装 golangci-lint
参考 [Install golangci-lint](https://golangci-lint.run/usage/install/)
#### 设置 pre-push github hook
将`.github/pre-push` 复制到本仓库的 `.git` 目录下
