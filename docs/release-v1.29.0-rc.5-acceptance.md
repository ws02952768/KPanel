# KPanel v1.29.0-rc.5 发布验收记录

日期：2026-11-02；发布级别 L3；releaseChannel=preview；releaseTrain=1.29.0。

候选/标签：`d52cd0b073b53000a1a682d1addf69ba937c58be` / `v1.29.0-rc.5`。[公开预览 Release](https://github.com/kejilion/KPanel/releases/tag/v1.29.0-rc.5)，ID 430712455，draft=false、prerelease=true。产物已发布，生产未部署。

稳定默认/回滚点：v1.28.0，`sha256:bd90d2bf69f059b4da29e456054bfb1785b47a6b1d453e5941e299a4a7de78c4`；上一预览 v1.29.0-rc.3，`sha256:4d6ed51757f6d4d9f91981a1e4c376265f7ca269c16d97d1e5e82f65f4aade64`。唯一列车 `release/v1.29.0-candidate` 保留；不归档单次 RC，不覆盖旧标签或不可变镜像。

## 发布画像与范围

- 域名绑定在设置与应用市场显示提交、完成、最近日志及错误反馈；选中应用间不串任务。反代准备 Python/flock 依赖，后续步骤失败即停止，KPanel 绑定后保留直连。
- Panel 与 SSH 共用 /usr/local/bin/k；安装/更新同步校验过的配套脚本到 /root/kejilion.sh 与 /usr/local/bin/k，/usr/bin/k 指向后者。更新会替换自定义脚本，需先备份；继承许可、区域和统计偏好；首次升级按当前SSH设置双向重新继承，global入口缺失时回退可信旧专用脚本。
- 首次从旧生命周期升级时，新 Agent 等待原锁/事务结束后迁移；业务写入暂缓，健康与更新恢复控制保持可用。旧自动更新单元暂用小型受保护转发文件，下一次正常更新移除。迁移失败恢复原内容、权限、链接或缺失状态；已提交迁移清理旧快照后才恢复写入。
- 保留 FastAdmin PHP74、LinuxMirrors/Armbian、command install 与 AI SQLite 并发修复；采用公开脚本的 WAF 来源修正、LDNMP 卸载反馈及工具调优。未增加 API endpoint、数据 Schema、端口、Compose 格式或 Agent 权限；新增迁移期间的认证后写入就绪检查。风险 L3 来自实际发布与宿主脚本写入。
- 来源 aedef1fb、7c7503d2、2155e4b8；首次升级缺口由发布者在 df55aa82 修复。独立复核发现 FREE-1（已提交 wrapper 后残留旧回退快照），de809509/c8158eb8 修正并回归；416b944 记录复核。RC4 e91已形成不可变标签，但发布run41316174091在公开前因偏好继承缺陷被取消，没有RC4公开Release；ad79b025修正旧专用脚本CN/关闭统计覆盖当前SSH V6/开启统计的问题，d52cd0b0记录本轮独立复核，`d52cd0b073b53000a1a682d1addf69ba937c58be` 是同树复核收据。
- 原79a0候选虽曾CI通过，但未发布且因真实旧升级复现撤回；原repro与独立FAIL均保留，不能用旧成功结果替代新候选。未纳入其余历史未决分支；不召回已交付写任务。

## 跨仓库联动及来源

scriptLinkageState=coupled，变更集 proxy-shared-host-script-preferences-preview-20261101。镜像固定公开 `kejilion/sh@d2156444add457973d4a84634653aa10cece712f`，根SHA-256 `f60a7023ca0d6dd9ee5f3557669675b84ca7e2c0f3de9fe4b8da22e740478235`，CN `d05a3dc5eb33916ce6cd1e377d08f363cbe144ce41298e719eb0d91fd69998a3`。WAF固定1ad9ad8e57c83317d4fe1f418054195f653e79cb，七个公开文件逐一核对。脚本已先公开于main，本轮未再次写入sh/apps/nginx仓库；非KPanel仓库没有推送功能或归档分支。
- FastAdmin 与 LinuxMirrors/command install 的实际函数字节与上一兼容脚本核对相同；前序FastAdmin真实Linux PHP74/MySQL/Nginx证据复用，不重建测试数据库。应用市场公开旧kpanel.conf的首次升级行为已用不可变RC3生命周期复现并纳入canonical回归；没有靠修改公共apps模板规避已安装的旧生命周期。

## 多维质量结论

| 维度 | 状态 | 证据与边界 |
| --- | --- | --- |
| 业务正确性与双端互通 | 已验证 | typed调用、依赖失败停止、共享入口与偏好继承；精确RC3升级→真实新Agent迁移→后续正常更新。真实外网反代/ACME、全LDNMP与换源未覆盖 |
| 网络入侵与供应链安全 | 已验证 | 固定来源/摘要、同SHA CI/L3扫描、双架构来源、SBOM/provenance；37项既有发现/处置保留，不宣称零漏洞或全审计 |
| 稳定性、失败恢复与兼容 | 已验证 | 真实flock/未提交事务拒绝、危险权限拒绝、最终rename失败恢复、提交后清理失败/重启、下一次失败更新恢复当前版本；备份/中断/回滚矩阵。真实服务崩溃与跨发行版仍有边界 |
| 性能与资源预算 | 已实现未实机验证 | 单工作者、30秒单次超时，仅锁忙每秒有界轮询；非忙失败保持关闭且不盲重试。没有吞吐或长期宿主压力基准，不承诺性能提升 |
| 用户体验与可访问性 | 已验证 | 新冻结mock重做设置成功、证书失败无切换、OpenList反馈、Enter、无控制台错误、14px/21px status；390px/既有截图仅按web/src精确相等复用。未新测全主题/语言/200% |
| 数据、配置与迁移 | 已验证 | 不变数据库格式；只迁移共享脚本和兼容入口，保留现有事务及冷备份。真实宿主业务数据/自定义脚本仍需备份；不把预览验证当稳定版真机准入 |

## 自动门禁及独立复核

- 可信控制器e91的push CI [run 41313999645](https://github.com/kejilion/KPanel/actions/runs/41313999645)通过；主线CI来自已取消RC4发布 [run 41316174091](https://github.com/kejilion/KPanel/actions/runs/41316174091) 的独立成功子作业124881204992/124881204997（精确源码e91），父run整体取消不记为成功。RC5冻结候选push CI [run 41321406097](https://github.com/kejilion/KPanel/actions/runs/41321406097)，Linux/Windows同SHA成功。原生Windows SCM安装、SYSTEM/账户隔离、签名遥测、实际stop/start与真实磁盘apply/rollback通过；未验证OS重启，也不关闭历史首次启动/Agent502根因。
- canonical发布 [run 41322833168](https://github.com/kejilion/KPanel/actions/runs/41322833168) / attempt1：候选CI、完整L3、main/tag、最终主线CI、Release、双架构与公开验收全部有真实成功回执。
- L3唯一外层run-release-l3.mjs --github-hosted；2026-11-01T15:46:23.088Z 至 2026-11-01T16:04:44.511Z，github-hosted，ubuntu24/20261026.354.1；go version go1.27.3 linux/amd64、Nodev24.22.1、Docker28.0.4。入口SHA-256 `181efee78ef2f8c197bf34c90a074550286acebc7a10630635386494b57c2fa3`；image_e2e、旧生命周期与真实锁/提交后回滚回归实际执行并通过。
- canonical1258源文件 `1f1a8b4a581b70e655d6778f35750b8d589b34700272df69dd0778554f39f542`；六目标actual/original闭包均相等，closure=`25a7547c1ea2576476553cf9cf5b2d681253e4a9ffa6ec2dafd0a68d4c6e586f`。Linux Agent新增一个runtime Go文件；模块图、111生产模块、vendor/构建摘要、37处置及2026-10-25/2026-11-08窗口未变。来源绑定本身不算扫描成功，CI/L3是另行实际结果。
- Web332文件、3351通过、6项既有skip；本机受影响Go包Linux全测通过。Windows完整Agent包4项文件操作失败在原基线同样复现，原日志保留，不计为通过；发布范围Windows Node的原生门禁另行通过。
- 三次独立范围分别为初始源码/恢复、最终元数据；恢复和最终为PASS WITH FOLLOW-UP，FREE-1原始HIGH/FAIL已修复，未覆盖原件。同provider干净会话fallback provider-unavailable，其他provider无可用CLI。旧RC4 OCR恢复10/10及2人工排除、元数据4/4及2人工排除，本轮RC5 5/5及3人工排除，均先自由臂后约束臂，新增H0/M0/L0。
- 候选普通安全盲审两包54个runtime Go精确复用快照，新恢复增量独立检查；H0/M1/L0，1既有中级事项单独登记负责人/11-15期限。相对v1.28.0四组cmd/kejilion-agent、internal/agent、internal/ai、internal/appmarket均recorded；历史无scope的skip trailer警告保留，不改写历史。quality-audit稳定补漏/CF深审未执行，不宣称全面覆盖。
- 本稳定周期OCR样本含既有有效/跳过/不可量化观察；本轮原始自由臂发现1条正确性HIGH、约束新增0，修正后复核0。仅如实记录本候选，不提前作工具有效性或退出判定。

## 依赖与技术栈

- 复用前序rc.3依赖发现报告（2026-11-01T07:27:55.408Z，9/9源，350信号=38直接/基座+312传递归属）；模块/工具链/lock依赖无变化，仅发布版本元数据变化。不伪称本RC新跑了dependency-report。Go1.27.3、Node24.22.1及现有固定Action/镜像继续使用，只有配套受管脚本来源更新。
- 2026-11-01安全通告/EOL既有证据保留；EOL复核下次2027-01-29，TS7/@typesNode26有期例外不变。本候选CI/L3实际重验Go、Node、源及镜像安全，不把日常发现报告当扫描结果。

## 隔离Linux及浏览器

- local-wsl-dr仅用于候选验证；新测试使用固定Bash容器、只读源码挂载且无宿主Docker/systemd socket。旧生命周期文件来自不可变fd703371；生产迁移CLI与真实flock、故障文件写入、权限/链接恢复被实际执行。服务/Docker外部动作在生命周期夹具中明确mock，不称真实systemd/OpenRC全宿主安装。
- 本机治理366/366通过，canonical six-target输入和原始27条记录摘要/退出码逐项回读。默认mock缺self-app场景、CRLF/路径/fixture/source作用域等首轮失败分别保留，修正后才计对应有效结果。
- 公开GitHub-hosted真实Panel/Agent initial48、restart26、published45、upgrade31，共150项通过，errors=0、assetFailures=0；真实发布信息接口和更新弹窗的版本、digest、8条摘要及3条升级提示匹配。测试库存/CA是夹具，不是生产宿主；runner ID `sha256:a01d7d5ee1f22449b564bf37ed4920eea32c13796c6dc55e373ee55d695afebe`。三个controller资源原始/recipe摘要一致且productImageUnchanged=true，三份cleanup回执通过。
- 本地mock绑定最终commit；截图接口两种形式均不可用，没有新截图主张，只复用精确相同web/src的前序截图/窄屏证据。旧版首开终端缺tab原失败与唯一只读复核仍保留，当前新预览通过不等于旧根因已修复；2026-11-03前复核，再现前先诊断，不追加盲目fresh-host重试。
- 没有新增soak需求；物理ARM64、真实主机反代/ACME、全LDNMP、完整换源、真实服务进程崩溃、认证登录及完整视觉矩阵未覆盖。PHP74维护风险、较宽调优smoke在旧新基线同样失败、Windows历史启动根因与AI历史总量预算待办均保留。

## 公开产物与自更新通道

- GitHub Release正文与冻结RELEASE_NOTES.md摘要相同，19附件/15校验项全部公开回读；SBOM/provenance及KPANEL_MODIFIED_COMPONENTS身份一致。版本与preview OCI index=`sha256:14e6020e825169d1ce479ce6b155b8252dc5621f29e689643559ecad9e8facd9`。
- linux/amd64：`sha256:e00d87af6c0d52b86e67c6937b3b5d887df9f3e7f731a063d1a958cd69bd4749`。
- linux/arm64：`sha256:bdf22795835dbb1b5805fe3948524fe50109cf5e7fd3f9763414d00491f84cce`。
- GitHub Latest和Docker latest仍为v1.28.0/`sha256:bd90d2bf69f059b4da29e456054bfb1785b47a6b1d453e5941e299a4a7de78c4`；上一preview `sha256:4d6ed51757f6d4d9f91981a1e4c376265f7ca269c16d97d1e5e82f65f4aade64`保持不可变回滚。加入预览只切换来源/检查，不自动安装；自动安装与一次性立即安装独立，退出预览不降级。
- CLI/systemd手动/自动、OpenRC手动更新按现有事务/冷备份/回滚矩阵验证，OpenRC自动调度及Node边界不变。真实宿主完整升级与回退须在稳定版准入前另验；当前公开验收不代替此项。

## 来源归档与本地资源

| 原分支 | 精确tip | 远端归档（已回读） | 本地处置 |
| --- | --- | --- | --- |
| `fix/proxy-install-feedback-20261101` | `aedef1fb60324d749153f374d25825ced7a7bd18` | `refs/heads/archive/fix/proxy-install-feedback-20261101` | 作者工作树/忽略依赖/在用预览保留，本地待处置 |
| `refactor/shared-host-script-20261101` | `7c7503d2d83c300a1f9b7dc385a96e21ae5aa16f` | `refs/heads/archive/refactor/shared-host-script-20261101` | 作者工作树/忽略依赖/在用预览保留，本地待处置 |
| `fix/sync-host-script-on-panel-update-20261101` | `2155e4b8c514c636f93867976c2dd26ec7cbc334` | `refs/heads/archive/fix/sync-host-script-on-panel-update-20261101` | 作者工作树/忽略依赖/在用预览保留，本地待处置 |
| `fix/rc4-trusted-automation-20261101` | `1b4bf5b945700bd2c5be77de3c518f466b1cbdd1` | `refs/heads/archive/fix/rc4-trusted-automation-20261101` | 本任务工作树按清理回执处置 |
| `fix/rc4-shared-upgrade-boundary-20261101` | `416b94439ccf5bd42e99ff29367d3d3fa2fbef39` | `refs/heads/archive/fix/rc4-shared-upgrade-boundary-20261101` | 本任务工作树按清理回执处置 |
| `fix/rc5-legacy-script-preferences-20261101` | `d52cd0b073b53000a1a682d1addf69ba937c58be` | `refs/heads/archive/fix/rc5-legacy-script-preferences-20261101` | 本任务工作树按清理回执处置 |

- 六个archive ref精确保存，四个来源原为local-only，两个本发布者远端分支以expected-SHA lease原子保存并移出活跃区；均可由`v1.29.0-rc.5`恢复。唯一RC列车保留。
- 本任务RC4临时Linux三目录及RC5测试副本已移除，自有mock预览已停止。Windows自有工作树清理回执：C:\GitHub\_release-admission\kpanel-v129-rc4: removed；C:\GitHub\_release-fixes\kpanel-rc4-shared-upgrade: removed；C:\GitHub\_release-fixes\kpanel-rc5-legacy-preferences: removed。作者三工作树、忽略依赖与在用预览按所有权保留，本地待处置，不作为下一轮未发布候选。
- C盘本次自有来源工作树清理阶段实际净可用空间变化 239538176 字节；临时Linux副本的宿主卷变化另见清理回执。只报告实测净变化，不将源码逻辑大小当释放量，不执行WSL磁盘压缩；并发系统写入可能影响卷余量。
- 原始失败、截图/trace、当前及上一稳定恢复证据保留；不执行全局Docker/cache prune、不缩减共享WSL、不清理未知卷或他人预览。验收工作树与其分支在本文件同SHA候选/主线CI完成后另行归档回收，外部closeout回读为准，不提前声称其已删除。

## 生产与回滚

- 不适用（预览版禁止生产部署）。本发布任务未连接prod-108、154或kpanel-eu，未执行生产备份、升级或postdeploy。来源任务此前独立授权的154直连恢复不是本候选部署，不复用为本版生产验收。
- 源码/tag与上一预览/稳定不可变镜像保留；兼容回退须配套恢复Panel/Agent及更新前数据，宿主自定义脚本用先前保存的原件。退出预览不会自动降级；生产实际版本/健康不作未执行采样。公共默认通道继续稳定v1.28.0。

## 交付节奏数据

<!-- kpanel-release-metrics:start -->
- 首个纳入提交时间：2026-11-01T18:09:32+08:00
- 候选冻结时间：2026-11-01T15:28:26.118Z
- 生产完成时间：不适用
- 提交到生产用时：不适用
- 是否回滚、紧急热修复或重复发布：否
- 若发生失败，发现时间、恢复时间和逃逸门禁：不适用
<!-- kpanel-release-metrics:end -->

<!-- kpanel-release-process-metrics:start -->
- 已记录发布流程异常或无效证据拦截次数：25
- 其中生产写操作开始后异常次数：0
<!-- kpanel-release-process-metrics:end -->

已记录的是准备/执行器/夹具/取证通道问题；首次升级真实缺陷、FREE-1与RC4偏好继承缺陷属于正常质量拦截；RC4公开前主动取消和新RC5不可变版本另行记录，Windows既有测试失败另记。原件与重测分别保留，无生产写入。

<!-- kpanel-release-process-incidents:start -->
[
  {
    "fingerprint": "preparation/local-validation/image-pull-cancelled",
    "count": 1,
    "position": "before-production-write",
    "impact": "本地L2镜像准备受网络拉取影响被取消，未作为完整通过证据。",
    "recoveryEvidence": "准备目录保留local-L2/构建原件；当前精确GitHub CI通过。",
    "permanentAction": "本轮完整发布只用已批准GitHub-hosted唯一入口；不以局部WSL验证替代L3。",
    "historicalReleases": []
  },
  {
    "fingerprint": "preparation/preview/missing-self-app-fixture-option",
    "count": 2,
    "position": "before-production-write",
    "impact": "两次默认mock未启用已安装KPanel场景，预期设置表单不可用。",
    "recoveryEvidence": "准备目录及corrected-preview保留首轮stopped manifest；带KPANEL_MOCK_SELF_APP=1的新manifest绑定精确提交并通过旅程。",
    "permanentAction": "本轮已固定场景环境并核验真实表单；发布维护者下次同场景启动先检查fixture，2026-11-03复核。",
    "historicalReleases": []
  },
  {
    "fingerprint": "preparation/browser-capture/screenshot-backend-unavailable",
    "count": 2,
    "position": "before-production-write",
    "impact": "两个已文档化截图接口失败，没有生成本轮新截图。",
    "recoveryEvidence": "旧preview/browser-result.json和工具原件；改为声明DOM/字体证据边界，只复用精确相同web/src的既有截图。",
    "permanentAction": "上游接口例外：KPanel发布维护者2026-11-03复核，截图后端恢复前不宣称新截图；不可替代正式公开浏览器证据。",
    "historicalReleases": []
  },
  {
    "fingerprint": "preparation/browser-selector/incorrect-alert-target",
    "count": 2,
    "position": "before-production-write",
    "impact": "旧候选预览的URL前缀与全页多alert使两次定位无效。",
    "recoveryEvidence": "旧preview/browser-result.json保留两项；改用已观察的HTTPS区域alert，未延长或削弱断言。",
    "permanentAction": "本轮最终预览使用精确区域及角色定位，成功/失败/禁止切换均回读。",
    "historicalReleases": []
  },
  {
    "fingerprint": "qualification/task-evidence/timestamp-precision",
    "count": 1,
    "position": "before-production-write",
    "impact": "原任务契约7位小数时间戳不符合入口要求。",
    "recoveryEvidence": "local-input-qualification.json失败保留，-2及corrected-local-qualification.json合规通过。",
    "permanentAction": "证据时间使用ISO毫秒；不改写原始高精度日志。",
    "historicalReleases": []
  },
  {
    "fingerprint": "evidence/job-log/in-progress-unavailable",
    "count": 1,
    "position": "before-production-write",
    "impact": "运行中作业日志连接器返回不可用，未取得可用证据。",
    "recoveryEvidence": "工具原件404保留；完成后通过GitHub job logs取得全文并校验摘要及27条rawEvidence。",
    "permanentAction": "仅在作业完成后下载不可变日志，不以连接器404推断作业失败或通过。",
    "historicalReleases": []
  },
  {
    "fingerprint": "evidence/log-download/direct-network-stall",
    "count": 1,
    "position": "before-production-write",
    "impact": "直连下载只获得缓慢不完整日志并停止。",
    "recoveryEvidence": "candidate-linux-direct-partial.log及停止收据保留；已配置用户代理取得完整日志。",
    "permanentAction": "本轮只读下载显式使用现有本机代理，TLS和发布源保持正常验证。",
    "historicalReleases": []
  },
  {
    "fingerprint": "cleanup/owned-download/utc-identity-guard",
    "count": 1,
    "position": "before-production-write",
    "impact": "停止自有下载进程的首轮时间身份比较因本地时区不匹配而拒绝。",
    "recoveryEvidence": "原工具拒绝原件保留；精确PID、命令行及UTC创建时间再验证后停止自有进程。",
    "permanentAction": "使用明确UTC字符串校验，禁止按进程名批量停止。",
    "historicalReleases": []
  },
  {
    "fingerprint": "validation/legacy-rootfs/missing-deploy-fixture",
    "count": 1,
    "position": "before-production-write",
    "impact": "首个旧升级复现夹具未提供IPC部署脚本而失败。",
    "recoveryEvidence": "upgrade-boundary/repro.log保留；补入精确deploy文件后repro-2.log复现真实产品缺陷。",
    "permanentAction": "夹具纳入目标源码固定依赖；产品缺陷与夹具缺失分别记录。",
    "historicalReleases": []
  },
  {
    "fingerprint": "validation/migration-shell/wrong-platform-launch",
    "count": 1,
    "position": "before-production-write",
    "impact": "一次原拟语法检查的Linux迁移命令被作为Windows Bash脚本调用，无合格通过证据。",
    "recoveryEvidence": "工具调用及migration-command.sh原件保留；随后Linux bash-n及显式隔离rootfs验证。",
    "permanentAction": "使用Linux文件启动器和rootfs，语法检查显式bash-n；不复用Windows调用作为Linux结果。",
    "historicalReleases": []
  },
  {
    "fingerprint": "validation/rootfs-launch/crlf-test-argument",
    "count": 1,
    "position": "before-production-write",
    "impact": "生成的CRLF启动器使测试flag带尾随CR，在测试前被拒绝。",
    "recoveryEvidence": "upgrade-boundary/fix-rootfs.log保留；LF文件启动器后续执行。",
    "permanentAction": "本轮Linux启动器固定UTF8/LF文件，避免内联跨shell转义。",
    "historicalReleases": []
  },
  {
    "fingerprint": "validation/rootfs-image/missing-bash-alias",
    "count": 1,
    "position": "before-production-write",
    "impact": "固定Bash容器缺少产品约定/bin/bash别名，测试未启动。",
    "recoveryEvidence": "fix-rootfs-2.log保留；只在一次性测试容器建立正确别名后-3通过。",
    "permanentAction": "canonical harness显式创建容器内/bin/bash别名；不改变宿主机。",
    "historicalReleases": []
  },
  {
    "fingerprint": "validation/governance-launch/missing-go-path",
    "count": 1,
    "position": "before-production-write",
    "impact": "本机治理预检因Go未加入该进程PATH而失败，没有执行测试。",
    "recoveryEvidence": "upgrade-boundary/governance.log保留；明确Go1.27.3 PATH后的governance-2.log 366/366。",
    "permanentAction": "每个本机入口显式固定Go工具链路径。",
    "historicalReleases": []
  },
  {
    "fingerprint": "validation/linux-go-launch/inline-path-quoting",
    "count": 1,
    "position": "before-production-write",
    "impact": "内联跨shell PATH含空格括号导致启动器语法失败，测试未运行。",
    "recoveryEvidence": "linux-full-affected-go.log保留；文件启动器-2完成两包全测。",
    "permanentAction": "本轮Linux命令使用LF脚本文件；禁止将Windows PATH插值为shell代码。",
    "historicalReleases": []
  },
  {
    "fingerprint": "qualification/recovery-helper/mixed-module-syntax",
    "count": 1,
    "position": "before-production-write",
    "impact": "恢复来源辅助器混用require与顶层await失败，后续本地提交曾保留旧绑定。",
    "recoveryEvidence": "原工具错误和de809提交保留；c815在外部发布前修正ESM辅助器、测试夹具和来源绑定。",
    "permanentAction": "当前辅助器固定ESM import，提交/推送前逐条检查原生命令退出码和canonical来源；无共享历史改写。",
    "historicalReleases": []
  },
  {
    "fingerprint": "validation/recovery-test/lifecycle-source-scope",
    "count": 1,
    "position": "before-production-write",
    "impact": "新增后续回退测试曾在函数外source含local的生命周期，夹具失败。",
    "recoveryEvidence": "harness-2.log保留；c815将source放回函数作用域，harness-4完整通过。",
    "permanentAction": "回归在canonical CI/L3使用相同production入口与required函数调用约定。",
    "historicalReleases": []
  },
  {
    "fingerprint": "validation/linux-overlay/stale-policy-context",
    "count": 1,
    "position": "before-production-write",
    "impact": "WSL测试副本尚未带新policy时，overlay patch上下文不符，未跑测试。",
    "recoveryEvidence": "harness-3.log保留；-4仅应用精确测试源；另在clean候选独立执行canonical六目标资格。",
    "permanentAction": "WSL副本只声明候选验证，不充当精确Git/CI/L3；新候选以官方来源导出核对。",
    "historicalReleases": []
  },
  {
    "fingerprint": "qualification/source-inventory/missing-output-directory",
    "count": 1,
    "position": "before-production-write",
    "impact": "canonical inventory首轮输出目录不存在而失败。",
    "recoveryEvidence": "fix-inventory.log保留；创建独占目录后fix-inventory-2.log通过。",
    "permanentAction": "本轮每个新inventory使用独立已创建目录，不覆盖失败原件。",
    "historicalReleases": []
  },
  {
    "fingerprint": "freeze/review-receipt/wrong-result-field",
    "count": 1,
    "position": "before-production-write",
    "impact": "冻结辅助器错误读取coverage字段，守卫在任何推送前拒绝。",
    "recoveryEvidence": "原工具AssertionError保留；按实际overallResult及reviewedScopeVerdict、head和计数严格校验后冻结。",
    "permanentAction": "读取真实复核schema并核对四项覆盖、零发现和冻结head，不绕过失败守卫。",
    "historicalReleases": []
  },
  {
    "fingerprint": "qualification/rc5-ready/missing-validation-plan",
    "count": 1,
    "position": "before-production-write",
    "impact": "RC5首次Ready契约缺少验证计划，被入口拒绝，未写源码。",
    "recoveryEvidence": "ready.json失败及ready-2.json通过原件；随后才修改并独立复核。",
    "permanentAction": "RC5写入前补齐实际focused-go/source-inventory计划，不改写失败回执。",
    "historicalReleases": []
  },
  {
    "fingerprint": "evidence/helper-path-normalization/powershell-inline-escaping",
    "count": 1,
    "position": "before-production-write",
    "impact": "一次外部辅助器路径规范化内联命令因PowerShell引号失败，未修改文件或执行验收。",
    "recoveryEvidence": "原SyntaxError保留；改为独立ESM文件统一正斜杠路径。",
    "permanentAction": "后续外部辅助器只用文件参数和显式路径，不使用嵌套内联代码。",
    "historicalReleases": []
  },
  {
    "fingerprint": "evidence/copied-helper/windows-path-literal",
    "count": 1,
    "position": "before-production-write",
    "impact": "监控辅助器复制时将Windows原生路径写入JS字符串，读取本地证据失败；GitHub发布已单次成功触发且未重发。",
    "recoveryEvidence": "identify-publication首次ENOENT保留；五个辅助器规范为正斜杠路径后识别原run。",
    "permanentAction": "生成JS代码时先规范路径，禁止将原生反斜杠作为未转义源码替换值；监控失败仅恢复取证。",
    "historicalReleases": []
  }
]
<!-- kpanel-release-process-incidents:end -->

## 遗留风险与后续准入

- 既有本机生命周期锁中级待办和staging清理低级待办：KPanel lifecycle maintainers，2026-11-15；既有进度脱敏/备份导出父路径线索同日核实。AI历史总量预算中级待办截至2026-11-08；旧终端首开截至11-03，Windows启动根因截至11-07。细节只在仓库外安全待办保留。
- 自定义脚本会被同步替换；先备份。部分清理或损坏快照会保持业务写关闭，可能需人工恢复；不承诺所有异常都自动继续。全宿主/跨发行版与物理ARM64需独立后续准入，不据预览通过宣称稳定版普遍适用。
- 本轮没有新增知识或工作流；复用release-kpanel4.3、官方OCR、普通候选盲审、版本管理和安全资源清理规则。公开Release、本验收及实际workflow/artifact/Git refs共同构成交付事实。
