
# API Docs

This document enumerates the Custom Resource Definitions used by the M3DB Operator. It is auto-generated from code comments.

## Table of Contents
* [ClusterCondition](#clustercondition)
* [ClusterSpec](#clusterspec)
* [IsolationGroup](#isolationgroup)
* [M3DBCluster](#m3dbcluster)
* [M3DBClusterList](#m3dbclusterlist)
* [M3DBStatus](#m3dbstatus)
* [IndexOptions](#indexoptions)
* [Namespace](#namespace)
* [NamespaceOptions](#namespaceoptions)
* [RetentionOptions](#retentionoptions)
* [PodIdentity](#podidentity)
* [PodIdentityConfig](#podidentityconfig)

## ClusterCondition

ClusterCondition represents various conditions the cluster can be in.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| type | Type of cluster condition. | ClusterConditionType | false |
| status | Status of the condition (True, False, Unknown). | corev1.ConditionStatus | false |
| lastUpdateTime | Last time this condition was updated. | string | false |
| lastTransitionTime | Last time this condition transitioned from one status to another. | string | false |
| reason | Reason this condition last changed. | string | false |
| message | Human-friendly message about this condition. | string | false |

[Back to TOC](#table-of-contents)

## ClusterSpec

ClusterSpec defines the desired state for a M3 cluster to be converge to.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| image | Image specifies which docker image to use with the cluster | string | false |
| replicationFactor | ReplicationFactor defines how many replicas | int32 | false |
| numberOfShards | NumberOfShards defines how many shards in total | int32 | false |
| isolationGroups | IsolationGroups specifies a map of key-value pairs. Defines which isolation groups to deploy persistent volumes for data nodes | [][IsolationGroup](#isolationgroup) | false |
| namespaces | Namespaces specifies the namespaces this cluster will hold. | [][Namespace](#namespace) | false |
| configMapName | ConfigMapName specifies the ConfigMap to use for this cluster. If unset a sane default will be used. | *string | false |
| podIdentityConfig | PodIdentityConfig sets the configuration for pod identity. If unset only pod name and UID will be used. | *PodIdentityConfig | false |
| containerResources | Resources defines memory / cpu constraints for each container in the cluster. | [corev1.ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.10/#resourcerequirements-v1-core) | false |
| dataDirVolumeClaimTemplate | DataDirVolumeClaimTemplate is the volume claim template for an M3DB instance's data. It claims PersistentVolumes for cluster storage, volumes are dynamically provisioned by when the StorageClass is defined. | *[corev1.PersistentVolumeClaim](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.10/#persistentvolumeclaim-v1-core) | false |
| labels | Labels sets the base labels that will be applied to resources created by the cluster. // TODO(schallert): design doc on labeling scheme. | map[string]string | false |

[Back to TOC](#table-of-contents)

## IsolationGroup

IsolationGroup defines the name of zone as well attributes for the zone configuration

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| name | Name | string | false |
| numInstances | NumInstances defines the number of instances | int32 | false |

[Back to TOC](#table-of-contents)

## M3DBCluster

M3DBCluster defines the cluster

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.10/#objectmeta-v1-meta) | false |
| type |  | string | true |
| spec |  | [ClusterSpec](#clusterspec) | true |
| status |  | [M3DBStatus](#m3dbstatus) | false |

[Back to TOC](#table-of-contents)

## M3DBClusterList

M3DBClusterList represents a list of M3DB Clusters

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.10/#listmeta-v1-meta) | false |
| items |  | [][M3DBCluster](#m3dbcluster) | true |

[Back to TOC](#table-of-contents)

## M3DBStatus

M3DBStatus contains the current state the M3DB cluster along with a human readable message

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| state | State is a enum of green, yellow, and red denoting the health of the cluster | M3DBState | false |
| conditions | Various conditions about the cluster. | [][ClusterCondition](#clustercondition) | false |
| message | Message is a human readable message indicating why the cluster is in it's current state | string | false |
| observedGeneration | ObservedGeneration is the last generation of the cluster the controller observed. Kubernetes will automatically increment metadata.Generation every time the cluster spec is changed. | int64 | false |

[Back to TOC](#table-of-contents)

## IndexOptions

IndexOptions defines parameters for indexing.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| enabled | Enabled controls whether metric indexing is enabled. | bool | false |
| blockSize | BlockSize controls the index block size. | time.Duration | false |

[Back to TOC](#table-of-contents)

## Namespace

Namespace defines an M3DB namespace or points to a preset M3DB namespace.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| name | Name is the namespace name. | string | false |
| preset | Preset indicates preset namespace options. | string | false |
| options | Options points to optional custom namespace configuration. | *[NamespaceOptions](#namespaceoptions) | false |

[Back to TOC](#table-of-contents)

## NamespaceOptions

NamespaceOptions defines parameters for an M3DB namespace. See https://m3db.github.io/m3/operational_guide/namespace_configuration/ for more details.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| bootstrapEnabled | BootstrapEnabled control if bootstrapping is enabled. | bool | false |
| flushEnabled | FlushEnabled controls whether flushing is enabled. | bool | false |
| writesToCommitLog | WritesToCommitLog controls whether commit log writes are enabled. | bool | false |
| cleanupEnabled | CleanupEnabled controls whether cleanups are enabled. | bool | false |
| repairEnabled | RepairEnabled controls whether repairs are enabled. | bool | false |
| snapshotEnabled | SnapshotEnabled controls whether snapshotting is enabled. | bool | false |
| retentionOptions | RetentionOptions sets the retention parameters. | [RetentionOptions](#retentionoptions) | false |
| indexOptions | IndexOptions sets the indexing parameters. | [IndexOptions](#indexoptions) | false |

[Back to TOC](#table-of-contents)

## RetentionOptions

RetentionOptions defines parameters for data retention.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| retentionPeriod | RetentionPeriod controls how long data for the namespace is retained. | time.Duration | false |
| blockSize | BlockSize controls the block size for the namespace. | time.Duration | false |
| bufferFuture | BufferFuture controls how far in the future metrics can be written. | time.Duration | false |
| bufferPast | BufferPast controls how far in the past metrics can be written. | time.Duration | false |
| blockDataExpiry | BlockDataExpiry controls the block expiry. | bool | false |
| blockDataExpiryAfterNotAccessPeriod | BlockDataExpiry controls the not after access period for expiration. | time.Duration | false |

[Back to TOC](#table-of-contents)

## PodIdentity

PodIdentity contains all the fields that may be used to identify a pod's identity in the M3DB placement. Any non-empty fields will be used to identity uniqueness of a pod for the purpose of M3DB replace operations.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| name |  | string | false |
| uid |  | string | false |
| nodeName |  | string | false |
| nodeExternalID |  | string | false |
| nodeProviderID |  | string | false |

[Back to TOC](#table-of-contents)

## PodIdentityConfig

PodIdentityConfig contains cluster-level configuration for deriving pod identity.

| Field | Description | Scheme | Required |
| ----- | ----------- | ------ | -------- |
| sources | Sources enumerates the sources from which to derive pod identity. Note that a pod's name will always be used. If empty, defaults to pod name and UID. | []PodIdentitySource | true |

[Back to TOC](#table-of-contents)
