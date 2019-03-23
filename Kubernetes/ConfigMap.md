很多生产环境中的应用程序配置较为复杂，可能需要多个config文件、命令行参数和环境变量的组合。使用容器部署时，把配置应该从应用程序镜像中解耦出来，以保证镜像的可移植性。尽管Secret允许类似于验证信息和秘钥等信息从应用中解耦出来，但在K8S1.2前并没有为了普通的或者非secret配置而存在的对象。在K8S1.2后引入ConfigMap来处理这种类型的配置数据。

ConfigMap是存储通用的配置变量的，类似于配置文件，使用户可以将分布式系统中用于不同模块的环境变量统一到一个对象中管理；而它与配置文件的区别在于它是存在集群的“环境”中的，并且支持K8S集群中所有通用的操作调用方式。

从数据角度来看，ConfigMap的类型只是键值组，用于存储被Pod或者其他资源对象（如RC）访问的信息。这与secret的设计理念有异曲同工之妙，主要区别在于ConfigMap通常不用于存储敏感信息，而只存储简单的文本信息。

ConfigMap可以保存环境变量的属性，也可以保存配置文件。

创建pod时，对configmap进行绑定，pod内的应用可以直接引用ConfigMap的配置。相当于configmap为应用/运行环境封装配置。

pod使用ConfigMap，通常用于：设置环境变量的值、设置命令行参数、创建配置文件



创建configmap
可使用kubectl create configmap从文件，目录或key-value字符串创建，也可通过kubectl create -f file创建

命令行
kubectl create configmap lykops-config --from-file=db_config_file=database.conf --from-file=ver.conf --from-literal=username=test --from-literal=hostname=localhost

–from-file表示来自文件，直接把文件内容写入configmap中，可以为目录也可以为文件，如果是文件的话，可以使用db_config_file=database.conf来修改key值 
–from-literal表示使用键值对配置


configmap用于保存配置数据的键值对，可以用来保存单个属性，也可保存配置文件。configmap与secret很类似，但它可以更方便的处理不包含敏感信息的字符串
