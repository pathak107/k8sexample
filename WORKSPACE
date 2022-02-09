load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "d6b2513456fe2229811da7eb67a444be7785f5323c6708b38d851d2b51e54d83",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.30.0/rules_go-v0.30.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.30.0/rules_go-v0.30.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains(version = "1.17.6")

gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_proto",
    sha256 = "66bfdf8782796239d3875d37e7de19b1d94301e8972b3cbd2446b332429b4df1",
    strip_prefix = "rules_proto-4.0.0",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

rules_proto_dependencies()

rules_proto_toolchains()

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:weqSxi/TMs1SqFRMHCtBgXRs8k3X39QIDEZ0pRcttUg=",
    version = "v1.44.0",
)

go_repository(
    name = "io_k8s_sigs_json",
    importpath = "sigs.k8s.io/json",
    sum = "h1:fD1pz4yfdADVNfFmcP2aBEtudwUQ1AlLnRBALr33v3s=",
    version = "v0.0.0-20211020170558-c049b76a60c6",
)

go_repository(
    name = "io_k8s_client_go",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/client-go",
    sum = "h1:23QYUmCQ/W6hW78xIwm3XqZrrKZM+LWDqW2zfo+szJs=",
    version = "v0.23.3",
)

go_repository(
    name = "io_k8s_apimachinery",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/apimachinery",
    sum = "h1:7IW6jxNzrXTsP0c8yXz2E5Yx/WTzVPTsHIx/2Vm0cIk=",
    version = "v0.23.3",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:Hsa8mG0dQ46ij8Sl2AYJDUv1oA9/d6Vk+3LG99Oe02g=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "io_k8s_klog",
    importpath = "k8s.io/klog",
    sum = "h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=",
    version = "v1.0.0",
)

go_repository(
    name = "io_k8s_klog_v2",
    importpath = "k8s.io/klog/v2",
    sum = "h1:bUO6drIvCIsvZ/XFgfxoGFQU/a4Qkh0iAlvUR7vlHJw=",
    version = "v2.30.0",
)

go_repository(
    name = "io_k8s_api",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/api",
    sum = "h1:KNrME8KHGr12Ozjf8ytOewKzZh6hl/hHUZeHddT3a38=",
    version = "v0.23.3",
)

go_repository(
    name = "com_github_shirou_gopsutil_v3",
    importpath = "github.com/shirou/gopsutil/v3",
    sum = "h1:XB/+p+kVnyYLuPHCfa99lxz2aJyvVhnyd+FxZqH/k7M=",
    version = "v3.21.4",
)

go_repository(
    name = "com_github_armon_circbuf",
    importpath = "github.com/armon/circbuf",
    sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
    version = "v0.0.0-20150827004946-bbbad097214e",
)

go_repository(
    name = "com_github_armon_go_metrics",
    importpath = "github.com/armon/go-metrics",
    sum = "h1:8GUt8eRujhVEGZFFEjBj46YV4rDjvGrNxb0KMWYkL2I=",
    version = "v0.0.0-20180917152333-f0300d1749da",
)

go_repository(
    name = "com_github_armon_go_radix",
    importpath = "github.com/armon/go-radix",
    sum = "h1:BUAU3CGlLvorLI26FmByPp2eC2qla6E1Tw+scpcg/to=",
    version = "v0.0.0-20180808171621-7fddfc383310",
)

go_repository(
    name = "com_github_bketelsen_crypt",
    importpath = "github.com/bketelsen/crypt",
    sum = "h1:+0HFd5KSZ/mm3JmhmrDukiId5iR6w4+BdFtfSy4yWIc=",
    version = "v0.0.3-0.20200106085610-5cbc8cc4026c",
)

go_repository(
    name = "com_github_chai2010_gettext_go",
    importpath = "github.com/chai2010/gettext-go",
    sum = "h1:7aWHqerlJ41y6FOsEUvknqgXnGmJyJSbjhAWq5pO4F8=",
    version = "v0.0.0-20160711120539-c6fed771bfd5",
)

go_repository(
    name = "com_github_google_nftables",
    importpath = "github.com/google/nftables",
    sum = "h1:bCm0Cf+suMHiri9F+ss5n5W0AVas85K5Z0Hekgpe7N0=",
    version = "v0.0.0-20210916140115-16a134723a96",
)

go_repository(
    name = "com_github_avast_retry_go",
    importpath = "github.com/avast/retry-go",
    sum = "h1:4SOWQ7Qs+oroOTQOYnAHqelpCO0biHSxpiH9JdtuBj0=",
    version = "v3.0.0+incompatible",
)

go_repository(
    name = "com_github_mdlayher_netlink",
    importpath = "github.com/mdlayher/netlink",
    sum = "h1:W3er9pI7mt2gOqOWzwvx20iJ8Akiqz1mUMTxU6wdvl8=",
    version = "v0.0.0-20191009155606-de872b0d824b",
)

go_repository(
    name = "com_github_koneu_natend",
    importpath = "github.com/koneu/natend",
    sum = "h1:MFX8DxRnKMY/2M3H61iSsVbo/n3h0MWGmWNN1UViOU0=",
    version = "v0.0.0-20150829182554-ec0926ea948d",
)

go_repository(
    name = "com_github_daviddengcn_go_colortext",
    importpath = "github.com/daviddengcn/go-colortext",
    sum = "h1:uVsMphB1eRx7xB1njzL3fuMdWRN8HtVzoUOItHMwv5c=",
    version = "v0.0.0-20160507010035-511bcaf42ccd",
)

go_repository(
    name = "com_github_docker_distribution",
    importpath = "github.com/docker/distribution",
    sum = "h1:a5mlkVzth6W5A4fOsS3D2EO5BUmsJpcB+cRlLU7cSug=",
    version = "v2.7.1+incompatible",
)

go_repository(
    name = "com_github_dvyukov_go_fuzz",
    importpath = "github.com/dvyukov/go-fuzz",
    sum = "h1:e1v4V9Heb+c4xQCCONROFvlzNs6Gq8aRZRwt+WzSEqY=",
    version = "v0.0.0-20201127111758-49e582c6c23d",
)

go_repository(
    name = "com_github_exponent_io_jsonpath",
    importpath = "github.com/exponent-io/jsonpath",
    sum = "h1:105gxyaGwCFad8crR9dcMQWvV9Hvulu6hwUh4tWPJnM=",
    version = "v0.0.0-20151013193312-d6023ce2651d",
)

go_repository(
    name = "com_github_fatih_camelcase",
    importpath = "github.com/fatih/camelcase",
    sum = "h1:hxNvNX/xYBp0ovncs8WyWZrOrpBNub/JfaMvbURyft8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_form3tech_oss_jwt_go",
    importpath = "github.com/form3tech-oss/jwt-go",
    sum = "h1:7ZaBxOI7TMoYBfyA3cQHErNNyAWIKUMIwqxEtgHOs5c=",
    version = "v3.2.3+incompatible",
)

go_repository(
    name = "com_github_fvbommel_sortorder",
    importpath = "github.com/fvbommel/sortorder",
    sum = "h1:dSnXLt4mJYH25uDDGa3biZNQsozaUWDSWeKJ0qqFfzE=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_go_ole_go_ole",
    importpath = "github.com/go-ole/go-ole",
    sum = "h1:nNBDSCOigTSiarFpYE9J/KtEA1IOW4CNeqT9TQDqCxI=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_golangplus_bytes",
    importpath = "github.com/golangplus/bytes",
    sum = "h1:7xqw01UYS+KCI25bMrPxwNYkSns2Db1ziQPpVq99FpE=",
    version = "v0.0.0-20160111154220-45c989fe5450",
)

go_repository(
    name = "com_github_golangplus_fmt",
    importpath = "github.com/golangplus/fmt",
    sum = "h1:f5gsjBiF9tRRVomCvrkGMMWI8W1f2OBFar2c5oakAP0=",
    version = "v0.0.0-20150411045040-2a5d6d7d2995",
)

go_repository(
    name = "com_github_golangplus_testing",
    importpath = "github.com/golangplus/testing",
    sum = "h1:KhcknUwkWHKZPbFy2P7jH5LKJ3La+0ZeknkkmrSgqb0=",
    version = "v0.0.0-20180327235837-af21d9c3145e",
)

go_repository(
    name = "com_github_hashicorp_consul_api",
    importpath = "github.com/hashicorp/consul/api",
    sum = "h1:BNQPM9ytxj6jbjjdRPioQ94T6YXriSopn0i8COv6SRA=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_hashicorp_consul_sdk",
    importpath = "github.com/hashicorp/consul/sdk",
    sum = "h1:LnuDWGNsoajlhGyHJvuWW6FVqRl8JOTPqS6CPTsYjhY=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix",
    importpath = "github.com/hashicorp/go-immutable-radix",
    sum = "h1:AKDB1HM5PWEA7i4nhcpwOrO2byshxBjXVn/J/3+z5/0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_msgpack",
    importpath = "github.com/hashicorp/go-msgpack",
    sum = "h1:zKjpN5BK/P5lMYrLmBHdBULWbJ0XpYR+7NGzqkZzoD4=",
    version = "v0.5.3",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    sum = "h1:iVjPR7a6H0tWELX5NxNe7bYopibicUzc7uPribsnS6o=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_net",
    importpath = "github.com/hashicorp/go.net",
    sum = "h1:sNCoNyDEvN1xa+X0baata4RdcpKwcMS6DH+xwfqPgjw=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_hashicorp_go_rootcerts",
    importpath = "github.com/hashicorp/go-rootcerts",
    sum = "h1:Rqb66Oo1X/eSV1x66xbDccZjhJigjg0+e82kpwzSwCI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_sockaddr",
    importpath = "github.com/hashicorp/go-sockaddr",
    sum = "h1:GeH6tui99pF4NJgfnhp+L6+FfobzVW3Ah46sLo0ICXs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_syslog",
    importpath = "github.com/hashicorp/go-syslog",
    sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_uuid",
    importpath = "github.com/hashicorp/go-uuid",
    sum = "h1:fv1ep09latC32wFoVwnqcnKJGnMSdBanPczbHAYm1BE=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_hashicorp_logutils",
    importpath = "github.com/hashicorp/logutils",
    sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_mdns",
    importpath = "github.com/hashicorp/mdns",
    sum = "h1:WhIgCr5a7AaVH6jPUwjtRuuE7/RDufnUvzIr48smyxs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_memberlist",
    importpath = "github.com/hashicorp/memberlist",
    sum = "h1:EmmoJme1matNzb+hMpDuR/0sbJSUisxyqBGG676r31M=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_hashicorp_serf",
    importpath = "github.com/hashicorp/serf",
    sum = "h1:YZ7UKsJv+hKjqGVUUbtE3HNj79Eln2oQ75tniF6iPt0=",
    version = "v0.8.2",
)

go_repository(
    name = "com_github_liggitt_tabwriter",
    importpath = "github.com/liggitt/tabwriter",
    sum = "h1:9TO3cAIGXtEhnIaL+V+BEER86oLrvS+kWobKpbJuye0=",
    version = "v0.0.0-20181228230101-89fcab3d43de",
)

go_repository(
    name = "com_github_lithammer_dedent",
    importpath = "github.com/lithammer/dedent",
    sum = "h1:VNzHMVCBNG1j0fh3OrsFRkVUwStdDArbgBWoPAffktY=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_makenowjust_heredoc",
    importpath = "github.com/MakeNowJust/heredoc",
    sum = "h1:sjQovDkwrZp8u+gxLtPgKGjk5hCxuy2hrRejBTA9xFU=",
    version = "v0.0.0-20170808103936-bb23615498cd",
)

go_repository(
    name = "com_github_miekg_dns",
    importpath = "github.com/miekg/dns",
    sum = "h1:9jZdLNd/P4+SfEJ0TNyxYpsK8N4GtfylBLqtbYN1sbA=",
    version = "v1.0.14",
)

go_repository(
    name = "com_github_mitchellh_cli",
    importpath = "github.com/mitchellh/cli",
    sum = "h1:iGBIsUe3+HZ/AD/Vd7DErOt5sU9fa8Uj7A2s1aggv1Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_ps",
    importpath = "github.com/mitchellh/go-ps",
    sum = "h1:i6ampVEEF4wQFF+bkYfwYgY+F/uYJDktmvLPf7qIgjc=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_testing_interface",
    importpath = "github.com/mitchellh/go-testing-interface",
    sum = "h1:fzU/JVNcaqHQEcVFAKeR41fkiLdIPrefOvVG1VZ96U0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_wordwrap",
    importpath = "github.com/mitchellh/go-wordwrap",
    sum = "h1:6GlHJ/LTGMrIJbwgdqdl2eEH8o+Exx/0m8ir9Gns0u4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_gox",
    importpath = "github.com/mitchellh/gox",
    sum = "h1:lfGJxY7ToLJQjHHwi0EX6uYBdK78egf954SQl13PQJc=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_mitchellh_iochan",
    importpath = "github.com/mitchellh/iochan",
    sum = "h1:C+X3KsSTLFVBr/tK1eYN/vs4rJcvsiLU338UhYPJWeY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opencontainers_go_digest",
    importpath = "github.com/opencontainers/go-digest",
    sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_pascaldekloe_goe",
    importpath = "github.com/pascaldekloe/goe",
    sum = "h1:Lgl0gzECD8GnQ5QCWA8o6BtfL6mDH5rQgM4/fX3avOs=",
    version = "v0.0.0-20180627143212-57f6aae5913c",
)

go_repository(
    name = "com_github_flosch_pongo2",
    importpath = "github.com/flosch/pongo2",
    sum = "h1:GY1+t5Dr9OKADM64SYnQjw/w99HMYvQ0A8/JoUkxVmc=",
    version = "v0.0.0-20190707114632-bbf5a6c351f4",
)

go_repository(
    name = "com_github_patrickmn_go_cache",
    importpath = "github.com/patrickmn/go-cache",
    sum = "h1:HRMgzkcYKYpi3C8ajMPV8OFXaaRUnok+kx1WdO15EQc=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_posener_complete",
    importpath = "github.com/posener/complete",
    sum = "h1:ccV59UEOTzVDnDUEFdT95ZzHVZ+5+158q8+SJb2QV5w=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_ryanuber_columnize",
    importpath = "github.com/ryanuber/columnize",
    sum = "h1:UFr9zpz4xgTnIE5yIMtWAMngCdZ9p/+q6lTbgelo80M=",
    version = "v0.0.0-20160712163229-9b3edd62028f",
)

go_repository(
    name = "com_github_sean_seed",
    importpath = "github.com/sean-/seed",
    sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
    version = "v0.0.0-20170313163322-e2103e2c3529",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
    version = "v0.0.0-20180927180507-b2de0cb4f26d",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    importpath = "github.com/smartystreets/goconvey",
    sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
    version = "v1.6.4",
)

go_repository(
    name = "com_github_stackexchange_wmi",
    importpath = "github.com/StackExchange/wmi",
    sum = "h1:G0m3OIz70MZUWq3EgK3CesDbo8upS2Vm9/P3FtgI+Jk=",
    version = "v0.0.0-20190523213315-cbe66965904d",
)

go_repository(
    name = "com_github_subosito_gotenv",
    importpath = "github.com/subosito/gotenv",
    sum = "h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_tklauser_go_sysconf",
    importpath = "github.com/tklauser/go-sysconf",
    sum = "h1:HT8SVixZd3IzLdfs/xlpq0jeSfTX57g1v6wB1EuzV7M=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_tklauser_numcpus",
    importpath = "github.com/tklauser/numcpus",
    sum = "h1:ct88eFm+Q7m2ZfXJdan1xYoXKlmwsfP+k88q05KvlZc=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_ugorji_go",
    importpath = "github.com/ugorji/go",
    sum = "h1:j4s+tAvLfL3bZyefP2SEWmhBzmuIlH/eqNuPdFPgngw=",
    version = "v1.1.4",
)

go_repository(
    name = "com_google_cloud_go_firestore",
    importpath = "cloud.google.com/go/firestore",
    sum = "h1:9x7Bx0A9R5/M9jibeJeZWqjeVEIxYW9fZYqB9a70/bY=",
    version = "v1.1.0",
)

go_repository(
    name = "io_k8s_cli_runtime",
    importpath = "k8s.io/cli-runtime",
    sum = "h1:W0/FHdbApnl9oB7xdG643c/Zaf7TZT+43I+zKxwqvhU=",
    version = "v0.20.2",
)

go_repository(
    name = "io_k8s_component_helpers",
    importpath = "k8s.io/component-helpers",
    sum = "h1:UM92JwXHAd2fN/LoJtm0zayIEypCYHxNBKaGhNRT7YE=",
    version = "v0.20.2",
)

go_repository(
    name = "io_k8s_metrics",
    importpath = "k8s.io/metrics",
    sum = "h1:o32EchiH4ukpUg86VLLAgkE9a9Ke0lijkzYxE+wSSRk=",
    version = "v0.20.2",
)

go_repository(
    name = "io_k8s_sigs_kustomize",
    importpath = "sigs.k8s.io/kustomize",
    sum = "h1:JUufWFNlI44MdtnjUqVnvh29rR37PQFzPbLXqhyOyX0=",
    version = "v2.0.3+incompatible",
)

go_repository(
    name = "org_uber_go_goleak",
    importpath = "go.uber.org/goleak",
    sum = "h1:z+mqJhf6ss6BSfSM671tgKyZBFPTTJM+HLxnhPC3wu0=",
    version = "v1.1.10",
)

go_repository(
    name = "org_uber_go_tools",
    importpath = "go.uber.org/tools",
    sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
    version = "v0.0.0-20190618225709-2cfd321de3ee",
)

go_repository(
    name = "com_github_google_go_github_v39",
    importpath = "github.com/google/go-github/v39",
    sum = "h1:rNNM311XtPOz5rDdsJXAp2o8F67X9FnROXTvto3aSnQ=",
    version = "v39.2.0",
)

go_repository(
    name = "com_github_google_go_querystring",
    importpath = "github.com/google/go-querystring",
    sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
    version = "v1.1.0",
)

go_rules_dependencies()

gazelle_dependencies()

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:UoveltGrhghAA7ePc+e+QYDHXrBps2PqFZiHkGR/xK8=",
    version = "v0.0.1-2020.1.4",
)

go_repository(
    name = "co_elastic_go_apm_module_apmhttp",
    importpath = "go.elastic.co/apm/module/apmhttp",
    sum = "h1:k/MjK0y2aLOXumoM8jcWXqxvIFlMS4U8Bn9cMUPdVX0=",
    version = "v1.11.0",
)

go_repository(
    name = "co_elastic_go_apm_module_apmot",
    importpath = "go.elastic.co/apm/module/apmot",
    sum = "h1:Qmol6ztDJgvGK/B2cRdcPRNw4qE7kRv1d0vo9ptZfIo=",
    version = "v1.11.0",
)

go_repository(
    name = "co_elastic_go_apm_module_apmgrpc",
    importpath = "go.elastic.co/apm/module/apmgrpc",
    sum = "h1:E1Cx6rwCj7JKDceI9S7Ix4NEJUW+jyy54Q66ftNh3bA=",
    version = "v1.11.0",
)

go_repository(
    name = "co_elastic_go_apm",
    importpath = "go.elastic.co/apm",
    sum = "h1:uJyt6nCW9880sZhfl1tB//Jy/5TadNoAd8edRUtgb3w=",
    version = "v1.11.0",
)

go_repository(
    name = "co_elastic_go_fastjson",
    importpath = "go.elastic.co/fastjson",
    sum = "h1:3MrGBWWVIxe/xvsbpghtkFoPciPhOCmjsR/HfwEeQR4=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_agnivade_levenshtein",
    importpath = "github.com/agnivade/levenshtein",
    sum = "h1:3oJU7J3FGFmyhn8KHjmVaZCN5hxTr7GxgRue+sxIXdQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_alecholmes_xfccparser",
    importpath = "github.com/alecholmes/xfccparser",
    sum = "h1:/PBnzDBxfHJ66AinLNglzZH4oWLrc1/QTKlSoNNnei8=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_alecthomas_participle",
    importpath = "github.com/alecthomas/participle",
    sum = "h1:P2PJWzwrSpuCWXKnzqvw0b0phSfH1kJo4p2HvLynVsI=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alexflint_go_filemutex",
    importpath = "github.com/alexflint/go-filemutex",
    sum = "h1:AMzIhMUqU3jMrZiTuW0zkYeKlKDAFD+DG20IoO421/Y=",
    version = "v0.0.0-20171022225611-72bdc8eae2ae",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:Hs82Z41s6SdL1CELW+XaDYmOH4hkBN4/N9og/AsOv7E=",
    version = "v0.0.0-20190717042225-c3de453c63f4",
)

go_repository(
    name = "com_github_andreyvit_diff",
    importpath = "github.com/andreyvit/diff",
    sum = "h1:bvNMNQO63//z+xNgfBlViaCIJKLlCJ6/fmUseuG0wVQ=",
    version = "v0.0.0-20170406064948-c7f18ee00883",
)

go_repository(
    name = "com_github_armon_consul_api",
    importpath = "github.com/armon/consul-api",
    sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
    version = "v0.0.0-20180202201655-eb2c6b5be1b6",
)

go_repository(
    name = "com_github_asaskevich_govalidator",
    importpath = "github.com/asaskevich/govalidator",
    sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
    version = "v0.0.0-20190424111038-f61b66f89f4a",
)

go_repository(
    name = "com_github_azure_go_ansiterm",
    importpath = "github.com/Azure/go-ansiterm",
    sum = "h1:w+iIsaOQNcT7OZ575w+acHgRric5iCyQh+xv+KJ4HB8=",
    version = "v0.0.0-20170929234023-d6e3b3328b78",
)

go_repository(
    name = "com_github_azure_go_autorest",
    importpath = "github.com/Azure/go-autorest",
    sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
    version = "v14.2.0+incompatible",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest",
    importpath = "github.com/Azure/go-autorest/autorest",
    sum = "h1:90Y4srNYrwOtAgVo3ndrQkTYn6kf1Eg/AjTFJ8Is2aM=",
    version = "v0.11.18",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_adal",
    importpath = "github.com/Azure/go-autorest/autorest/adal",
    sum = "h1:Mp5hbtOePIzM8pJVRa3YLrWWmZtoxRXqUEzCfJt3+/Q=",
    version = "v0.9.13",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_date",
    importpath = "github.com/Azure/go-autorest/autorest/date",
    sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_mocks",
    importpath = "github.com/Azure/go-autorest/autorest/mocks",
    sum = "h1:K0laFcLE6VLTOwNgSxaGbUcLPuGXlNkbVvq4cW4nIHk=",
    version = "v0.4.1",
)

go_repository(
    name = "in_gopkg_ini_v1",
    importpath = "gopkg.in/ini.v1",
    sum = "h1:duBzk771uxoUuOlyRLkHsygud9+5lrlGjdFBb4mSKDU=",
    version = "v1.62.0",
)

go_repository(
    name = "com_github_azure_go_autorest_logger",
    importpath = "github.com/Azure/go-autorest/logger",
    sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_azure_go_autorest_tracing",
    importpath = "github.com/Azure/go-autorest/tracing",
    sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_basgys_goxml2json",
    importpath = "github.com/basgys/goxml2json",
    sum = "h1:4ln5i4rseYfXNd86lGEB+Vi652IsIXIvggKM/BhUKVw=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_bgentry_speakeasy",
    importpath = "github.com/bgentry/speakeasy",
    sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_blang_semver",
    importpath = "github.com/blang/semver",
    sum = "h1:cQNTCjp13qL8KC3Nbxr/y2Bqb63oX6wdnnjpJbkM4JQ=",
    version = "v3.5.1+incompatible",
)

go_repository(
    name = "com_github_buger_jsonparser",
    importpath = "github.com/buger/jsonparser",
    sum = "h1:y853v6rXx+zefEcjET3JuKAqvhj+FKflQijjeaSv2iA=",
    version = "v0.0.0-20180808090653-f4dd9f5a6b44",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_cespare_xxhash",
    importpath = "github.com/cespare/xxhash",
    sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_chzyer_test",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_cockroachdb_datadriven",
    importpath = "github.com/cockroachdb/datadriven",
    sum = "h1:OaNxuTZr7kxeODyLWsRMC+OD03aFUH+mW6r2d+MWa5Y=",
    version = "v0.0.0-20190809214429-80d97fb3cbaa",
)

go_repository(
    name = "com_github_containernetworking_cni",
    build_file_generation = "on",
    importpath = "github.com/containernetworking/cni",
    sum = "h1:fE3r16wpSEyaqY4Z4oFrLMmIGfBYIKpPrHK31EJ9FzE=",
    version = "v0.7.1",
)

go_repository(
    name = "com_github_containernetworking_plugins",
    importpath = "github.com/containernetworking/plugins",
    sum = "h1:5lnwfsAYO+V7yXhysJKy3E1A2Gy9oVut031zfdOzI9w=",
    version = "v0.8.2",
)

go_repository(
    name = "com_github_coreos_bbolt",
    importpath = "github.com/coreos/bbolt",
    sum = "h1:wZwiHHUieZCquLkDL0B8UhzreNWsPHooDAG3q34zk0s=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    sum = "h1:8F3hqu9fGYLBifCmRCJsicFqDx/D68Rt3q1JMazcgBQ=",
    version = "v3.3.13+incompatible",
)

go_repository(
    name = "com_github_coreos_go_etcd",
    importpath = "github.com/coreos/go-etcd",
    sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_coreos_go_iptables",
    importpath = "github.com/coreos/go-iptables",
    sum = "h1:KH0EwId05JwWIfb96gWvkiT2cbuOu8ygqUaB+yPAwIg=",
    version = "v0.4.2",
)

go_repository(
    name = "com_github_coreos_go_oidc",
    importpath = "github.com/coreos/go-oidc",
    sum = "h1:sdJrfw8akMnCuUlaZU3tE/uYXFgfqom8DBE9so9EBsM=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_coreos_go_systemd",
    importpath = "github.com/coreos/go-systemd",
    sum = "h1:Wf6HqHfScWJN9/ZjdUKyjop4mf3Qdd+1TvvltAvM3m8=",
    version = "v0.0.0-20190321100706-95778dfbb74e",
)

go_repository(
    name = "com_github_coreos_pkg",
    importpath = "github.com/coreos/pkg",
    sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
    version = "v0.0.0-20180928190104-399ea9e2e55f",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man",
    importpath = "github.com/cpuguy83/go-md2man",
    sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
    version = "v1.0.10",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man_v2",
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
    version = "v1.1.9",
)

go_repository(
    name = "com_github_d2g_dhcp4",
    importpath = "github.com/d2g/dhcp4",
    sum = "h1:Xo2rK1pzOm0jO6abTPIQwbAmqBIOj132otexc1mmzFc=",
    version = "v0.0.0-20170904100407-a1d1b6c41b1c",
)

go_repository(
    name = "com_github_d2g_dhcp4client",
    importpath = "github.com/d2g/dhcp4client",
    sum = "h1:suYBsYZIkSlUMEz4TAYCczKf62IA2UWC+O8+KtdOhCo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_d2g_dhcp4server",
    importpath = "github.com/d2g/dhcp4server",
    sum = "h1:+CpLbZIeUn94m02LdEKPcgErLJ347NUwxPKs5u8ieiY=",
    version = "v0.0.0-20181031114812-7d4a0a7f59a5",
)

go_repository(
    name = "com_github_d2g_hardwareaddr",
    importpath = "github.com/d2g/hardwareaddr",
    sum = "h1:itqmmf1PFpC4n5JW+j4BU7X4MTfVurhYRTjODoPb2Y8=",
    version = "v0.0.0-20190221164911-e7d9fbe030e4",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_deepmap_oapi_codegen",
    importpath = "github.com/deepmap/oapi-codegen",
    sum = "h1:Xx8/OynzWhQeKFWr182hg6Ja2evS1gcqdja7qMn05yU=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
    version = "v3.2.0+incompatible",
)

go_repository(
    name = "com_github_dgryski_go_sip13",
    importpath = "github.com/dgryski/go-sip13",
    sum = "h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
    version = "v0.0.0-20181026042036-e10d5fee7954",
)

go_repository(
    name = "com_github_docker_docker",
    importpath = "github.com/docker/docker",
    sum = "h1:w3NnFcKR5241cfmQU5ZZAsf0xcpId6mWOupTvJlUX2U=",
    version = "v0.7.3-0.20190327010347-be7ac8be2ae0",
)

go_repository(
    name = "com_github_docker_go_units",
    importpath = "github.com/docker/go-units",
    sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_docker_spdystream",
    importpath = "github.com/docker/spdystream",
    sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
    version = "v0.0.0-20160310174837-449fdfce4d96",
)

go_repository(
    name = "com_github_docopt_docopt_go",
    importpath = "github.com/docopt/docopt-go",
    sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
    version = "v0.0.0-20180111231733-ee0de3bc6815",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_elastic_go_elasticsearch_v8",
    importpath = "github.com/elastic/go-elasticsearch/v8",
    sum = "h1:/auNW6OUXX8ItIHeHSy5fZXVMld3UW/LndRfqEL/hXg=",
    version = "v8.0.0-20210216171300-6f546f74982b",
)

go_repository(
    name = "com_github_elastic_go_sysinfo",
    importpath = "github.com/elastic/go-sysinfo",
    sum = "h1:FiOJvd3KSHa8ALx/7EPsFcJFsMMhCfgG7NPUZwm3ybk=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_elazarl_goproxy",
    importpath = "github.com/elazarl/goproxy",
    sum = "h1:yUdfgN0XgIJw7foRItutHYUIhlcKzcSf5vDpdhQAKTc=",
    version = "v0.0.0-20180725130230-947c36da3153",
)

go_repository(
    name = "com_github_emicklei_go_restful",
    importpath = "github.com/emicklei/go-restful",
    sum = "h1:H2pdYOb3KQ1/YsqVWoWNLQO+fusocsw354rqGTZtAgw=",
    version = "v0.0.0-20170410110728-ff4f55a20633",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:fP+fF0up6oPY49OrjPrhIJ8yQfdIM85NXMLkMg1EXVs=",
    version = "v0.9.10-0.20210907150352-cf90f659a021",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_evanphx_json_patch",
    importpath = "github.com/evanphx/json-patch",
    sum = "h1:4onqiflcdA9EOZ4RxV643DvftH5pOlLGNtQ5lPWQu84=",
    version = "v4.12.0+incompatible",
)

go_repository(
    name = "com_github_fatih_color",
    importpath = "github.com/fatih/color",
    sum = "h1:DkWD4oS2D8LGGgTQ6IvwJJXSL5Vp2ffcQg58nFV38Ys=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:hsms1Qyu0jgnwNXIxa+/V/PDsU6CfLf6CNO8H7IWoS4=",
    version = "v1.4.9",
)

go_repository(
    name = "com_github_getkin_kin_openapi",
    importpath = "github.com/getkin/kin-openapi",
    sum = "h1:j77zg3Ec+k+r+GA3d8hBoXpAc6KX9TbBPrwQGBIy2sY=",
    version = "v0.76.0",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_globalsign_mgo",
    importpath = "github.com/globalsign/mgo",
    sum = "h1:D4uzjWwKYQ5XnAvUbuvHW93esHg7F8N/OYeBBcJoTr0=",
    version = "v0.0.0-20180905125535-1ca0a4f7cbcb",
)

go_repository(
    name = "com_github_glycerine_goconvey",
    importpath = "github.com/glycerine/goconvey",
    sum = "h1:gclg6gY70GLy3PbkQ1AERPfmLMMagS60DKF78eWwLn8=",
    version = "v0.0.0-20190410193231-58a59202ab31",
)

go_repository(
    name = "com_github_glycerine_go_unsnap_stream",
    importpath = "github.com/glycerine/go-unsnap-stream",
    sum = "h1:Ujru1hufTHVb++eG6OuNDKMxZnGIvF6o/u8q/8h2+I4=",
    version = "v0.0.0-20181221182339-f9677308dec2",
)

go_repository(
    name = "com_github_gobuffalo_envy",
    importpath = "github.com/gobuffalo/envy",
    sum = "h1:X3is06x7v0nW2xiy2yFbbIjwHz57CD6z6MkvqULTCm8=",
    version = "v1.6.5",
)

go_repository(
    name = "com_github_gobuffalo_flect",
    importpath = "github.com/gobuffalo/flect",
    sum = "h1:GPoRjEN0QObosV4XwuoWvSd5uSiL0N3e91/xqyY4crQ=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_godbus_dbus",
    importpath = "github.com/godbus/dbus",
    sum = "h1:RBUpb2b14UnmRHNd2uHz20ZHLDK+SW5Us/vWF5IHRaY=",
    version = "v0.0.0-20180201030542-885f9cc04c9c",
)

go_repository(
    name = "com_github_go_gl_glfw",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
    version = "v0.0.0-20190409004039-e6da0acd62b1",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
    version = "v0.0.0-20200222043503-6f7a984d4dc4",
)

go_repository(
    name = "com_github_go_ini_ini",
    importpath = "github.com/go-ini/ini",
    sum = "h1:7VJT/ZXjzqSrvtraFp4ONq80hTcRQth1c9ZnQ3uNQvU=",
    version = "v1.62.0",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    sum = "h1:wDJmvq38kDhkVxi50ni9ykkdUr1PKgqKOoi01fa0Mdk=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_golang_groupcache",
    importpath = "github.com/golang/groupcache",
    sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
    version = "v0.0.0-20210331224755-41bb18bfe9da",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:jlYHihg//f7RRwuPfptm04yp4s7O6Kw8EZiVYIGcH0g=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    sum = "h1:Qgr9rKW7uDUkrbSmQeiDsGa8SjGyCOGtuasMWwvp2P4=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:MP4Eh7ZCb31lleYCFuwm0oe4/YGak+5l1vA2NOE80nA=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:QK40JKJyMdUDz+h+xvCsru/bJhvG0UxvePV0ufL/AcE=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_go_logr_zapr",
    importpath = "github.com/go-logr/zapr",
    sum = "h1:v6Ji8yBW77pva6NkJKQdHLAJKrIJKRHz0RXwPqCHSR4=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:sjZBwGj9Jlw33ImPtvFviGYvseOtDM7hkSKB7+Tv3SM=",
    version = "v2.0.5",
)
go_repository(
    name = "com_github_googleapis_gnostic",
    importpath = "github.com/googleapis/gnostic",
    sum = "h1:2qsuRm+bzgwSIKikigPASa2GhW8H2Dn4Qq7UxD8K/48=",
    version = "v0.5.3",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:gK4Kx5IaGY9CD5sPJ36FHiBJ6ZXl0kilRiiCj+jdYp4=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:Khx7svrCpmxxtHBq5j2mp/xVjsi8hQMfNLvJFAlrGgU=",
    version = "v0.5.5",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:zIaiqGYDQwa4HVx5wGRTXbx38Pqxjemn4BP98wpzpXo=",
    version = "v0.0.0-20210226084205-cbba55b83ad5",
)

go_repository(
    name = "com_github_google_renameio",
    importpath = "github.com/google/renameio",
    sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:EVhdT+1Kseyi1/pUmXKaFxYsDNy9RQYkMWRH68J/W7Y=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_go_openapi_analysis",
    importpath = "github.com/go-openapi/analysis",
    sum = "h1:8b2ZgKfKIUTVQpTb77MoRDIMEIwvDVw40o3aOXdfYzI=",
    version = "v0.19.5",
)

go_repository(
    name = "com_github_go_openapi_errors",
    importpath = "github.com/go-openapi/errors",
    sum = "h1:a2kIyV3w+OS3S97zxUndRVD46+FhGOUBDFY7nmu4CsY=",
    version = "v0.19.2",
)

go_repository(
    name = "com_github_go_openapi_jsonpointer",
    importpath = "github.com/go-openapi/jsonpointer",
    sum = "h1:gZr+CIYByUqjcgeLXnQu2gHYQC9o73G2XUeOFYEICuY=",
    version = "v0.19.5",
)

go_repository(
    name = "com_github_go_openapi_jsonreference",
    importpath = "github.com/go-openapi/jsonreference",
    sum = "h1:5cxNfTy0UVC3X8JL5ymxzyoUZmo8iZb+jeTWn7tUa8o=",
    version = "v0.19.3",
)

go_repository(
    name = "com_github_go_openapi_loads",
    importpath = "github.com/go-openapi/loads",
    sum = "h1:5I4CCSqoWzT+82bBkNIvmLc0UOsoKKQ4Fz+3VxOB7SY=",
    version = "v0.19.4",
)

go_repository(
    name = "com_github_go_openapi_runtime",
    importpath = "github.com/go-openapi/runtime",
    sum = "h1:csnOgcgAiuGoM/Po7PEpKDoNulCcF3FGbSnbHfxgjMI=",
    version = "v0.19.4",
)

go_repository(
    name = "com_github_go_openapi_spec",
    importpath = "github.com/go-openapi/spec",
    sum = "h1:0XRyw8kguri6Yw4SxhsQA/atC88yqrk0+G4YhI2wabc=",
    version = "v0.19.3",
)

go_repository(
    name = "com_github_go_openapi_strfmt",
    importpath = "github.com/go-openapi/strfmt",
    sum = "h1:eRfyY5SkaNJCAwmmMcADjY31ow9+N7MCLW7oRkbsINA=",
    version = "v0.19.3",
)

go_repository(
    name = "com_github_go_openapi_swag",
    importpath = "github.com/go-openapi/swag",
    sum = "h1:lTz6Ys4CmqqCQmZPBlbQENR1/GucA2bzYTE12Pw4tFY=",
    version = "v0.19.5",
)

go_repository(
    name = "com_github_go_openapi_validate",
    importpath = "github.com/go-openapi/validate",
    sum = "h1:QhCBKRYqZR+SKo4gl1lPhPahope8/RLt6EVgY8X80w0=",
    version = "v0.19.5",
)

go_repository(
    name = "com_github_gophercloud_gophercloud",
    importpath = "github.com/gophercloud/gophercloud",
    sum = "h1:P/nh25+rzXouhytV2pUHBb65fnds26Ghl8/391+sT5o=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:twflg0XRTjwKpxb/jFExr4HGq6on2dEOmnL6FV+fgPw=",
    version = "v0.0.0-20190910122728-9d188e94fb99",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:+/TMaTYc4QFitKJxsQ7Yye35DkWvkdLcvGKqM+x0Ufc=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    importpath = "github.com/gregjones/httpcache",
    sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
    version = "v0.0.0-20180305231024-9cad4c3443a7",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    sum = "h1:THDBEeQ9xZ8JEaCLyLQqXMMdRqNr0QAUJTIkQAUtFjg=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
    sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    build_file_generation = "on",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    #sum = "h1:UImYN5qQ8tuGpGE16ZmjvcTtTw24zw1QAp/SlnNrZhI=",
    #version = "v1.9.5",
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_opentracing",
    importpath = "github.com/grpc-ecosystem/grpc-opentracing",
    sum = "h1:MJG/KsmcqMwFAkh8mTnAwhyKoB+sTAnY4CACC110tbU=",
    version = "v0.0.0-20180507213350-8e809c8a8645",
)

go_repository(
    name = "com_github_hako_durafmt",
    importpath = "github.com/hako/durafmt",
    sum = "h1:BpJ2o0OR5FV7vrkDYfXYVJQeMNWa8RhklZOpW2ITAIQ=",
    version = "v0.0.0-20200710122514-c0fb7b4da026",
)

go_repository(
    name = "com_github_hashicorp_go_retryablehttp",
    importpath = "github.com/hashicorp/go-retryablehttp",
    sum = "h1:92lWxgpa+fF3FozM4B3UZtHZMJX8T5XT+TFdCxsPyWs=",
    version = "v0.6.8",
)

go_repository(
    name = "com_github_hashicorp_go_cleanhttp",
    importpath = "github.com/hashicorp/go-cleanhttp",
    sum = "h1:dH3aiDG9Jvb5r5+bYHsikaOUIpcM0xvgMXVoDkXMzJM=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    sum = "h1:0hERBMJE1eitiLkihrMvRVBYAkpHzc/J3QdDN+dAcgU=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_ianlancetaylor_demangle",
    importpath = "github.com/ianlancetaylor/demangle",
    sum = "h1:mV02weKRL81bEnm8A0HT1/CAelMQDBuQIfLw8n+d6xI=",
    version = "v0.0.0-20200824232613-28f6c0f3b639",
)

go_repository(
    name = "com_github_imdario_mergo",
    importpath = "github.com/imdario/mergo",
    sum = "h1:JboBksRwiiAJWvIYJVo46AfV+IAIKZpfrSzVKj42R4Q=",
    version = "v0.3.5",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_influxdata_influxdb_client_go",
    importpath = "github.com/influxdata/influxdb-client-go/v2",
    sum = "h1:O0CGIuIwQafvAxttAJ/VqMKfbWWn2Mt8rbOmaM2Zj4w=",
    version = "v2.2.2",
)

go_repository(
    name = "com_github_influxdata_line_protocol",
    importpath = "github.com/influxdata/line-protocol",
    sum = "h1:W9WBk7wlPfJLvMCdtV4zPulc4uCPrlywQOmbFOhgQNU=",
    version = "v0.0.0-20200327222509-2487e7298839",
)

go_repository(
    name = "com_github_j_keck_arping",
    importpath = "github.com/j-keck/arping",
    sum = "h1:742eGXur0715JMq73aD95/FU0XpVKXqNuTnEfXsLOYQ=",
    version = "v0.0.0-20160618110441-2cf9dc699c56",
)

go_repository(
    name = "com_github_joeshaw_multierror",
    importpath = "github.com/joeshaw/multierror",
    sum = "h1:rp+c0RAYOWj8l6qbCUTSiRLG/iKnW3K3/QfPPuSsBt4=",
    version = "v0.0.0-20140124173710-69b34d4ec901",
)

go_repository(
    name = "com_github_joho_godotenv",
    importpath = "github.com/joho/godotenv",
    sum = "h1:Zjp+RcGpHhGlrMbJzXTrZZPrWj+1vfm90La1wgB6Bhc=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_jonboulle_clockwork",
    importpath = "github.com/jonboulle/clockwork",
    sum = "h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    importpath = "github.com/jstemmer/go-junit-report",
    sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
    version = "v4.20.0+incompatible",
)

go_repository(
    name = "com_github_juju_errors",
    importpath = "github.com/juju/errors",
    sum = "h1:wnhMXidtb70kDZCeLt/EfsVtkXS5c8zLnE9y/6DIRAU=",
    version = "v0.0.0-20180806074554-22422dad46e1",
)

go_repository(
    name = "com_github_juju_loggo",
    importpath = "github.com/juju/loggo",
    sum = "h1:UUHMLvzt/31azWTN/ifGWef4WUqvXk0iRqdhdy/2uzI=",
    version = "v0.0.0-20190526231331-6e530bcce5d8",
)

go_repository(
    name = "com_github_juju_testing",
    importpath = "github.com/juju/testing",
    sum = "h1:ZUgTbk8oHgP0jpMieifGC9Lv47mHn8Pb3mFX3/Ew4iY=",
    version = "v0.0.0-20190613124551-e81189438503",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:TDTW5Yz1mjftljbcKqRcrYhd4XeOoI98t+9HbQbYf7g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_k8snetworkplumbingwg_network_attachment_definition_client",
    importpath = "github.com/k8snetworkplumbingwg/network-attachment-definition-client",
    sum = "h1:9QbZltQGRFe7temwcTDjj8rIbow48Gv6mIKOxuks+OI=",
    version = "v1.1.1-0.20201119153432-9d213757d22d",
)

go_repository(
    name = "com_github_kelseyhightower_envconfig",
    importpath = "github.com/kelseyhightower/envconfig",
    sum = "h1:Im6hONhd3pLkfDFsbRgu68RDNkGF1r3dvMUtDTo2cv8=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_kevinpollet_nego",
    importpath = "github.com/kevinpollet/nego",
    sum = "h1:BaIpmhcqpBnz4+NZjUjVGxKNA+/E7ovKsjmwqjXcGYc=",
    version = "v0.0.0-20200324111829-b3061ca9dd9d",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_kr_logfmt",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:s5hAObm+yFO5uHYt5dYjxi2rXrsnmRpJx4OYvIWUaQs=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_kubernetes_incubator_reference_docs",
    importpath = "github.com/kubernetes-incubator/reference-docs",
    sum = "h1:3MPc52KuzSwQ8SfQ9WKDuszK8r7XX+oXCeHaQNCUOV8=",
    version = "v0.0.0-20170929004150-fcf65347b256",
)

go_repository(
    name = "com_github_labstack_echo",
    importpath = "github.com/labstack/echo",
    sum = "h1:pGRcYk231ExFAyoAjAfD85kQzRJCRI8bbnE7CX5OEgg=",
    version = "v3.3.10+incompatible",
)

go_repository(
    name = "com_github_labstack_echo_v4",
    importpath = "github.com/labstack/echo/v4",
    sum = "h1:PQIBaRplyRy3OjwILGkPg89JRtH2x5bssi59G2EL3fo=",
    version = "v4.1.17",
)

go_repository(
    name = "com_github_labstack_gommon",
    importpath = "github.com/labstack/gommon",
    sum = "h1:JEeO0bvc78PKdyHxloTKiF8BD5iGrH8T6MSeGvSgob0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_lpar_gzipped_v2",
    importpath = "github.com/lpar/gzipped/v2",
    sum = "h1:y7FjyTH07f8dX0YQ5o0sg2DTbRnmS3oT1pUvxViQ//o=",
    version = "v2.0.2",
)

go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    sum = "h1:ZC2Vc7/ZFkGmsVC9KvOjumD+G5lXy2RtTKyzRKO2BQ4=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:hB2xlXdHp/pmPZq0y3QnmWAArdw9PqbmotexnWx/FU8=",
    version = "v0.0.0-20190626092158-b2ccc519800e",
)

go_repository(
    name = "com_github_markbates_inflect",
    importpath = "github.com/markbates/inflect",
    sum = "h1:5fh1gzTFhfae06u3hzHYO9xe3l3v3nW5Pwt3naLTP5g=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:UVL0vNpWh04HeJXV0KLcaT7r06gOH2l4OW6ddYRUIY4=",
    version = "v0.0.9",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:bnP0vzxcAdeI1zdubAl5PjU6zsERjGZb7raWodagDYs=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:UnlwIPBGaTZfPQ6T1IGzPI0EkYAQmT9fAEJ/poFC63o=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_mattn_go_shellwords",
    importpath = "github.com/mattn/go-shellwords",
    sum = "h1:K/VxK7SZ+cvuPgFSLKi5QPI9Vr/ipOf4C1gN+ntueUk=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:I0XW9+e1XWDxdcEniV4rQAIOPUGDq67JSCiRCgGCZLI=",
    version = "v1.0.2-0.20181231171920-c182affec369",
)

go_repository(
    name = "com_github_microsoft_go_winio",
    importpath = "github.com/Microsoft/go-winio",
    sum = "h1:zoIOcVf0xPN1tnMVbTtEdI+P8OofVk3NObnwOQ6nK2Q=",
    version = "v0.4.11",
)

go_repository(
    name = "com_github_microsoft_hcsshim",
    importpath = "github.com/Microsoft/hcsshim",
    sum = "h1:ZfF0+zZeYdzMIVMZHKtDKJvLHj76XCuVae/jNkjj0IA=",
    version = "v0.8.6",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    importpath = "github.com/mitchellh/go-homedir",
    sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_moby_term",
    importpath = "github.com/moby/term",
    sum = "h1:aY7OQNf2XqY/JQ6qREWamhI/81os/agb2BAGpcx5yWI=",
    version = "v0.0.0-20200312100748-672ec06f55cd",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mschoch_smat",
    importpath = "github.com/mschoch/smat",
    sum = "h1:VeRdUYdCw49yizlSbMEn2SZ+gT+3IUKx8BqxyQdz+BY=",
    version = "v0.0.0-20160514031455-90eadee771ae",
)

go_repository(
    name = "com_github_munnerz_goautoneg",
    importpath = "github.com/munnerz/goautoneg",
    sum = "h1:7PxY7LVfSZm7PEeBTyK1rj1gABdCO2mbri6GKO1cMDs=",
    version = "v0.0.0-20120707110453-a547fc61f48d",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:F9x/1yl3T2AeKLr2AMdilSD8+f9bvMnNN8VS5iDtovc=",
    version = "v0.0.0-20161129095857-cc309e4a2223",
)

go_repository(
    name = "com_github_mxk_go_flowrate",
    importpath = "github.com/mxk/go-flowrate",
    sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
    version = "v0.0.0-20140419014527-cca7078d478f",
)

go_repository(
    name = "com_github_natefinch_lumberjack",
    importpath = "github.com/natefinch/lumberjack",
    sum = "h1:4QJd3OLAMgj7ph+yZTuX13Ld4UpgHp07nNdFX7mqFfM=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_nxadm_tail",
    importpath = "github.com/nxadm/tail",
    sum = "h1:DQuhQpB1tVlglWS2hLQ5OV6B5r8aGxSrPc5Qo6uTN78=",
    version = "v1.4.4",
)

go_repository(
    name = "com_github_nytimes_gziphandler",
    importpath = "github.com/NYTimes/gziphandler",
    sum = "h1:lsxEuwrXEAokXB9qhlbKWPpo3KMLZQ5WB5WLQRW1uq0=",
    version = "v0.0.0-20170623195520-56545f4a5d46",
)

go_repository(
    name = "com_github_oklog_ulid",
    importpath = "github.com/oklog/ulid",
    sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_olekukonko_tablewriter",
    importpath = "github.com/olekukonko/tablewriter",
    sum = "h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_oneofone_xxhash",
    importpath = "github.com/OneOfOne/xxhash",
    sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:2mOpI4JVVPBN+WQRa0WKH2eXR+Ey+uK4n7Zj0aYpIQA=",
    version = "v1.14.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:o0+MgICZLuZ7xjH7Vx6zS/zcu93/BEp1VwkIW1mEXCE=",
    version = "v1.10.1",
)

go_repository(
    name = "com_github_pborman_uuid",
    importpath = "github.com/pborman/uuid",
    sum = "h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    sum = "h1:T5zMGML61Wp+FlcbWjRDT7yAxhJNAiPPLOFECq181zc=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_peterbourgon_diskv",
    importpath = "github.com/peterbourgon/diskv",
    sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
    version = "v2.0.1+incompatible",
)

go_repository(
    name = "com_github_philhofer_fwd",
    importpath = "github.com/philhofer/fwd",
    sum = "h1:UbZqGr5Y38ApvM/V/jEljVxwocdweyH+vmYvRPBnbqQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_pquerna_cachecontrol",
    importpath = "github.com/pquerna/cachecontrol",
    sum = "h1:0XM1XL/OFFJjXsYXlG30spTkV/E9+gmd5GD1w2HE8xM=",
    version = "v0.0.0-20171018203845-0dec1b30a021",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:NTGy1Ja9pByO+xAeH/qiWnLrKtr3hJPNjaVUwnjpdpA=",
    version = "v1.7.1",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
    version = "v0.0.0-20190812154241-14fe0d1b01d4",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:RyRA7RzGXQZiW+tGMr7sxa85G1z0yOpM1qq5c8lNawc=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:wH4vA7pcjKuZzjF7lM8awk4fnuJO6idemZXoKnULUx4=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_prometheus_tsdb",
    importpath = "github.com/prometheus/tsdb",
    sum = "h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=",
    version = "v0.7.1",
)

go_repository(
    name = "com_github_prometheus_operator_apis",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring",
    # sum = "h1:DUQJRh1osK5195HZKed21IobNQdQ7oxRTqb5Yqt8ttw=",
    sum = "h1:OuRwJIT6F/FTh/R2WJRNj2utkBBwo5My5RI+uGDGGPI=",
    version = "v0.44.0",
)

go_repository(
    name = "com_github_puerkitobio_purell",
    importpath = "github.com/PuerkitoBio/purell",
    sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_puerkitobio_urlesc",
    importpath = "github.com/PuerkitoBio/urlesc",
    sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
    version = "v0.0.0-20170810143723-de5bf2ad4578",
)

go_repository(
    name = "com_github_roaringbitmap_roaring",
    importpath = "github.com/RoaringBitmap/roaring",
    sum = "h1:naNqvO1mNnghk2UvcsqnzHDBn9DRbCIRy94GmDTRVTQ=",
    version = "v0.5.5",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:RR9dF3JtopPvtkroDZuVD7qquD0bnHlKSqaQhgwt8yk=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_russross_blackfriday",
    importpath = "github.com/russross/blackfriday",
    sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_safchain_ethtool",
    importpath = "github.com/safchain/ethtool",
    sum = "h1:eskphjc5kRCykOJyX7HHVbJCs25/8knprttvrVvEd8o=",
    version = "v0.0.0-20201023143004-874930cb3ce0",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    sum = "h1:Kpca3qRNrduNnOQeazBd0ysaKrUJiIuISHxogkT9RPQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_shurcool_sanitized_anchor_name",
    importpath = "github.com/shurcooL/sanitized_anchor_name",
    sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_soheilhy_cmux",
    importpath = "github.com/soheilhy/cmux",
    sum = "h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    importpath = "github.com/spaolacci/murmur3",
    sum = "h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
    version = "v0.0.0-20180118202830-f09979ecbc72",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    sum = "h1:5jhuqJyZCZf2JRofRvN/nIFgIWNzPa3/Vz8mYylgbWc=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    sum = "h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    sum = "h1:KfztREH0tPxJJ+geloSLaAkaPkr4ki2Er5quFV1TDo4=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    sum = "h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    sum = "h1:xVKxvI7ouOI5I+U9s2eeiUfMaWBVoXA3AWskkrqK0VM=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_tidwall_pretty",
    importpath = "github.com/tidwall/pretty",
    sum = "h1:HsD+QiTn7sK6flMKIvNmpqz1qrpP3Ps6jOKIKMooyg4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_tinylib_msgp",
    importpath = "github.com/tinylib/msgp",
    sum = "h1:9fQd+ICuRIu/ue4vxJZu6/LzxN0HwMds2nq/0cFvxHU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_tmc_grpc_websocket_proxy",
    importpath = "github.com/tmc/grpc-websocket-proxy",
    sum = "h1:LnC5Kc/wtumK+WB441p7ynQJzVuNRJiqddSIE3IlSEQ=",
    version = "v0.0.0-20190109142713-0ad062ec5ee5",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    importpath = "github.com/ugorji/go/codec",
    sum = "h1:C5VurWRRCKjuENsbM6GYVw8W++WVW9rSxoACKIvxzz8=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_urfave_cli",
    importpath = "github.com/urfave/cli",
    sum = "h1:fDqGv3UG/4jbVl/QkFwEdddtEDjh/5Ov6X+0B/3bPaw=",
    version = "v1.20.0",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    importpath = "github.com/valyala/bytebufferpool",
    sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_fasttemplate",
    importpath = "github.com/valyala/fasttemplate",
    sum = "h1:TVEnxayobAdVkhQfrfes2IzOB6o+z4roRkPF52WA1u4=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_vektah_gqlparser",
    importpath = "github.com/vektah/gqlparser",
    sum = "h1:ZsyLGn7/7jDNI+y4SEhI4yAxRChlv15pUHMjijT+e68=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_vishvananda_netlink",
    importpath = "github.com/vishvananda/netlink",
    sum = "h1:1iyaYNBLmP6L0220aDnYQpo1QEV4t4hJ+xEEhhJH8j0=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_vishvananda_netns",
    importpath = "github.com/vishvananda/netns",
    sum = "h1:OviZH7qLw/7ZovXvuNyL3XQl8UFofeikI1NW1Gypu7k=",
    version = "v0.0.0-20191106174202-0a2b9b5464df",
)

go_repository(
    name = "com_github_willf_bitset",
    importpath = "github.com/willf/bitset",
    sum = "h1:N7Z7E9UvjW+sGsEl7k/SJrvY2reP1A07MrGuCjIOjRE=",
    version = "v1.1.11",
)

go_repository(
    name = "com_github_xiang90_probing",
    importpath = "github.com/xiang90/probing",
    sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
    version = "v0.0.0-20190116061207-43a291ad63a2",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    importpath = "github.com/xordataexchange/crypt",
    sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
    version = "v0.0.3-0.20170626215501-b2862e3d0a77",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:dPmz1Snjq0kmkz159iL7S6WzdahUTHnHB5M56WFVifs=",
    version = "v1.3.5",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:at8Tk2zUz63cLPR0JPWm5vp77pEZmzxEQBEfRKn1VV8=",
    version = "v0.81.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:PQcPefKFdaIzjQFbiyOgAqyx8q5djaE7x9Sqe712DPA=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_datastore",
    importpath = "cloud.google.com/go/datastore",
    sum = "h1:/May9ojXjRkPBNVrq+oWLqmWCkr4OU5uRY29bu0mRyQ=",
    version = "v1.1.0",
)

go_repository(
    name = "com_google_cloud_go_pubsub",
    importpath = "cloud.google.com/go/pubsub",
    sum = "h1:ukjixP1wl0LpnZ6LWtZJ0mX5tBmjp1f8Sqer8Z2OMUU=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_storage",
    importpath = "cloud.google.com/go/storage",
    sum = "h1:STgFzyU5/8miMl0//zKh2aQeTyeaUH3WN9bSUiJ09bA=",
    version = "v1.10.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
    version = "v0.0.0-20190408044501-666a987793e9",
)

go_repository(
    name = "in_gopkg_airbrake_gobrake_v2",
    importpath = "gopkg.in/airbrake/gobrake.v2",
    sum = "h1:7z2uVWwn7oVeeugY1DtlPAy5H+KYgB1KeKTnqjNatLo=",
    version = "v2.0.9",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:BLraFXnmrev5lT+xlilqcH8XK9/i0At2xKjWk4p6zsU=",
    version = "v1.0.0-20200227125254-8fa46927fb4f",
)

go_repository(
    name = "in_gopkg_cheggaaa_pb_v1",
    importpath = "gopkg.in/cheggaaa/pb.v1",
    sum = "h1:Ev7yu1/f6+d+b3pi5vPdRPc6nNtP1umSfcWiEfRqv6I=",
    version = "v1.0.25",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_gemnasium_logrus_airbrake_hook_v2",
    importpath = "gopkg.in/gemnasium/logrus-airbrake-hook.v2",
    sum = "h1:OAj3g0cR6Dx/R07QgQe8wkA9RNjB2u4i700xBkIT4e0=",
    version = "v2.1.2",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "in_gopkg_mgo_v2",
    importpath = "gopkg.in/mgo.v2",
    sum = "h1:xcEWjVhvbDy+nHP67nPDDpbYrY+ILlfndk4bRioVHaU=",
    version = "v2.0.0-20180705113604-9856a29383ce",
)

go_repository(
    name = "in_gopkg_natefinch_lumberjack_v2",
    importpath = "gopkg.in/natefinch/lumberjack.v2",
    sum = "h1:1Lc07Kr7qY4U2YPouBjpCLxpiyxIVoxqXgkXLknAOE8=",
    version = "v2.0.0",
)

go_repository(
    name = "in_gopkg_resty_v1",
    importpath = "gopkg.in/resty.v1",
    sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
    version = "v1.12.0",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    importpath = "gopkg.in/square/go-jose.v2",
    sum = "h1:orlkJ3myw8CN1nVQHBFfloD+L3egixIa4FvUP6RosSA=",
    version = "v2.2.2",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
    version = "v2.4.0",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:h8qDotaEPuJATrMmW04NCwg7v22aHH28wwpauUhK9Oo=",
    version = "v3.0.0-20210107192922-496545a6307b",
)

go_repository(
    name = "io_etcd_go_bbolt",
    importpath = "go.etcd.io/bbolt",
    sum = "h1:XAzx9gjCb0Rxj7EoqcClPD1d5ZBxZJk0jbuoPHenBt0=",
    version = "v1.3.5",
)

go_repository(
    name = "io_etcd_go_etcd",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "go.etcd.io/etcd",
    sum = "h1:1JFLBqwIgdyHN1ZtgjTBwO+blA6gVOmZurpiMEsETKo=",
    version = "v0.5.0-alpha.5.0.20200910180754-dd1b699fc489",
)

go_repository(
    name = "io_k8s_apiextensions_apiserver",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/apiextensions-apiserver",
    sum = "h1:rfrMWQ87lhd8EzQWRnbQ4gXrniL/yTRBgYH1x1+BLlo=",
    version = "v0.20.2",
)

go_repository(
    name = "io_k8s_apiserver",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/apiserver",
    sum = "h1:lGno2t3gcZnLtzsKH4oG0xA9/4GTiBzMO1DGp+K+Bak=",
    version = "v0.20.2",
)

go_repository(
    name = "io_k8s_code_generator",
    importpath = "k8s.io/code-generator",
    sum = "h1:SQaysped4EtUDk3u1zphnUJiOAwFdhHx9xS3WKAE0x8=",
    version = "v0.20.2",
)

go_repository(
    name = "io_k8s_component_base",
    importpath = "k8s.io/component-base",
    sum = "h1:LMmu5I0pLtwjpp5009KLuMGFqSc2S2isGw8t1hpYKLE=",
    version = "v0.20.2",
)

go_repository(
    name = "io_k8s_gengo",
    importpath = "k8s.io/gengo",
    sum = "h1:GohjlNKauSai7gN4wsJkeZ3WAJx4Sh+oT/b5IYn5suA=",
    version = "v0.0.0-20210813121822-485abfe95c7c",
)

go_repository(
    name = "io_k8s_kube_aggregator",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/kube-aggregator",
    sum = "h1:iDJILLwIKjojE0bjZHKMGp8Ry5U1ugsJzrb/A9lD+00=",
    version = "v0.19.2",
)

go_repository(
    name = "io_k8s_kube_openapi",
    importpath = "k8s.io/kube-openapi",
    sum = "h1:E3J9oCLlaobFUqsjG9DfKbP2BmgwBL2p7pn0A3dG9W4=",
    version = "v0.0.0-20211115234752-e816edb12b65",
)

go_repository(
    name = "io_k8s_kubectl",
    importpath = "k8s.io/kubectl",
    sum = "h1:mXExF6N4eQUYmlfXJmfWIheCBLF6/n4VnwQKbQki5iE=",
    version = "v0.20.2",
)

go_repository(
    name = "io_k8s_sigs_apiserver_builder_alpha",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "sigs.k8s.io/apiserver-builder-alpha",
    patches = [
        "//:patches/api-server-alpha.patch",
    ],
    sum = "h1:i+XvReFGEJAUgHmATTPouxfDrSdR+YOcbT0SIEQ+5gE=",
    version = "v1.18.1-0.20201012071248-ca5d7287e44a",
)

go_repository(
    name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
    importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
    sum = "h1:TihvEz9MPj2u0KWds6E2OBUXfwaL4qRJ33c7HGiJpqk=",
    version = "v0.0.14",
)

go_repository(
    name = "io_k8s_sigs_controller_runtime",
    importpath = "sigs.k8s.io/controller-runtime",
    sum = "h1:O0K2CJ2JavK8/Tf4LfcpAwRxOFBhv8DjyrbmE6Qw59s=",
    version = "v0.8.1",
)

go_repository(
    name = "io_k8s_sigs_controller_tools",
    importpath = "sigs.k8s.io/controller-tools",
    sum = "h1:la1h46EzElvWefWLqfsXrnsO3lZjpkI0asTpX6h8PLA=",
    version = "v0.2.4",
)

go_repository(
    name = "io_k8s_sigs_kubebuilder",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "sigs.k8s.io/kubebuilder",
    sum = "h1:zviasl0ksxaXMRKBlDMtzODlq1xwVJpNkkYCITbcAvg=",
    version = "v1.0.9-0.20200925141511-a2f239880b04",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff",
    importpath = "sigs.k8s.io/structured-merge-diff",
    sum = "h1:4Z09Hglb792X0kfOBBJUPFEyvVfQWrYT/l8h5EKA6JQ=",
    version = "v0.0.0-20190525122527-15d366b2352e",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff_v3",
    importpath = "sigs.k8s.io/structured-merge-diff/v3",
    sum = "h1:dOmIZBMfhcHS09XZkMyUgkq5trg3/jRyJYFZUiaOp8E=",
    version = "v3.0.0",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff_v4",
    importpath = "sigs.k8s.io/structured-merge-diff/v4",
    sum = "h1:bKCqE9GvQ5tiVHn5rfn1r+yao3aLQEaLzkkmAkf+A6Y=",
    version = "v4.2.1",
)

go_repository(
    name = "io_k8s_sigs_testing_frameworks",
    importpath = "sigs.k8s.io/testing_frameworks",
    sum = "h1:cP2l8fkA3O9vekpy5Ks8mmA0NW/F7yBdXf8brkWhVrs=",
    version = "v0.1.1",
)

go_repository(
    name = "io_k8s_sigs_yaml",
    importpath = "sigs.k8s.io/yaml",
    sum = "h1:kr/MCeFWJWTwyaHoR9c8EjH9OumOmoF9YGiZd7lFm/Q=",
    version = "v1.2.0",
)

go_repository(
    name = "io_k8s_utils",
    importpath = "k8s.io/utils",
    sum = "h1:ck1fRPWPJWsMd8ZRFsWc6mh/zHp5fZ/shhbrgPUxDAE=",
    version = "v0.0.0-20211116205334-6203023598ed",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    sum = "h1:gqCw0LfLxScz8irSi8exQc7fyQ0fKQU/qnC/X8+V/1M=",
    version = "v0.23.0",
)

go_repository(
    name = "io_rsc_binaryregexp",
    importpath = "rsc.io/binaryregexp",
    sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
    version = "v0.2.0",
)

go_repository(
    name = "io_rsc_quote_v3",
    importpath = "rsc.io/quote/v3",
    sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
    version = "v3.1.0",
)

go_repository(
    name = "io_rsc_sampler",
    importpath = "rsc.io/sampler",
    sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
    version = "v1.3.0",
)

go_repository(
    name = "net_howett_plist",
    importpath = "howett.net/plist",
    sum = "h1:jhnBjNi9UFpfpl8YZhA9CrOqpnJdvzuiHsl/dnxl11M=",
    version = "v0.0.0-20181124034731-591f970eefbb",
)

go_repository(
    name = "net_launchpad_xmlpath",
    importpath = "launchpad.net/xmlpath",
    sum = "h1:B8nNZBUrx8YufDCAJjvO/lVs4GxXMQHyrjwJdJzXMFg=",
    version = "v0.0.0-20130614043138-000000000004",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    sum = "h1:4sAyIHT6ZohtAQDoxws+ez7bROYmUlOVvsUscYCDTqA=",
    version = "v0.43.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
    version = "v1.6.7",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:E7wSQBXkH3T3diucK+9Z1kjn4+/9tNG7lZLr75oOhh8=",
    version = "v0.0.0-20210402141018-6c239bbf2bb1",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:SnqbnDw1V7RiZcXPx5MEeqPv2s79L9i7BJUlG/+RurQ=",
    version = "v1.27.1",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:HWj/xjIHfjYU5nVXpTM0s39J9CbLn7Cc5a7IC5rwsMQ=",
    version = "v0.0.0-20210817164053-32db794688a5",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:QE6XYQK6naiK1EPAe1g/ILLxN5RBoH5xkJk3CqlMI/Y=",
    version = "v0.0.0-20200224162631-6cc2880d07d6",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
    version = "v0.0.0-20190802002840-cff245a6509b",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:2M3HP5CCK1Si9FQhwnzYhXdG6DXeebvUHFpre8QvbyI=",
    version = "v0.0.0-20201208152925-83fdc39ff7b5",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
    version = "v0.0.0-20190719004257-d2bd2a29d028",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:Gz96sIWK3OalVv/I/qNygP42zyoKp3xptRVCWRFEBvo=",
    version = "v0.4.2",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:Qmd2pbz05z7z6lm0DrgQVVPuBm92jqujBKMHMOlOQEw=",
    version = "v0.0.0-20210819190943-2bc19b11175f",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:5KslGYwFpkhGh+Q16bwMP3cOontH8FOep7tGV86Y7SQ=",
    version = "v0.0.0-20210220032951-036812b2e83c",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:id054HUawV2/6IGm2IV8KZQjqtwAOo2CYlOToYqa0d0=",
    version = "v0.0.0-20211019181941-9d821ace8654",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=",
    version = "v0.3.7",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:7zkz7BUtwNFFqcowJ+RIgu2MaV/MapERkDIy+mwPyjs=",
    version = "v0.0.0-20210723032227-1f47c861a9ac",
)


go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:ouewzE6p+/VEB31YYnTbEJdi8pFqKp4P4n85vwo3DHA=",
    version = "v0.1.5",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_repository(
    name = "org_mongodb_go_mongo_driver",
    importpath = "go.mongodb.org/mongo-driver",
    sum = "h1:jxcFYjlkl8xaERsgLo+RNquI0epW6zuy/ZRQs6jnrFA=",
    version = "v1.1.2",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:Ezj3JGmsOnG1MoRWQkPBsKLe9DwWD9QeXzTRzzldNVk=",
    version = "v1.6.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:KCa4XfM8CWFCpxXRGok+Q0SS/0XBhMDbHHGABQLvD2A=",
    version = "v1.5.0",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:ue41HOKd1vGURxrmeKIgELGb3jPW9DMUDGtsinblHwI=",
    version = "v1.19.1",
)

go_repository(
    name = "com_github_antchfx_xmlquery",
    build_file_generation = "on",
    importpath = "github.com/antchfx/xmlquery",
    sum = "h1:HYmadPG0uz8CySdL68rB4DCLKXz2PurCjS3mnkVF4CQ=",
    version = "v1.3.3",
)

go_repository(
    name = "com_github_antchfx_xpath",
    build_file_generation = "on",
    importpath = "github.com/antchfx/xpath",
    sum = "h1:WOFtK8TVAjLm3lbgqeP0arlHpvCEeTANeWZ/csPpJkQ=",
    version = "v1.1.11",
)

go_repository(
    name = "tools_gotest",
    importpath = "gotest.tools",
    sum = "h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
    version = "v2.2.0+incompatible",
)

go_repository(
    name = "tools_gotest_v3",
    importpath = "gotest.tools/v3",
    sum = "h1:kG1BFyqVHuQoVQiR1bWGnfz/fmHvvuiSPIV7rvl360E=",
    version = "v3.0.2",
)

go_repository(
    name = "xyz_gomodules_jsonpatch_v2",
    importpath = "gomodules.xyz/jsonpatch/v2",
    sum = "h1:Phva6wqu+xR//Njw6iorylFFgn/z547tw5Ne3HZPQ+k=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_bazel_gomock",
    commit = "fde78c91cf1783cc1e33ba278922ba67a6ee2a84",
    importpath = "github.com/jmhodges/bazel_gomock",
)

# inet.af/netaddr v0.0.0-20210311133851-b21affee3d06 h1:+12VMoF1ykxcVvOICCLi5jSPdbzPjLNzOHrSluK7tVA=
go_repository(
    name = "af_inet_netaddr",
    importpath = "inet.af/netaddr",
    sum = "h1:+12VMoF1ykxcVvOICCLi5jSPdbzPjLNzOHrSluK7tVA=",
    version = "v0.0.0-20210311133851-b21affee3d06",
)

# go4.org/intern v0.0.0-20210108033219-3eb7198706b2 h1:VFTf+jjIgsldaz/Mr00VaCSswHJrI2hIjQygE/W4IMg=
go_repository(
    name = "org_go4_intern",
    importpath = "go4.org/intern",
    sum = "h1:VFTf+jjIgsldaz/Mr00VaCSswHJrI2hIjQygE/W4IMg=",
    version = "v0.0.0-20210108033219-3eb7198706b2",
)

# go4.org/unsafe/assume-no-moving-gc v0.0.0-20201222180813-1025295fd063 h1:1tk03FUNpulq2cuWpXZWj649rwJpk0d20rxWiopKRmc=
go_repository(
    name = "org_go4_unsafe_assume-no-moving-gc",
    importpath = "go4.org/unsafe/assume-no-moving-gc",
    sum = "h1:1tk03FUNpulq2cuWpXZWj649rwJpk0d20rxWiopKRmc=",
    version = "v0.0.0-20201222180813-1025295fd063",
)
