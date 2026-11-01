# KPanel v1.29.0-rc.3 发布验收记录

日期：2026-11-01；发布级别 L3；releaseChannel=preview；releaseTrain=1.29.0。

候选提交 / 标签：`fd7033717f6f47d66277da760e7d255bd85d3c25` / `v1.29.0-rc.3`。公开 [Release](https://github.com/kejilion/KPanel/releases/tag/v1.29.0-rc.3)，ID 430453650，draft=false、prerelease=true。产物已发布，生产未部署。

上一稳定版本/回滚点为 v1.28.0、`sha256:bd90d2bf69f059b4da29e456054bfb1785b47a6b1d453e5941e299a4a7de78c4`；前一预览 v1.29.0-rc.1、`sha256:655e84cc25d5a6343ca932f548e462fa6984d3bd1b0a2a2e6e08b06ccea45f12` 保留。批准基线 `aa9776784569dc3a9bc88b0b6d64cc98d545bc91`；当前唯一候选 `release/v1.29.0-candidate` 保留到本列车稳定版，不覆盖 rc.1 或既有镜像。

## 发布画像与范围

- 业务域：内置脚本的换源、Nginx/LDNMP 安装兼容、设置页 HTTPS DNS 指引，以及 AI 多会话创建任务准入。展示/只读与既有写入入口的兼容修复；无新 API、Agent 权限、端口、数据 Schema、Compose 格式或 WAF 版本。
- LinuxMirrors 修复 `38b9fcd` 和 Armbian 前置 `ed58dea` 在 rc.1 脚本基线上重放为 `36ad597`、`9872766`；`54a337d` 完整重放为 `496d363`，包括 command install、重复安装保留配置和失败提前返回。
- HTTPS 指引来源 `050c7a0a` 重放为 `b5246143`，复用既有组件/语言资源和公网地址接口；查询失败不会阻断域名输入。浏览器发现 @ 内联提示继承缩放后为11.44px，独立修复 `37e0a033` 经 `f441c64e` 将其限定为12px。
- 版本/源码资格在 `8cc53bcf` 准备，command install 资格在 `c3f6fc95` 补入，`454e9ac7` 记录扩大范围；rc.2 的 `aa977678` 仅补已公开脚本文档。rc.2 源码门禁发现 SQLite517 后，新增 AI 准入方法体修复和16并发回归，source `38716407` 经独立复核，发布派生元数据 `b91d429e` 再独立资格化，最终冻结为 `fd7033717f6f47d66277da760e7d255bd85d3c25`。原 HTTPS 组件/测试与454e9ac7逐字节相同；当前 web tree `210bb1a1fd498eae2c49c91c60d553f9b3ec5ced`。
- FastAdmin 函数、固定包/模板与七文件 WAF 清单逐字节保留。未纳入脚本主线其余 WAF/调优更新，也未将历史本地分支按名称推断为待发布；原盘点166项已在批准main、3项精确tip已归档、64项历史未决保留。
- 风险 L3 来自实际版本/镜像发布；本轮不执行生产部署、真实主机换源、Docker 安装、TLS 初始化、认证登录或宿主自动更新回滚。

## 跨仓库联动及来源

scriptLinkageState=coupled，变更集 linuxmirrors-preview-20261101。旧镜像脚本 `5b8327f` / `2ba489e89be5320537135a1ca0bdd3c1b69f5ff927c548593ec6ee9ae1ed51e4`；新固定脚本 `496d3633d850ddc68bcef654b8d9d11054a28a68`，根 SHA-256 `9cd4b90c9b1d509fba453ab7a68ef6b33cb82c24ae599031bc65430205c5b96d`、CN `9b95e51696ccd08c67aeb994b4c8746ed87499d0af1144b987fb05717cec9a09`。

脚本已先在独立clone本地合入main，再仅推送 main:main；公开main `5e085c8c23efa01d47d58dce77535ca78017e6b6` 的文件树与777908e相同，仅保存兼容提交可达性，没有推送功能分支/PR。匿名公开根/CN与模板字节匹配；实际GitHub镜像ADD及公开双架构标签/字节通过精确候选核对。FastAdmin模板保持 `nginx@c830bf34ae6da8e183b4892860fbce30ea6dc651`，SHA-256 `986b57db53211332331c1c0022e7894546c8e8d3a1e4baeef5f2c95f01c23f46`。

根/CN归一化相等，两次原始脚本重放及54a337d的stable patch ID相同。WAF revision/manifest及FastAdmin函数与5b8327f完全一致；不以脚本main头的额外功能替代本版固定来源。回退仍需配套数据备份，不自动删除既有数据库/站点。

## 多维质量

| 维度 | 状态 | 证据与边界 |
| --- | --- | --- |
| 业务正确性/双端兼容 | 已验证 | 八语言换源兼容、Armbian夹具、首次/重复安装、完整性失败和提前停止；新入口仍沿用脚本契约。全宿主换源/安装未覆盖 |
| 网络入侵/供应链 | 已验证 | 固定提交/摘要、公开字节、CI/L3扫描及provenance；保留现有37项漏洞处置，不声称零漏洞或全面安全 |
| 稳定/失败恢复 | 已验证 | 参数/stdin/退出码、下载失败、配置保留、摘要失败无半份配置、容器启动前返回；真实宿主升级/回退未覆盖 |
| 性能/资源 | 已验证 | AI只在创建任务的短事务开始预留写入锁，不增加队列/重试/连接池/超时；Linux原取消与新并发测试200轮、AI整包和Windows定向20轮通过。未做吞吐基准，不承诺任意外部长期持锁无错误；Nginx容器128MiB/1CPU/pids64/network-none |
| 体验/可访问性 | 已验证 | 11项定向测试、类型检查/构建，mock空IP/键盘与窄宽视口；14/13/12px，guide宽/scroll320，page380，焦点3px。原生200%、其他主题/语言和截图未验证 |
| 数据/配置迁移 | 不适用 | 无数据格式迁移；既有Nginx/Compose/WAF保护通过夹具和真实Nginx语法检查，非全宿主安装成功声明 |

## 独立复核与安全审查

干净Codex复核30371c9e..8cc53bcf与脚本5b8327f..36ad597为PASS WITH FOLLOW-UP，仅针对本地准备，其他提供商CLI不可用。扩大到54a337d后由发布者核对原始补丁、逐字节保留范围及定向回归；先前独立复核只复用不变部分，不冒充扩大范围全新独立审查。原脚本/HTTPS主要业务来源于独立作者；发布者修复AI真实锁冲突，并由干净独立Codex上下文另行复核生产方法/测试和最终派生元数据，结论PASS，其他提供商CLI仍不可用。独立overlay补验持锁取消无插入/锁释放、8并发retry恰一成功及父任务权限等7项通过。

OCR1.12.11先完成自由臂再圈选，最初7/7覆盖、有效H0/M0/L0、constrained-only=0；扩大范围已见过重叠规则，明确blind=false、constrained-only=unreported，再核对7/7，不计新盲测收益。AI源码自由臂0、约束臂2/2、constrained-only=0；最终发布扩大范围已见重叠规则，6/6、constrained-only=unreported。安全敏感自动范围internal/ai由无实现上下文的候选盲审实际执行：原始H0/M2/L0，核实1、修复0、待办1；一项既有累计历史容量风险由KPanel release maintainers负责，2026-11-08前复核，吞吐扩展或ai.db存储压力时触发；另一连接池外键线索被精确驱动支持和双连接删除级联探针反证。原始线索/核实仅留仓库外。审查模型使用父任务继承默认配置，运行标识未暴露；早期元数据中误标的具体型号以最终inherited-default声明更正。本版未执行CF审计或稳定前补漏，不把未执行记录为通过，也不改变稳定周期退出判定。

## 自动门禁与供应链资格

- 候选push CI [41284813874](https://github.com/kejilion/KPanel/actions/runs/41284813874) 同SHA成功，Linux/Windows均核对；依赖report [41284813899](https://github.com/kejilion/KPanel/actions/runs/41284813899) 成功。
- 原发布run [41285840378](https://github.com/kejilion/KPanel/actions/runs/41285840378) / attempt1的候选CI、L3、main/tag、主线CI、Release成功，但首次公开浏览器失败，外层conclusion=failure、preview-completion跳过，原件见workflow-failed.json。唯一受限只读复核 [41290512546](https://github.com/kejilion/KPanel/actions/runs/41290512546) / attempt1成功，原snapshot/镜像/标签/正文/附件未重建或覆盖；新验证controller为fd7033717f6f47d66277da760e7d255bd85d3c25。三个controller资源的original/recipe摘要与原轮逐字节相同，实际runner镜像ID另记；只说明新隔离环境验收通过，不证明环境根因或原缺陷已修复。
- L3唯一入口 run-release-l3.mjs --github-hosted，回执candidate `fd7033717f6f47d66277da760e7d255bd85d3c25`、run/attempt 41285840378/1，2026-11-01T07:43:09.143Z 至 2026-11-01T07:59:00.982Z，status=passed。
- Runner ubuntu24/20261026.354.1；go version go1.27.3 linux/amd64、Nodev24.22.1、Docker28.0.4；回执entrySha256（内层scripts/run-release-gate.sh，外层入口仍为run-release-l3.mjs） `43e57f4fd321846204020d6cf27eb75961f0aefa1284a9187e786601704c2d40`。源码/场景/镜像/生命周期及image_e2e通过，源码身份和完整闭包独立回读。
- 1256个官方firstPartySources；还原Store/version并移除新增测试后精确重现旧1255文件基线。相对失败rc.2仅Store方法体、新并发测试和Go版本字串变化，新source `1093d537602653c43fa5a8521c3e730f8fb5af8db5a8326a4b53932008567136`，相对上一公开rc.1的Docker仅脚本pin/checksum变化，相对失败rc.2的Docker完全不变。模块图、GoFiles/导入和六平台生产闭包、37 findings/dispositions及2026-11-08有效期保留；资格重绑定本身不算扫描成功，CI/L3另有真实结果。
- 发布说明草稿/最终均由共享生成器和运行时解析器核对；公开摘要5条、升级提示2条。零digest的本地draft仅作生成器测试，不作为已发布证据。

## 依赖与维护

本候选依赖报告生成于2026-11-01T07:27:55.408Z，新report作业41284813899成功，必需检测源9/9；350个稳定通道信号=38直接/基座行动项+312传递归属信号。首轮相同输入报告（06:05:14.595Z）结果相同，原件保留。EOL current，下次2027-01-29；TS7与@typesNode26既有有期例外保持，不为清空信号机械升级。只采用本轮兼容脚本；上游main额外WAF变化留后续独立评估。最近每日安全审计run41186235799/job124493612780成功（原日志保存），本候选CI/L3另行重验当前图。每日审计与版本发现不同，push的security-advisories按原规则skipped未计为本轮通过。

## Linux、浏览器与未覆盖场景

local-wsl-dr/candidate-validation用途检查通过。真实Linux install文件写入、八语言首次/重复/完整性/失败停止通过；固定Nginx镜像 `sha256:3d6cab9c2941df2441024a2f02a65ebe1b8916c3cbe7b9527f62627accea8004` 三次实际nginx -t均通过，fixture保留WAF include/状态配置。下载以固定本地模板夹具替换；未声称真实ModSecurity/Fail2Ban迁移、全LDNMP/宿主安装或换源已完成。每次--rm、network-none、资源限额，最终own containers=0。

较宽调优smoke在当前WSL的SSH-service夹具失败；旧5b8327f同样失败，原日志script-smoke/tuning-diagnostic/tuning-baseline保留，未通过重跑或降低断言改记为绿。此项与本轮修复的回归及正式CI/L3分别记录。

本轮没有新启动rc.3本地mock。原aa977678 manifest及准备版的HTTPS组件/测试字节与rc.3完全相同，复用相同组件输入的14/13/12px、窄宽视口、空IP提示及键盘3px测量；Web包版本树不同，不声称整树相同。A/AAAA成功和接口失败由11项单元测试覆盖。前序已停止manifest保留；未执行真实DNS绑定或TLS写入。

公开GitHub-hosted真实Panel/Agent四组浏览器initial48/restart26/published45/upgrade31共150项通过，errors=0、assetFailures=0；真实发布信息接口/更新弹窗的版本、digest、摘要及升级提示一致。测试库存/CA明确为夹具，不挂宿主Docker/systemd socket，不等于真实宿主安装。无新soak需求；物理ARM64、全宿主/systemd/OpenRC更新、认证登录、200%专项视觉和跨发行版真实安装仍未覆盖。FastAdmin保留前次精确应用函数的Linux证据及PHP74维护风险，不重建数据库测试环境。

Runner实际ID `sha256:8cbc94f41f3b13ce4d8d305fe284b87d2232df78e6d05fbdcaa35f0950d0c45d`；recipe与controller原件摘要一致、productImageUnchanged=true，三份cleanup回执确认移除测试容器/网络。

## 公开产物与更新通道

版本/preview index `sha256:4d6ed51757f6d4d9f91981a1e4c376265f7ca269c16d97d1e5e82f65f4aade64`；amd64 `sha256:57b389bf17375650a54310ed4208a75c844aff8e1ae4bbcd9d9bd1854eabf608`；arm64 `sha256:08d6508c631fd4ed4b7ac1f0ac673fce74e00f7ccee5c343e16d8fd52667b151`。19附件名称/大小/digest与冻结清单、公开API及下载磁盘回读匹配；15条SHA256SUMS、modified component provenance通过。

公开正文SHA-256 `f21eb18ec2e56f2e7dae3359e29d5915c3dec3dbd91d273b4a9bded9ebb6e6f3`，snapshot `d5a18c60356f107e427e3afca4614463319ffab28131d31521b70be713654beb`，附件清单 `5368dc5aed8e9774b386ac918889b87beb5f3992b5eb5377e8105594154f53f3`；controller `aa9776784569dc3a9bc88b0b6d64cc98d545bc91`。GitHub Latest仍v1.28.0，Docker latest仍 `sha256:bd90d2bf69f059b4da29e456054bfb1785b47a6b1d453e5941e299a4a7de78c4`；前一preview冻结快照为v1.29.0-rc.1 / `sha256:655e84cc25d5a6343ca932f548e462fa6984d3bd1b0a2a2e6e08b06ccea45f12`，不可变产物保留。

加入预览只切来源并检查，不自动安装；自动安装与一次性安装独立、退出不自动降级，旧状态默认stable。生产全部不适用（预览版禁止生产部署）；prod-108及其他生产机器本轮未连接、检查、部署或升级。人工回退需选固定digest并恢复配套备份，不以退出通道代替降级。

## 来源归档与本地收尾

- fix/https-dns-guide / `050c7a0a48ccc8ca6e7490e8e4b7689591d17186` → `refs/heads/archive/fix/https-dns-guide`，纳入v1.29.0-rc.3的范围及差异映射已核对。
- fix/rc2-dns-readable-code / `37e0a0339a9a3c0c06141d9b75becbbeea012e2b` → `refs/heads/archive/fix/rc2-dns-readable-code`，纳入v1.29.0-rc.3的范围及差异映射已核对。
- fix/ai-sqlite-writer-20261101 / `ed71dea8d20bb37d9cc80a0fcca30c73740071f7` → `refs/heads/archive/fix/ai-sqlite-writer-20261101`，纳入v1.29.0-rc.3的范围及差异映射已核对。

原HTTPS工作树/分支按未知原始所有权保留；自己字号修复及脚本兼容分支仅按已核实所有权本地归档，sh远端仍只有main写入。当前RC列车保留，历史64项未决分支原样保留，归档不等于所有原tip内容原样发布。验收文档候选需同SHA候选Linux CI→主线快进→主线CI成功后归档，最终精确引用由closeout.json保存，不再次构建产品。

## 交付节奏与流程异常

<!-- kpanel-release-metrics:start -->
- 首个纳入提交时间：2026-10-30T00:22:15+08:00
- 候选冻结时间：2026-11-01T07:26:32.476Z
- 生产完成时间：不适用
- 提交到生产用时：不适用
- 是否回滚、紧急热修复或重复发布：否
- 若发生失败，发现时间、恢复时间和逃逸门禁：不适用
<!-- kpanel-release-metrics:end -->

首个源码时间取本轮脚本前置ed58dea原始提交；RC不进入正式发布/生产部署频率及生产lead time。复用既有release-metrics滚动14天/最近20正式版本口径，不把缺失生产事实当零耗时。新RC是聚合独立缺陷/指引更新，不重发同一版本产物。

权威报告生成于2026-11-01T08:41:29.540Z：滚动14天正式标签4个、分布3天、单日最多2，有生产完成证据0次。最近20个正式版本验收覆盖100%，生产完成覆盖45%（9份），生产lead time中位数6.82小时（仅9份有值），变更失败1/20；缺失生产记录不推断为从未部署。

<!-- kpanel-release-process-metrics:start -->
- 已记录发布流程异常或无效证据拦截次数：26
- 其中生产写操作开始后异常次数：0
<!-- kpanel-release-process-metrics:end -->

<!-- kpanel-release-process-incidents:start -->
[
  {
    "fingerprint": "preparation/source-qualification/helper-array-syntax",
    "count": 1,
    "position": "before-production-write",
    "impact": "The external script identity helper initially failed JavaScript parsing due to a missing closing array bracket; no identity success receipt or external write resulted from that invocation.",
    "recoveryEvidence": "Original tool transcript contains SyntaxError Unexpected token ';' at verify-script-identity.mjs:20. The single bracket was corrected; helper then verified exact upstream patch IDs, unchanged WAF/FastAdmin bytes and publication-main tree equality. Current/final supplemental receipts remain in the preparation evidence root.",
    "permanentAction": "Validate local helper syntax before executing qualification; retain original transcript and do not infer successful qualification from a later unrelated shell command.",
    "historicalReleases": []
  },
  {
    "fingerprint": "preparation/review-trailers/separate-message-paragraphs",
    "count": 1,
    "position": "before-production-write",
    "impact": "Metadata commit f060f544 separated each trailer into a Git message paragraph; completion checker reported OCR missing because Git recognized only the final block.",
    "recoveryEvidence": "Immutable local history retains f060f544 and de124adf. A new metadata-only commit used one contiguous trailer block from --file; exact checker then reported OCR and independent review recorded. No source tests were falsely rebound or shared history rewritten.",
    "permanentAction": "Use a message file with one contiguous final trailer block and inspect canonical candidate-checkpoint output before freezing.",
    "historicalReleases": []
  },
  {
    "fingerprint": "preparation/preview/missing-self-app-fixture-option",
    "count": 1,
    "position": "before-production-write",
    "impact": "The first mock preview used the default inventory without KPanel installed, so the intended DNS domain guide journey was unavailable despite services being ready.",
    "recoveryEvidence": "Original preview/manifest.json in the preparation root is preserved stopped. Repository-supported KPANEL_MOCK_SELF_APP=1 enabled the intended scenario in separate preview-self-app and final manifests. DOM proof observed the DNS empty-IP guide and retained input; no real app installation or host write occurred.",
    "permanentAction": "Set the documented mock self-app option during scenario preparation; distinguish service readiness from visible journey completion.",
    "historicalReleases": []
  },
  {
    "fingerprint": "preparation/browser-measurement/unsupported-evaluator-global",
    "count": 1,
    "position": "before-production-write",
    "impact": "A read-only DOM measurement used parseFloat, which was unavailable in the browser evaluator scope, so that measurement returned TypeError instead of evidence.",
    "recoveryEvidence": "Original CUA tool transcript preserves TypeError parseFloat is not a function. The supported read-only query returned computed font-size strings directly; actual final guide sizes were 14/13/12px and narrow guide/client scroll widths both320, page widths both380. Browser viewport reset and all own tabs/previews closed.",
    "permanentAction": "Return supported DOM values from evaluator and parse in host code when needed; record measurement failures rather than treating them as pass.",
    "historicalReleases": []
  },
  {
    "fingerprint": "preflight/orchestrator-cli/unsupported-help",
    "count": 1,
    "position": "before-production-write",
    "impact": "Preparation invoked local-feature-preview.mjs --help, which is unsupported and returned usage failure before any preview startup.",
    "recoveryEvidence": "Original tool transcript retained; supported start/status/stop arguments were read from the canonical workflow and used successfully.",
    "permanentAction": "Read the existing workflow/CLI parser for unsupported help; do not repeat known invalid invocation or interpret it as preview success.",
    "historicalReleases": [
      "v1.15.0",
      "v1.14.0"
    ]
  },
  {
    "fingerprint": "preparation/file-discovery/assumed-literal-path",
    "count": 5,
    "position": "before-production-write",
    "impact": "Two source-discovery calls guessed a scan-policy location or inventory helper filenames that do not exist; neither produced validation evidence.",
    "recoveryEvidence": "Original rg errors retained in tool transcript. Located actual scripts/torrent-backport-scan-scope.json and torrent-backport-baseline.mjs, reused known qualification helpers, and reproduced the official 1255-file inventory. During SQLite follow-up, one search again guessed two nonexistent inventory helper filenames, and one publication-artifact inspection guessed the reusable workflow name before reading its actual caller reference; rg --files identified the existing torrent-backport-baseline.mjs, then canonical identity was independently reproduced. No guessed-path validation was accepted.",
    "permanentAction": "Use rg --files before content searches and reuse verified project paths; inspect each result rather than the final shell exit alone.",
    "historicalReleases": []
  },
  {
    "fingerprint": "preparation/file-discovery/wrong-repository",
    "count": 1,
    "position": "before-production-write",
    "impact": "One exploratory source-helper search ran in the paired sh worktree instead of KPanel, returning missing files; no gate was run or marked passed.",
    "recoveryEvidence": "Original errors retained; source qualification then used the exact KPanel absolute import and independently reproduced its baseline.",
    "permanentAction": "Bind repository and workdir in the command spec before cross-repository source discovery.",
    "historicalReleases": []
  },
  {
    "fingerprint": "preparation/file-discovery/powershell-unexpanded-glob",
    "count": 4,
    "position": "before-production-write",
    "impact": "Two exploratory rg calls supplied shell-style wildcard path arguments that PowerShell did not expand, causing literal-path errors. This shares the cause of historical rollback-research/rg/windows-glob; no verification pass was inferred.",
    "recoveryEvidence": "Original tool transcripts retained. Followed with focused rg --files and known exact script paths; all actual source/preview checks used canonical entries.",
    "permanentAction": "Pass rg -g patterns or known exact filenames; do not rely on PowerShell glob expansion for native-tool path arguments.",
    "historicalReleases": [
      "v1.15.1",
      "v1.16.0"
    ]
  },
  {
    "fingerprint": "qualification/l3-evidence/inline-powershell-quoting",
    "count": 1,
    "position": "before-production-write",
    "impact": "A local schema-inspection prototype used nested quotes in node -e through PowerShell and failed parsing before reading any evidence; remote release jobs were unaffected.",
    "recoveryEvidence": "Original tool transcript preserves SyntaxError Invalid string escape. Moved the exact inspection to inspect-l3-schema.mjs, ran node --check, then successfully read only selected inventory identity fields from the prior immutable L3 artifact.",
    "permanentAction": "Use a file-based Node entry for evidence parsing with JSON marker strings; syntax-check before executing and never infer validation from an unrelated final shell command.",
    "historicalReleases": [
      "v1.27.2"
    ]
  },
  {
    "fingerprint": "qualification/l3-receipt/wrong-entry-file-expectation",
    "count": 1,
    "position": "before-production-write",
    "impact": "Local read-only L3 verifier initially compared entrySha256 to run-release-l3.mjs instead of the inner run-release-gate.sh recorded by the producer. It rejected the receipt before writing a success record; no remote job or product input changed.",
    "recoveryEvidence": "Original assertion actual43e57f versus expected9f8cdf is retained in tool transcript; failed verifier preserved as verify-l3-before-entry-fix.mjs. Read scripts/github-release-l3.mjs producer line50, matched candidate Git blob scripts/run-release-gate.sh, then verified exact candidate/run/attempt, all1255 first-party records, inventory hash and image_e2e.",
    "permanentAction": "Derive receipt field semantics from its canonical producer and compare canonical Git source bytes; preserve failed helper and never change a product or receipt to satisfy an incorrect local assumption.",
    "historicalReleases": []
  },
  {
    "fingerprint": "qualification/linux-test-entry/argument-line-ending",
    "count": 2,
    "position": "before-production-write",
    "impact": "Initial Linux binary baseline invocation was rejected by flag parsing; the first external shell wrapper then retained CRLF in timeout180s and was also rejected. Neither ran a valid test or produced a test-pass receipt.",
    "recoveryEvidence": "Original first output showed flag provided but not defined -test; next wrapper reported invalid timeout180s\\r. Preserved final normalized LF shell entry ran original100 iterations successfully; new semantic baseline regression separately reproduced actual517/5 on first run.",
    "permanentAction": "Use a file-based bounded entry with explicit UTF-8 LF, separate flag/value argv and saved exit status; do not count usage output as a test.",
    "historicalReleases": []
  },
  {
    "fingerprint": "preflight/tool-qualification/go-path-omitted",
    "count": 1,
    "position": "before-production-write",
    "impact": "Source task readiness initially reported missing Go because one fresh PowerShell process omitted the known per-call toolchain PATH.",
    "recoveryEvidence": "Original ready statusfailed is preserved in tool transcript. Repeated ready with explicit Go1.27.3 PATH passed; subsequent source tests and rc.3 handoff qualified actual candidate.",
    "permanentAction": "Prefix every Go-dependent fresh invocation with exact verified toolchain bin and inspect preflight JSON status, not only final shell exit.",
    "historicalReleases": []
  },
  {
    "fingerprint": "qualification/review-metadata/empty-amend-flag",
    "count": 1,
    "position": "before-production-write",
    "impact": "Amending an unpushed metadata-only review commit without --allow-empty was rejected; its HEAD remained unchanged.",
    "recoveryEvidence": "Git reported No changes and requested --allow-empty. Root inspected unchanged head and applied --amend --allow-empty only to the local source metadata; source code tree unchanged. Final released review annotation is a later contiguous trailer block.",
    "permanentAction": "When an intentional metadata-only review commit is amended before publication, use --allow-empty explicitly and verify resulting exact commit.",
    "historicalReleases": []
  },
  {
    "fingerprint": "qualification/reviewer-identity/unexposed-model-assumption",
    "count": 1,
    "position": "before-production-write",
    "impact": "A review metadata trailer initially named a concrete inherited model whose runtime identifier was not exposed; this was unsupported attribution, not a code or scanner failure.",
    "recoveryEvidence": "Before candidate push the final review metadata and security verification explicitly corrected reviewer to codex/inherited-default, exact model/reasoning unknown. Raw security report and reviewed snapshot remain unchanged.",
    "permanentAction": "Record only actual explicit override or exposed runtime identity; otherwise state inherited-default and uncertainty. Do not infer model from available-model catalog.",
    "historicalReleases": []
  },
  {
    "fingerprint": "publication/run-discovery/eventual-consistency",
    "count": 1,
    "position": "before-production-write",
    "impact": "The first read-only workflow-run list immediately after successful dispatch did not yet return the new run, so strict identity binding rejected zero matches. A premature monitor call then refused missing publicationRun. No second dispatch occurred.",
    "recoveryEvidence": "CLI returned original run41285840378. Direct API and a later list independently matched full candidatefd703371, controlleraa977678, main workflow_dispatch, created07:42:15Z and attempt1; only that original run was bound. Dispatch request file prevents duplicate dispatch.",
    "permanentAction": "Use the original CLI run URL and validate candidate/controller/event/path/attempt; treat delayed list visibility as pending and gate monitoring on verified identity, never redispatch to repair evidence visibility.",
    "historicalReleases": []
  },
  {
    "fingerprint": "qualification/ci-log-summary/ansi-controls-not-stripped",
    "count": 1,
    "position": "before-production-write",
    "impact": "Local summary extraction initially matched the colored raw Vitest log without stripping terminal controls; it rejected the known successful332/3350/6 counters and wrote no success summary. Remote CI remained successful.",
    "recoveryEvidence": "Original helper preserved as summarize-candidate-tests-before-ansi-fix.mjs; original log untouched. Node stripVTControlCharacters is applied only to the parsing copy, and bounded boolean assertions now verify exact counters and AI/native Windows success before writing candidate-test-summary.json.",
    "permanentAction": "Normalize terminal controls for text summaries, keep original bytes/hash, and use bounded diagnostics rather than dump a whole multi-MB log on a failed match.",
    "historicalReleases": []
  },
  {
    "fingerprint": "qualification/bulk-run-log-download/transport-timeout",
    "count": 1,
    "position": "before-production-write",
    "impact": "Optional whole-run ZIP retrieval timed out after180s with partial bytes and wrote no completed ZIP or success receipt.",
    "recoveryEvidence": "Exact failed-public-job log was independently downloaded in full, and original L3/snapshot/publication/public artifacts and workflow/job statuses are retained. No publication or browser job was retried to repair the optional bulk download.",
    "permanentAction": "Keep bounded per-failure raw logs and immutable artifacts as primary evidence; retrieve optional bulk archives only with an appropriate transfer budget and preserve incomplete status.",
    "historicalReleases": []
  },
  {
    "fingerprint": "public-browser/terminal-first-load/missing-failure-diagnostics",
    "count": 1,
    "position": "before-production-write",
    "impact": "The original published-mode UI assertion failed before a terminal tab became visible. The verifier closed the browser before preserving the failing URL/network/DOM/screenshot, leaving root cause inconclusive; the assertion itself is a separate unclassified quality failure.",
    "recoveryEvidence": "Original run41285840378/attempt1 remainsFAIL with8 published checks and full log/report; initial48/restart26 and public assets passed. One canonical read-only recovery run41290512546/attempt1 passed150 checks on a fresh runner with identical3 controller resource bytes, original SHA/digest/snapshot and unchanged assertions. This is additional environment evidence, not a fix or root-cause closure.",
    "permanentAction": "Owner KPanel release maintainers; due2026-11-03. Any recurrence blocks further fresh-host retries; add bounded original-browser failure diagnostics under a separately reviewed governance candidate, then diagnose and repair. Original failure and needs-validation remain open.",
    "historicalReleases": []
  }
]
<!-- kpanel-release-process-incidents:end -->

失败rc.2 run41280745365/attempt1在main/tag提升后被Verify source拦截：event_hub_test.go:277 SQLite517；保留首次日志与标签，没有Release/公开镜像。旧生产代码新增16并发回归首轮11/16成功并复现517/5；修复写前读升级冲突后Linux200轮、AI整包、Windows20轮、独立取消/retry探针及本rc.3源码/主线/发布门禁通过；公开门禁由唯一受限只读复核补齐，原失败保留。不重跑rc.2碰绿、不改原断言、不覆盖标签；该产品缺陷为生产前质量门禁捕获，未计入生产失败或执行流程异常。

本rc.3首轮公开验收在published模式首次goto终端页后，等待terminal-tab__close可见30秒超时，尚未执行reload；原initial48/restart26及published前8项PTY/bootstrap/API检查通过，但后续published37项与upgrade31项当轮未执行。Panel/Agent健康、未OOM、资源充足；终端/auth/router/fixture源码与rc.1未变。无原页面URL/网络/失败截图，根因INCONCLUSIVE，不能推断上游瞬态或同一Agent502。独立只读复核确认该证据边界。按6.2.2只做一次新canonical runner实验，完整real-feature/published/upgrade全部断言保持；新轮150项通过，原FAIL与needs-validation仍保留。负责人KPanel release maintainers，2026-11-03前定位；任何同指纹再现停止fresh-host追加尝试，先补原浏览器失败诊断并修复入口。诊断allowlist扩展属于另行治理候选，本轮未改发布工作流或断言。

最近五个稳定记录的原始指纹已读取并保留引用；前次Windows首次启动超时及Agent版本信息502上游原因仍未确认，不能因本轮通过关闭。若同指纹再现先诊断，不继续换runner重试。当前质量/流程失败以原日志及明细为准，不把产品断言、额外历史夹具失败、生产前流程拦截混算生产变更失败。

全部本轮mock进程停止；临时Nginx容器0，WSL shell fixture由trap清除、Go临时目录按测试teardown回收。管理树、当前RC候选、原HTTPS未知所有权树、sh主克隆、唯一证据及归属不明数据卷保留；本任务已归档工作树须在最终文档主线CI通过后逐项核对，再记录回收或跳过原因。精确路径、恢复点、C盘实测结束值和净变化由closeout.json保存；未做全局prune、缩VHD或其他任务清理。复用现有工作流，未新增知识/治理流程。
