// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/gatekeeper-rendered/admissionregistration.k8s.io_v1_mutatingwebhookconfiguration_gatekeeper-mutating-webhook-configuration.yaml
// config/gatekeeper-rendered/admissionregistration.k8s.io_v1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_assign.mutations.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_assignmetadata.mutations.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_configs.config.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_expansiontemplate.expansion.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_modifyset.mutations.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_mutatorpodstatuses.status.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_providers.externaldata.gatekeeper.sh.yaml
// config/gatekeeper-rendered/apps_v1_deployment_gatekeeper-audit.yaml
// config/gatekeeper-rendered/apps_v1_deployment_gatekeeper-controller-manager.yaml
// config/gatekeeper-rendered/policy_v1_poddisruptionbudget_gatekeeper-controller-manager.yaml
// config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml
// config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml
// config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml
// config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml
// config/gatekeeper-rendered/v1_namespace_gatekeeper-system.yaml
// config/gatekeeper-rendered/v1_resourcequota_gatekeeper-critical-pods.yaml
// config/gatekeeper-rendered/v1_secret_gatekeeper-webhook-server-cert.yaml
// config/gatekeeper-rendered/v1_service_gatekeeper-webhook-service.yaml
// config/gatekeeper-rendered/v1_serviceaccount_gatekeeper-admin.yaml
package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configGatekeeperRenderedAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml = []byte(`apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-mutating-webhook-configuration
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/mutate
  failurePolicy: Ignore
  matchPolicy: Exact
  name: mutation.gatekeeper.sh
  namespaceSelector:
    matchExpressions:
    - key: admission.gatekeeper.sh/ignore
      operator: DoesNotExist
    - key: kubernetes.io/metadata.name
      operator: NotIn
      values:
      - gatekeeper-system
  rules:
  - apiGroups:
    - '*'
    apiVersions:
    - '*'
    operations:
    - CREATE
    - UPDATE
    resources:
    - '*'
  sideEffects: None
  timeoutSeconds: 1
`)

func configGatekeeperRenderedAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml, nil
}

func configGatekeeperRenderedAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/admissionregistration.k8s.io_v1_mutatingwebhookconfiguration_gatekeeper-mutating-webhook-configuration.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml = []byte(`apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-validating-webhook-configuration
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/admit
  failurePolicy: Ignore
  matchPolicy: Exact
  name: validation.gatekeeper.sh
  namespaceSelector:
    matchExpressions:
    - key: admission.gatekeeper.sh/ignore
      operator: DoesNotExist
    - key: kubernetes.io/metadata.name
      operator: NotIn
      values:
      - gatekeeper-system
  rules:
  - apiGroups:
    - '*'
    apiVersions:
    - '*'
    operations:
    - CREATE
    - UPDATE
    resources:
    - '*'
    - pods/ephemeralcontainers
    - pods/exec
    - pods/log
    - pods/eviction
    - pods/portforward
    - pods/proxy
    - pods/attach
    - pods/binding
    - deployments/scale
    - replicasets/scale
    - statefulsets/scale
    - replicationcontrollers/scale
    - services/proxy
    - nodes/proxy
    - services/status
  sideEffects: None
  timeoutSeconds: 3
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/admitlabel
  failurePolicy: Fail
  matchPolicy: Exact
  name: check-ignore-label.gatekeeper.sh
  namespaceSelector:
    matchExpressions:
    - key: kubernetes.io/metadata.name
      operator: NotIn
      values:
      - gatekeeper-system
  rules:
  - apiGroups:
    - ""
    apiVersions:
    - '*'
    operations:
    - CREATE
    - UPDATE
    resources:
    - namespaces
  sideEffects: None
  timeoutSeconds: 3
`)

func configGatekeeperRenderedAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml, nil
}

func configGatekeeperRenderedAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/admissionregistration.k8s.io_v1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: assign.mutations.gatekeeper.sh
spec:
  group: mutations.gatekeeper.sh
  names:
    kind: Assign
    listKind: AssignList
    plural: assign
    singular: assign
  preserveUnknownFields: false
  scope: Cluster
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: Assign is the Schema for the assign API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            properties:
              name:
                maxLength: 63
                type: string
            type: object
          spec:
            description: AssignSpec defines the desired state of Assign.
            properties:
              applyTo:
                description: ApplyTo lists the specific groups, versions and kinds
                  a mutation will be applied to. This is necessary because every mutation
                  implies part of an object schema and object schemas are associated
                  with specific GVKs.
                items:
                  description: ApplyTo determines what GVKs items the mutation should
                    apply to. Globs are not allowed.
                  properties:
                    groups:
                      items:
                        type: string
                      type: array
                    kinds:
                      items:
                        type: string
                      type: array
                    versions:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              location:
                description: 'Location describes the path to be mutated, for example:
                  ` + "`" + `spec.containers[name: main]` + "`" + `.'
                type: string
              match:
                description: Match allows the user to limit which resources get mutated.
                  Individual match criteria are AND-ed together. An undefined match
                  criteria matches everything.
                properties:
                  excludedNamespaces:
                    description: 'ExcludedNamespaces is a list of namespace names.
                      If defined, a constraint only applies to resources not in a
                      listed namespace. ExcludedNamespaces also supports a prefix
                      or suffix based glob.  For example, ` + "`" + `excludedNamespaces: [kube-*]` + "`" + `
                      matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and ` + "`" + `excludedNamespaces:
                      [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: 'LabelSelector is the combination of two optional
                      fields: ` + "`" + `matchLabels` + "`" + ` and ` + "`" + `matchExpressions` + "`" + `.  These two fields
                      provide different methods of selecting or excluding k8s objects
                      based on the label keys and values included in object metadata.  All
                      selection expressions from both sections are ANDed to determine
                      if an object meets the cumulative requirements of the selector.'
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  name:
                    description: 'Name is the name of an object.  If defined, it will
                      match against objects with the specified name.  Name also supports
                      a prefix or suffix glob.  For example, ` + "`" + `name: pod-*` + "`" + ` would match
                      both ` + "`" + `pod-a` + "`" + ` and ` + "`" + `pod-b` + "`" + `, and ` + "`" + `name: *-pod` + "`" + ` would match both
                      ` + "`" + `a-pod` + "`" + ` and ` + "`" + `b-pod` + "`" + `.'
                    pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                    type: string
                  namespaceSelector:
                    description: NamespaceSelector is a label selector against an
                      object's containing namespace or the object itself, if the object
                      is a namespace.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    description: 'Namespaces is a list of namespace names. If defined,
                      a constraint only applies to resources in a listed namespace.  Namespaces
                      also supports a prefix or suffix based glob.  For example, ` + "`" + `namespaces:
                      [kube-*]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and
                      ` + "`" + `namespaces: [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  scope:
                    description: Scope determines if cluster-scoped and/or namespaced-scoped
                      resources are matched.  Accepts ` + "`" + `*` + "`" + `, ` + "`" + `Cluster` + "`" + `, or ` + "`" + `Namespaced` + "`" + `.
                      (defaults to ` + "`" + `*` + "`" + `)
                    type: string
                  source:
                    description: Source determines whether generated or original resources
                      are matched. Accepts ` + "`" + `Generated` + "`" + `|` + "`" + `Original` + "`" + `|` + "`" + `All` + "`" + ` (defaults
                      to ` + "`" + `All` + "`" + `). A value of ` + "`" + `Generated` + "`" + ` will only match generated
                      resources, while ` + "`" + `Original` + "`" + ` will only match regular resources.
                    enum:
                    - All
                    - Generated
                    - Original
                    type: string
                type: object
              parameters:
                description: Parameters define the behavior of the mutator.
                properties:
                  assign:
                    description: Assign.value holds the value to be assigned
                    properties:
                      externalData:
                        description: ExternalData describes the external data provider
                          to be used for mutation.
                        properties:
                          dataSource:
                            default: ValueAtLocation
                            description: DataSource specifies where to extract the
                              data that will be sent to the external data provider
                              as parameters.
                            enum:
                            - ValueAtLocation
                            - Username
                            type: string
                          default:
                            description: Default specifies the default value to use
                              when the external data provider returns an error and
                              the failure policy is set to "UseDefault".
                            type: string
                          failurePolicy:
                            default: Fail
                            description: FailurePolicy specifies the policy to apply
                              when the external data provider returns an error.
                            enum:
                            - UseDefault
                            - Ignore
                            - Fail
                            type: string
                          provider:
                            description: Provider is the name of the external data
                              provider.
                            type: string
                        type: object
                      fromMetadata:
                        description: FromMetadata assigns a value from the specified
                          metadata field.
                        properties:
                          field:
                            description: Field specifies which metadata field provides
                              the assigned value. Valid fields are ` + "`" + `namespace` + "`" + ` and
                              ` + "`" + `name` + "`" + `.
                            type: string
                        type: object
                      value:
                        description: Value is a constant value that will be assigned
                          to ` + "`" + `location` + "`" + `
                        x-kubernetes-preserve-unknown-fields: true
                    type: object
                  pathTests:
                    items:
                      description: "PathTest allows the user to customize how the
                        mutation works if parent paths are missing. It traverses the
                        list in order. All sub paths are tested against the provided
                        condition, if the test fails, the mutation is not applied.
                        All ` + "`" + `subPath` + "`" + ` entries must be a prefix of ` + "`" + `location` + "`" + `. Any
                        glob characters will take on the same value as was used to
                        expand the matching glob in ` + "`" + `location` + "`" + `. \n Available Tests:
                        * MustExist    - the path must exist or do not mutate * MustNotExist
                        - the path must not exist or do not mutate."
                      properties:
                        condition:
                          description: Condition describes whether the path either
                            MustExist or MustNotExist in the original object
                          enum:
                          - MustExist
                          - MustNotExist
                          type: string
                        subPath:
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: AssignStatus defines the observed state of Assign.
            properties:
              byPod:
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus.
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system.
                        properties:
                          message:
                            type: string
                          type:
                            description: Type indicates a specific class of error
                              for use by controller code. If not present, the error
                              should be treated as not matching any known type.
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: Assign is the Schema for the assign API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AssignSpec defines the desired state of Assign.
            properties:
              applyTo:
                description: ApplyTo lists the specific groups, versions and kinds
                  a mutation will be applied to. This is necessary because every mutation
                  implies part of an object schema and object schemas are associated
                  with specific GVKs.
                items:
                  description: ApplyTo determines what GVKs items the mutation should
                    apply to. Globs are not allowed.
                  properties:
                    groups:
                      items:
                        type: string
                      type: array
                    kinds:
                      items:
                        type: string
                      type: array
                    versions:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              location:
                description: 'Location describes the path to be mutated, for example:
                  ` + "`" + `spec.containers[name: main]` + "`" + `.'
                type: string
              match:
                description: Match allows the user to limit which resources get mutated.
                  Individual match criteria are AND-ed together. An undefined match
                  criteria matches everything.
                properties:
                  excludedNamespaces:
                    description: 'ExcludedNamespaces is a list of namespace names.
                      If defined, a constraint only applies to resources not in a
                      listed namespace. ExcludedNamespaces also supports a prefix
                      or suffix based glob.  For example, ` + "`" + `excludedNamespaces: [kube-*]` + "`" + `
                      matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and ` + "`" + `excludedNamespaces:
                      [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: 'LabelSelector is the combination of two optional
                      fields: ` + "`" + `matchLabels` + "`" + ` and ` + "`" + `matchExpressions` + "`" + `.  These two fields
                      provide different methods of selecting or excluding k8s objects
                      based on the label keys and values included in object metadata.  All
                      selection expressions from both sections are ANDed to determine
                      if an object meets the cumulative requirements of the selector.'
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  name:
                    description: 'Name is the name of an object.  If defined, it will
                      match against objects with the specified name.  Name also supports
                      a prefix or suffix glob.  For example, ` + "`" + `name: pod-*` + "`" + ` would match
                      both ` + "`" + `pod-a` + "`" + ` and ` + "`" + `pod-b` + "`" + `, and ` + "`" + `name: *-pod` + "`" + ` would match both
                      ` + "`" + `a-pod` + "`" + ` and ` + "`" + `b-pod` + "`" + `.'
                    pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                    type: string
                  namespaceSelector:
                    description: NamespaceSelector is a label selector against an
                      object's containing namespace or the object itself, if the object
                      is a namespace.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    description: 'Namespaces is a list of namespace names. If defined,
                      a constraint only applies to resources in a listed namespace.  Namespaces
                      also supports a prefix or suffix based glob.  For example, ` + "`" + `namespaces:
                      [kube-*]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and
                      ` + "`" + `namespaces: [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  scope:
                    description: Scope determines if cluster-scoped and/or namespaced-scoped
                      resources are matched.  Accepts ` + "`" + `*` + "`" + `, ` + "`" + `Cluster` + "`" + `, or ` + "`" + `Namespaced` + "`" + `.
                      (defaults to ` + "`" + `*` + "`" + `)
                    type: string
                  source:
                    description: Source determines whether generated or original resources
                      are matched. Accepts ` + "`" + `Generated` + "`" + `|` + "`" + `Original` + "`" + `|` + "`" + `All` + "`" + ` (defaults
                      to ` + "`" + `All` + "`" + `). A value of ` + "`" + `Generated` + "`" + ` will only match generated
                      resources, while ` + "`" + `Original` + "`" + ` will only match regular resources.
                    enum:
                    - All
                    - Generated
                    - Original
                    type: string
                type: object
              parameters:
                description: Parameters define the behavior of the mutator.
                properties:
                  assign:
                    description: Assign.value holds the value to be assigned
                    properties:
                      externalData:
                        description: ExternalData describes the external data provider
                          to be used for mutation.
                        properties:
                          dataSource:
                            default: ValueAtLocation
                            description: DataSource specifies where to extract the
                              data that will be sent to the external data provider
                              as parameters.
                            enum:
                            - ValueAtLocation
                            - Username
                            type: string
                          default:
                            description: Default specifies the default value to use
                              when the external data provider returns an error and
                              the failure policy is set to "UseDefault".
                            type: string
                          failurePolicy:
                            default: Fail
                            description: FailurePolicy specifies the policy to apply
                              when the external data provider returns an error.
                            enum:
                            - UseDefault
                            - Ignore
                            - Fail
                            type: string
                          provider:
                            description: Provider is the name of the external data
                              provider.
                            type: string
                        type: object
                      fromMetadata:
                        description: FromMetadata assigns a value from the specified
                          metadata field.
                        properties:
                          field:
                            description: Field specifies which metadata field provides
                              the assigned value. Valid fields are ` + "`" + `namespace` + "`" + ` and
                              ` + "`" + `name` + "`" + `.
                            type: string
                        type: object
                      value:
                        description: Value is a constant value that will be assigned
                          to ` + "`" + `location` + "`" + `
                        x-kubernetes-preserve-unknown-fields: true
                    type: object
                  pathTests:
                    items:
                      description: "PathTest allows the user to customize how the
                        mutation works if parent paths are missing. It traverses the
                        list in order. All sub paths are tested against the provided
                        condition, if the test fails, the mutation is not applied.
                        All ` + "`" + `subPath` + "`" + ` entries must be a prefix of ` + "`" + `location` + "`" + `. Any
                        glob characters will take on the same value as was used to
                        expand the matching glob in ` + "`" + `location` + "`" + `. \n Available Tests:
                        * MustExist    - the path must exist or do not mutate * MustNotExist
                        - the path must not exist or do not mutate."
                      properties:
                        condition:
                          description: Condition describes whether the path either
                            MustExist or MustNotExist in the original object
                          enum:
                          - MustExist
                          - MustNotExist
                          type: string
                        subPath:
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: AssignStatus defines the observed state of Assign.
            properties:
              byPod:
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus.
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system.
                        properties:
                          message:
                            type: string
                          type:
                            description: Type indicates a specific class of error
                              for use by controller code. If not present, the error
                              should be treated as not matching any known type.
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: Assign is the Schema for the assign API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AssignSpec defines the desired state of Assign.
            properties:
              applyTo:
                description: ApplyTo lists the specific groups, versions and kinds
                  a mutation will be applied to. This is necessary because every mutation
                  implies part of an object schema and object schemas are associated
                  with specific GVKs.
                items:
                  description: ApplyTo determines what GVKs items the mutation should
                    apply to. Globs are not allowed.
                  properties:
                    groups:
                      items:
                        type: string
                      type: array
                    kinds:
                      items:
                        type: string
                      type: array
                    versions:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              location:
                description: 'Location describes the path to be mutated, for example:
                  ` + "`" + `spec.containers[name: main]` + "`" + `.'
                type: string
              match:
                description: Match allows the user to limit which resources get mutated.
                  Individual match criteria are AND-ed together. An undefined match
                  criteria matches everything.
                properties:
                  excludedNamespaces:
                    description: 'ExcludedNamespaces is a list of namespace names.
                      If defined, a constraint only applies to resources not in a
                      listed namespace. ExcludedNamespaces also supports a prefix
                      or suffix based glob.  For example, ` + "`" + `excludedNamespaces: [kube-*]` + "`" + `
                      matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and ` + "`" + `excludedNamespaces:
                      [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: 'LabelSelector is the combination of two optional
                      fields: ` + "`" + `matchLabels` + "`" + ` and ` + "`" + `matchExpressions` + "`" + `.  These two fields
                      provide different methods of selecting or excluding k8s objects
                      based on the label keys and values included in object metadata.  All
                      selection expressions from both sections are ANDed to determine
                      if an object meets the cumulative requirements of the selector.'
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  name:
                    description: 'Name is the name of an object.  If defined, it will
                      match against objects with the specified name.  Name also supports
                      a prefix or suffix glob.  For example, ` + "`" + `name: pod-*` + "`" + ` would match
                      both ` + "`" + `pod-a` + "`" + ` and ` + "`" + `pod-b` + "`" + `, and ` + "`" + `name: *-pod` + "`" + ` would match both
                      ` + "`" + `a-pod` + "`" + ` and ` + "`" + `b-pod` + "`" + `.'
                    pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                    type: string
                  namespaceSelector:
                    description: NamespaceSelector is a label selector against an
                      object's containing namespace or the object itself, if the object
                      is a namespace.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    description: 'Namespaces is a list of namespace names. If defined,
                      a constraint only applies to resources in a listed namespace.  Namespaces
                      also supports a prefix or suffix based glob.  For example, ` + "`" + `namespaces:
                      [kube-*]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and
                      ` + "`" + `namespaces: [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  scope:
                    description: Scope determines if cluster-scoped and/or namespaced-scoped
                      resources are matched.  Accepts ` + "`" + `*` + "`" + `, ` + "`" + `Cluster` + "`" + `, or ` + "`" + `Namespaced` + "`" + `.
                      (defaults to ` + "`" + `*` + "`" + `)
                    type: string
                  source:
                    description: Source determines whether generated or original resources
                      are matched. Accepts ` + "`" + `Generated` + "`" + `|` + "`" + `Original` + "`" + `|` + "`" + `All` + "`" + ` (defaults
                      to ` + "`" + `All` + "`" + `). A value of ` + "`" + `Generated` + "`" + ` will only match generated
                      resources, while ` + "`" + `Original` + "`" + ` will only match regular resources.
                    enum:
                    - All
                    - Generated
                    - Original
                    type: string
                type: object
              parameters:
                description: Parameters define the behavior of the mutator.
                properties:
                  assign:
                    description: Assign.value holds the value to be assigned
                    properties:
                      externalData:
                        description: ExternalData describes the external data provider
                          to be used for mutation.
                        properties:
                          dataSource:
                            default: ValueAtLocation
                            description: DataSource specifies where to extract the
                              data that will be sent to the external data provider
                              as parameters.
                            enum:
                            - ValueAtLocation
                            - Username
                            type: string
                          default:
                            description: Default specifies the default value to use
                              when the external data provider returns an error and
                              the failure policy is set to "UseDefault".
                            type: string
                          failurePolicy:
                            default: Fail
                            description: FailurePolicy specifies the policy to apply
                              when the external data provider returns an error.
                            enum:
                            - UseDefault
                            - Ignore
                            - Fail
                            type: string
                          provider:
                            description: Provider is the name of the external data
                              provider.
                            type: string
                        type: object
                      fromMetadata:
                        description: FromMetadata assigns a value from the specified
                          metadata field.
                        properties:
                          field:
                            description: Field specifies which metadata field provides
                              the assigned value. Valid fields are ` + "`" + `namespace` + "`" + ` and
                              ` + "`" + `name` + "`" + `.
                            type: string
                        type: object
                      value:
                        description: Value is a constant value that will be assigned
                          to ` + "`" + `location` + "`" + `
                        x-kubernetes-preserve-unknown-fields: true
                    type: object
                  pathTests:
                    items:
                      description: "PathTest allows the user to customize how the
                        mutation works if parent paths are missing. It traverses the
                        list in order. All sub paths are tested against the provided
                        condition, if the test fails, the mutation is not applied.
                        All ` + "`" + `subPath` + "`" + ` entries must be a prefix of ` + "`" + `location` + "`" + `. Any
                        glob characters will take on the same value as was used to
                        expand the matching glob in ` + "`" + `location` + "`" + `. \n Available Tests:
                        * MustExist    - the path must exist or do not mutate * MustNotExist
                        - the path must not exist or do not mutate."
                      properties:
                        condition:
                          description: Condition describes whether the path either
                            MustExist or MustNotExist in the original object
                          enum:
                          - MustExist
                          - MustNotExist
                          type: string
                        subPath:
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: AssignStatus defines the observed state of Assign.
            properties:
              byPod:
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus.
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system.
                        properties:
                          message:
                            type: string
                          type:
                            description: Type indicates a specific class of error
                              for use by controller code. If not present, the error
                              should be treated as not matching any known type.
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_assign.mutations.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: assignmetadata.mutations.gatekeeper.sh
spec:
  group: mutations.gatekeeper.sh
  names:
    kind: AssignMetadata
    listKind: AssignMetadataList
    plural: assignmetadata
    singular: assignmetadata
  preserveUnknownFields: false
  scope: Cluster
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: AssignMetadata is the Schema for the assignmetadata API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            properties:
              name:
                maxLength: 63
                type: string
            type: object
          spec:
            description: AssignMetadataSpec defines the desired state of AssignMetadata.
            properties:
              location:
                type: string
              match:
                description: Match selects objects to apply mutations to.
                properties:
                  excludedNamespaces:
                    description: 'ExcludedNamespaces is a list of namespace names.
                      If defined, a constraint only applies to resources not in a
                      listed namespace. ExcludedNamespaces also supports a prefix
                      or suffix based glob.  For example, ` + "`" + `excludedNamespaces: [kube-*]` + "`" + `
                      matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and ` + "`" + `excludedNamespaces:
                      [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: 'LabelSelector is the combination of two optional
                      fields: ` + "`" + `matchLabels` + "`" + ` and ` + "`" + `matchExpressions` + "`" + `.  These two fields
                      provide different methods of selecting or excluding k8s objects
                      based on the label keys and values included in object metadata.  All
                      selection expressions from both sections are ANDed to determine
                      if an object meets the cumulative requirements of the selector.'
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  name:
                    description: 'Name is the name of an object.  If defined, it will
                      match against objects with the specified name.  Name also supports
                      a prefix or suffix glob.  For example, ` + "`" + `name: pod-*` + "`" + ` would match
                      both ` + "`" + `pod-a` + "`" + ` and ` + "`" + `pod-b` + "`" + `, and ` + "`" + `name: *-pod` + "`" + ` would match both
                      ` + "`" + `a-pod` + "`" + ` and ` + "`" + `b-pod` + "`" + `.'
                    pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                    type: string
                  namespaceSelector:
                    description: NamespaceSelector is a label selector against an
                      object's containing namespace or the object itself, if the object
                      is a namespace.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    description: 'Namespaces is a list of namespace names. If defined,
                      a constraint only applies to resources in a listed namespace.  Namespaces
                      also supports a prefix or suffix based glob.  For example, ` + "`" + `namespaces:
                      [kube-*]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and
                      ` + "`" + `namespaces: [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  scope:
                    description: Scope determines if cluster-scoped and/or namespaced-scoped
                      resources are matched.  Accepts ` + "`" + `*` + "`" + `, ` + "`" + `Cluster` + "`" + `, or ` + "`" + `Namespaced` + "`" + `.
                      (defaults to ` + "`" + `*` + "`" + `)
                    type: string
                  source:
                    description: Source determines whether generated or original resources
                      are matched. Accepts ` + "`" + `Generated` + "`" + `|` + "`" + `Original` + "`" + `|` + "`" + `All` + "`" + ` (defaults
                      to ` + "`" + `All` + "`" + `). A value of ` + "`" + `Generated` + "`" + ` will only match generated
                      resources, while ` + "`" + `Original` + "`" + ` will only match regular resources.
                    enum:
                    - All
                    - Generated
                    - Original
                    type: string
                type: object
              parameters:
                properties:
                  assign:
                    description: Assign.value holds the value to be assigned
                    properties:
                      externalData:
                        description: ExternalData describes the external data provider
                          to be used for mutation.
                        properties:
                          dataSource:
                            default: ValueAtLocation
                            description: DataSource specifies where to extract the
                              data that will be sent to the external data provider
                              as parameters.
                            enum:
                            - ValueAtLocation
                            - Username
                            type: string
                          default:
                            description: Default specifies the default value to use
                              when the external data provider returns an error and
                              the failure policy is set to "UseDefault".
                            type: string
                          failurePolicy:
                            default: Fail
                            description: FailurePolicy specifies the policy to apply
                              when the external data provider returns an error.
                            enum:
                            - UseDefault
                            - Ignore
                            - Fail
                            type: string
                          provider:
                            description: Provider is the name of the external data
                              provider.
                            type: string
                        type: object
                      fromMetadata:
                        description: FromMetadata assigns a value from the specified
                          metadata field.
                        properties:
                          field:
                            description: Field specifies which metadata field provides
                              the assigned value. Valid fields are ` + "`" + `namespace` + "`" + ` and
                              ` + "`" + `name` + "`" + `.
                            type: string
                        type: object
                      value:
                        description: Value is a constant value that will be assigned
                          to ` + "`" + `location` + "`" + `
                        x-kubernetes-preserve-unknown-fields: true
                    type: object
                type: object
            type: object
          status:
            description: AssignMetadataStatus defines the observed state of AssignMetadata.
            properties:
              byPod:
                description: 'INSERT ADDITIONAL STATUS FIELD - define observed state
                  of cluster Important: Run "make" to regenerate code after modifying
                  this file'
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus.
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system.
                        properties:
                          message:
                            type: string
                          type:
                            description: Type indicates a specific class of error
                              for use by controller code. If not present, the error
                              should be treated as not matching any known type.
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: AssignMetadata is the Schema for the assignmetadata API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AssignMetadataSpec defines the desired state of AssignMetadata.
            properties:
              location:
                type: string
              match:
                description: Match selects objects to apply mutations to.
                properties:
                  excludedNamespaces:
                    description: 'ExcludedNamespaces is a list of namespace names.
                      If defined, a constraint only applies to resources not in a
                      listed namespace. ExcludedNamespaces also supports a prefix
                      or suffix based glob.  For example, ` + "`" + `excludedNamespaces: [kube-*]` + "`" + `
                      matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and ` + "`" + `excludedNamespaces:
                      [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: 'LabelSelector is the combination of two optional
                      fields: ` + "`" + `matchLabels` + "`" + ` and ` + "`" + `matchExpressions` + "`" + `.  These two fields
                      provide different methods of selecting or excluding k8s objects
                      based on the label keys and values included in object metadata.  All
                      selection expressions from both sections are ANDed to determine
                      if an object meets the cumulative requirements of the selector.'
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  name:
                    description: 'Name is the name of an object.  If defined, it will
                      match against objects with the specified name.  Name also supports
                      a prefix or suffix glob.  For example, ` + "`" + `name: pod-*` + "`" + ` would match
                      both ` + "`" + `pod-a` + "`" + ` and ` + "`" + `pod-b` + "`" + `, and ` + "`" + `name: *-pod` + "`" + ` would match both
                      ` + "`" + `a-pod` + "`" + ` and ` + "`" + `b-pod` + "`" + `.'
                    pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                    type: string
                  namespaceSelector:
                    description: NamespaceSelector is a label selector against an
                      object's containing namespace or the object itself, if the object
                      is a namespace.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    description: 'Namespaces is a list of namespace names. If defined,
                      a constraint only applies to resources in a listed namespace.  Namespaces
                      also supports a prefix or suffix based glob.  For example, ` + "`" + `namespaces:
                      [kube-*]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and
                      ` + "`" + `namespaces: [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  scope:
                    description: Scope determines if cluster-scoped and/or namespaced-scoped
                      resources are matched.  Accepts ` + "`" + `*` + "`" + `, ` + "`" + `Cluster` + "`" + `, or ` + "`" + `Namespaced` + "`" + `.
                      (defaults to ` + "`" + `*` + "`" + `)
                    type: string
                  source:
                    description: Source determines whether generated or original resources
                      are matched. Accepts ` + "`" + `Generated` + "`" + `|` + "`" + `Original` + "`" + `|` + "`" + `All` + "`" + ` (defaults
                      to ` + "`" + `All` + "`" + `). A value of ` + "`" + `Generated` + "`" + ` will only match generated
                      resources, while ` + "`" + `Original` + "`" + ` will only match regular resources.
                    enum:
                    - All
                    - Generated
                    - Original
                    type: string
                type: object
              parameters:
                properties:
                  assign:
                    description: Assign.value holds the value to be assigned
                    properties:
                      externalData:
                        description: ExternalData describes the external data provider
                          to be used for mutation.
                        properties:
                          dataSource:
                            default: ValueAtLocation
                            description: DataSource specifies where to extract the
                              data that will be sent to the external data provider
                              as parameters.
                            enum:
                            - ValueAtLocation
                            - Username
                            type: string
                          default:
                            description: Default specifies the default value to use
                              when the external data provider returns an error and
                              the failure policy is set to "UseDefault".
                            type: string
                          failurePolicy:
                            default: Fail
                            description: FailurePolicy specifies the policy to apply
                              when the external data provider returns an error.
                            enum:
                            - UseDefault
                            - Ignore
                            - Fail
                            type: string
                          provider:
                            description: Provider is the name of the external data
                              provider.
                            type: string
                        type: object
                      fromMetadata:
                        description: FromMetadata assigns a value from the specified
                          metadata field.
                        properties:
                          field:
                            description: Field specifies which metadata field provides
                              the assigned value. Valid fields are ` + "`" + `namespace` + "`" + ` and
                              ` + "`" + `name` + "`" + `.
                            type: string
                        type: object
                      value:
                        description: Value is a constant value that will be assigned
                          to ` + "`" + `location` + "`" + `
                        x-kubernetes-preserve-unknown-fields: true
                    type: object
                type: object
            type: object
          status:
            description: AssignMetadataStatus defines the observed state of AssignMetadata.
            properties:
              byPod:
                description: 'INSERT ADDITIONAL STATUS FIELD - define observed state
                  of cluster Important: Run "make" to regenerate code after modifying
                  this file'
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus.
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system.
                        properties:
                          message:
                            type: string
                          type:
                            description: Type indicates a specific class of error
                              for use by controller code. If not present, the error
                              should be treated as not matching any known type.
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: AssignMetadata is the Schema for the assignmetadata API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AssignMetadataSpec defines the desired state of AssignMetadata.
            properties:
              location:
                type: string
              match:
                description: Match selects objects to apply mutations to.
                properties:
                  excludedNamespaces:
                    description: 'ExcludedNamespaces is a list of namespace names.
                      If defined, a constraint only applies to resources not in a
                      listed namespace. ExcludedNamespaces also supports a prefix
                      or suffix based glob.  For example, ` + "`" + `excludedNamespaces: [kube-*]` + "`" + `
                      matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and ` + "`" + `excludedNamespaces:
                      [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: 'LabelSelector is the combination of two optional
                      fields: ` + "`" + `matchLabels` + "`" + ` and ` + "`" + `matchExpressions` + "`" + `.  These two fields
                      provide different methods of selecting or excluding k8s objects
                      based on the label keys and values included in object metadata.  All
                      selection expressions from both sections are ANDed to determine
                      if an object meets the cumulative requirements of the selector.'
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  name:
                    description: 'Name is the name of an object.  If defined, it will
                      match against objects with the specified name.  Name also supports
                      a prefix or suffix glob.  For example, ` + "`" + `name: pod-*` + "`" + ` would match
                      both ` + "`" + `pod-a` + "`" + ` and ` + "`" + `pod-b` + "`" + `, and ` + "`" + `name: *-pod` + "`" + ` would match both
                      ` + "`" + `a-pod` + "`" + ` and ` + "`" + `b-pod` + "`" + `.'
                    pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                    type: string
                  namespaceSelector:
                    description: NamespaceSelector is a label selector against an
                      object's containing namespace or the object itself, if the object
                      is a namespace.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    description: 'Namespaces is a list of namespace names. If defined,
                      a constraint only applies to resources in a listed namespace.  Namespaces
                      also supports a prefix or suffix based glob.  For example, ` + "`" + `namespaces:
                      [kube-*]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and
                      ` + "`" + `namespaces: [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  scope:
                    description: Scope determines if cluster-scoped and/or namespaced-scoped
                      resources are matched.  Accepts ` + "`" + `*` + "`" + `, ` + "`" + `Cluster` + "`" + `, or ` + "`" + `Namespaced` + "`" + `.
                      (defaults to ` + "`" + `*` + "`" + `)
                    type: string
                  source:
                    description: Source determines whether generated or original resources
                      are matched. Accepts ` + "`" + `Generated` + "`" + `|` + "`" + `Original` + "`" + `|` + "`" + `All` + "`" + ` (defaults
                      to ` + "`" + `All` + "`" + `). A value of ` + "`" + `Generated` + "`" + ` will only match generated
                      resources, while ` + "`" + `Original` + "`" + ` will only match regular resources.
                    enum:
                    - All
                    - Generated
                    - Original
                    type: string
                type: object
              parameters:
                properties:
                  assign:
                    description: Assign.value holds the value to be assigned
                    properties:
                      externalData:
                        description: ExternalData describes the external data provider
                          to be used for mutation.
                        properties:
                          dataSource:
                            default: ValueAtLocation
                            description: DataSource specifies where to extract the
                              data that will be sent to the external data provider
                              as parameters.
                            enum:
                            - ValueAtLocation
                            - Username
                            type: string
                          default:
                            description: Default specifies the default value to use
                              when the external data provider returns an error and
                              the failure policy is set to "UseDefault".
                            type: string
                          failurePolicy:
                            default: Fail
                            description: FailurePolicy specifies the policy to apply
                              when the external data provider returns an error.
                            enum:
                            - UseDefault
                            - Ignore
                            - Fail
                            type: string
                          provider:
                            description: Provider is the name of the external data
                              provider.
                            type: string
                        type: object
                      fromMetadata:
                        description: FromMetadata assigns a value from the specified
                          metadata field.
                        properties:
                          field:
                            description: Field specifies which metadata field provides
                              the assigned value. Valid fields are ` + "`" + `namespace` + "`" + ` and
                              ` + "`" + `name` + "`" + `.
                            type: string
                        type: object
                      value:
                        description: Value is a constant value that will be assigned
                          to ` + "`" + `location` + "`" + `
                        x-kubernetes-preserve-unknown-fields: true
                    type: object
                type: object
            type: object
          status:
            description: AssignMetadataStatus defines the observed state of AssignMetadata.
            properties:
              byPod:
                description: 'INSERT ADDITIONAL STATUS FIELD - define observed state
                  of cluster Important: Run "make" to regenerate code after modifying
                  this file'
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus.
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system.
                        properties:
                          message:
                            type: string
                          type:
                            description: Type indicates a specific class of error
                              for use by controller code. If not present, the error
                              should be treated as not matching any known type.
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_assignmetadata.mutations.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: configs.config.gatekeeper.sh
spec:
  group: config.gatekeeper.sh
  names:
    kind: Config
    listKind: ConfigList
    plural: configs
    singular: config
  preserveUnknownFields: false
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: Config is the Schema for the configs API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ConfigSpec defines the desired state of Config.
            properties:
              match:
                description: Configuration for namespace exclusion
                items:
                  properties:
                    excludedNamespaces:
                      items:
                        description: 'A string that supports globbing at its front
                          or end. Ex: "kube-*" will match "kube-system" or "kube-public",
                          "*-system" will match "kube-system" or "gatekeeper-system".  The
                          asterisk is required for wildcard matching.'
                        pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                        type: string
                      type: array
                    processes:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              readiness:
                description: Configuration for readiness tracker
                properties:
                  statsEnabled:
                    type: boolean
                type: object
              sync:
                description: Configuration for syncing k8s objects
                properties:
                  syncOnly:
                    description: If non-empty, only entries on this list will be replicated
                      into OPA
                    items:
                      properties:
                        group:
                          type: string
                        kind:
                          type: string
                        version:
                          type: string
                      type: object
                    type: array
                type: object
              validation:
                description: Configuration for validation
                properties:
                  traces:
                    description: List of requests to trace. Both "user" and "kinds"
                      must be specified
                    items:
                      properties:
                        dump:
                          description: Also dump the state of OPA with the trace.
                            Set to ` + "`" + `All` + "`" + ` to dump everything.
                          type: string
                        kind:
                          description: Only trace requests of the following GroupVersionKind
                          properties:
                            group:
                              type: string
                            kind:
                              type: string
                            version:
                              type: string
                          type: object
                        user:
                          description: Only trace requests from the specified user
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: ConfigStatus defines the observed state of Config.
            type: object
        type: object
    served: true
    storage: true
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_configs.config.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: constraintpodstatuses.status.gatekeeper.sh
spec:
  group: status.gatekeeper.sh
  names:
    kind: ConstraintPodStatus
    listKind: ConstraintPodStatusList
    plural: constraintpodstatuses
    singular: constraintpodstatus
  preserveUnknownFields: false
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: ConstraintPodStatus is the Schema for the constraintpodstatuses
          API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          status:
            description: ConstraintPodStatusStatus defines the observed state of ConstraintPodStatus.
            properties:
              constraintUID:
                description: Storing the constraint UID allows us to detect drift,
                  such as when a constraint has been recreated after its CRD was deleted
                  out from under it, interrupting the watch
                type: string
              enforced:
                type: boolean
              errors:
                items:
                  description: Error represents a single error caught while adding
                    a constraint to OPA.
                  properties:
                    code:
                      type: string
                    location:
                      type: string
                    message:
                      type: string
                  required:
                  - code
                  - message
                  type: object
                type: array
              id:
                type: string
              observedGeneration:
                format: int64
                type: integer
              operations:
                items:
                  type: string
                type: array
            type: object
        type: object
    served: true
    storage: true
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: constrainttemplatepodstatuses.status.gatekeeper.sh
spec:
  group: status.gatekeeper.sh
  names:
    kind: ConstraintTemplatePodStatus
    listKind: ConstraintTemplatePodStatusList
    plural: constrainttemplatepodstatuses
    singular: constrainttemplatepodstatus
  preserveUnknownFields: false
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: ConstraintTemplatePodStatus is the Schema for the constrainttemplatepodstatuses
          API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          status:
            description: ConstraintTemplatePodStatusStatus defines the observed state
              of ConstraintTemplatePodStatus.
            properties:
              errors:
                items:
                  description: CreateCRDError represents a single error caught during
                    parsing, compiling, etc.
                  properties:
                    code:
                      type: string
                    location:
                      type: string
                    message:
                      type: string
                  required:
                  - code
                  - message
                  type: object
                type: array
              id:
                description: 'Important: Run "make" to regenerate code after modifying
                  this file'
                type: string
              observedGeneration:
                format: int64
                type: integer
              operations:
                items:
                  type: string
                type: array
              templateUID:
                description: UID is a type that holds unique ID values, including
                  UUIDs.  Because we don't ONLY use UUIDs, this is an alias to string.  Being
                  a type captures intent and helps make sure that UIDs and names do
                  not get conflated.
                type: string
            type: object
        type: object
    served: true
    storage: true
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  labels:
    gatekeeper.sh/system: "yes"
  name: constrainttemplates.templates.gatekeeper.sh
spec:
  group: templates.gatekeeper.sh
  names:
    kind: ConstraintTemplate
    listKind: ConstraintTemplateList
    plural: constrainttemplates
    singular: constrainttemplate
  preserveUnknownFields: false
  scope: Cluster
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: ConstraintTemplate is the Schema for the constrainttemplates
          API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ConstraintTemplateSpec defines the desired state of ConstraintTemplate.
            properties:
              crd:
                properties:
                  spec:
                    properties:
                      names:
                        properties:
                          kind:
                            type: string
                          shortNames:
                            items:
                              type: string
                            type: array
                        type: object
                      validation:
                        default:
                          legacySchema: false
                        properties:
                          legacySchema:
                            default: false
                            type: boolean
                          openAPIV3Schema:
                            type: object
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                    type: object
                type: object
              targets:
                items:
                  properties:
                    libs:
                      items:
                        type: string
                      type: array
                    rego:
                      type: string
                    target:
                      type: string
                  type: object
                type: array
            type: object
          status:
            description: ConstraintTemplateStatus defines the observed state of ConstraintTemplate.
            properties:
              byPod:
                items:
                  description: ByPodStatus defines the observed state of ConstraintTemplate
                    as seen by an individual controller
                  properties:
                    errors:
                      items:
                        description: CreateCRDError represents a single error caught
                          during parsing, compiling, etc.
                        properties:
                          code:
                            type: string
                          location:
                            type: string
                          message:
                            type: string
                        required:
                        - code
                        - message
                        type: object
                      type: array
                    id:
                      description: a unique identifier for the pod that wrote the
                        status
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
                type: array
              created:
                type: boolean
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: ConstraintTemplate is the Schema for the constrainttemplates
          API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ConstraintTemplateSpec defines the desired state of ConstraintTemplate.
            properties:
              crd:
                properties:
                  spec:
                    properties:
                      names:
                        properties:
                          kind:
                            type: string
                          shortNames:
                            items:
                              type: string
                            type: array
                        type: object
                      validation:
                        default:
                          legacySchema: true
                        properties:
                          legacySchema:
                            default: true
                            type: boolean
                          openAPIV3Schema:
                            type: object
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                    type: object
                type: object
              targets:
                items:
                  properties:
                    libs:
                      items:
                        type: string
                      type: array
                    rego:
                      type: string
                    target:
                      type: string
                  type: object
                type: array
            type: object
          status:
            description: ConstraintTemplateStatus defines the observed state of ConstraintTemplate.
            properties:
              byPod:
                items:
                  description: ByPodStatus defines the observed state of ConstraintTemplate
                    as seen by an individual controller
                  properties:
                    errors:
                      items:
                        description: CreateCRDError represents a single error caught
                          during parsing, compiling, etc.
                        properties:
                          code:
                            type: string
                          location:
                            type: string
                          message:
                            type: string
                        required:
                        - code
                        - message
                        type: object
                      type: array
                    id:
                      description: a unique identifier for the pod that wrote the
                        status
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
                type: array
              created:
                type: boolean
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: ConstraintTemplate is the Schema for the constrainttemplates
          API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ConstraintTemplateSpec defines the desired state of ConstraintTemplate.
            properties:
              crd:
                properties:
                  spec:
                    properties:
                      names:
                        properties:
                          kind:
                            type: string
                          shortNames:
                            items:
                              type: string
                            type: array
                        type: object
                      validation:
                        default:
                          legacySchema: true
                        properties:
                          legacySchema:
                            default: true
                            type: boolean
                          openAPIV3Schema:
                            type: object
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                    type: object
                type: object
              targets:
                items:
                  properties:
                    libs:
                      items:
                        type: string
                      type: array
                    rego:
                      type: string
                    target:
                      type: string
                  type: object
                type: array
            type: object
          status:
            description: ConstraintTemplateStatus defines the observed state of ConstraintTemplate.
            properties:
              byPod:
                items:
                  description: ByPodStatus defines the observed state of ConstraintTemplate
                    as seen by an individual controller
                  properties:
                    errors:
                      items:
                        description: CreateCRDError represents a single error caught
                          during parsing, compiling, etc.
                        properties:
                          code:
                            type: string
                          location:
                            type: string
                          message:
                            type: string
                        required:
                        - code
                        - message
                        type: object
                      type: array
                    id:
                      description: a unique identifier for the pod that wrote the
                        status
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
                type: array
              created:
                type: boolean
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_expansiontemplateExpansionGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: expansiontemplate.expansion.gatekeeper.sh
spec:
  group: expansion.gatekeeper.sh
  names:
    kind: ExpansionTemplate
    listKind: ExpansionTemplateList
    plural: expansiontemplate
    singular: expansiontemplate
  preserveUnknownFields: false
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: ExpansionTemplate is the Schema for the ExpansionTemplate API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ExpansionTemplateSpec defines the desired state of ExpansionTemplate.
            properties:
              applyTo:
                description: ApplyTo lists the specific groups, versions and kinds
                  of generator resources which will be expanded.
                items:
                  description: ApplyTo determines what GVKs items the mutation should
                    apply to. Globs are not allowed.
                  properties:
                    groups:
                      items:
                        type: string
                      type: array
                    kinds:
                      items:
                        type: string
                      type: array
                    versions:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              enforcementAction:
                description: EnforcementAction specifies the enforcement action to
                  be used for resources matching the ExpansionTemplate. Specifying
                  an empty value will use the enforcement action specified by the
                  Constraint in violation.
                type: string
              generatedGVK:
                description: GeneratedGVK specifies the GVK of the resources which
                  the generator resource creates.
                properties:
                  group:
                    type: string
                  kind:
                    type: string
                  version:
                    type: string
                type: object
              templateSource:
                description: TemplateSource specifies the source field on the generator
                  resource to use as the base for expanded resource. For Pod-creating
                  generators, this is usually spec.template
                type: string
            type: object
        type: object
    served: true
    storage: true
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_expansiontemplateExpansionGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_expansiontemplateExpansionGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_expansiontemplateExpansionGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_expansiontemplateExpansionGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_expansiontemplate.expansion.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_modifysetMutationsGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: modifyset.mutations.gatekeeper.sh
spec:
  group: mutations.gatekeeper.sh
  names:
    kind: ModifySet
    listKind: ModifySetList
    plural: modifyset
    singular: modifyset
  preserveUnknownFields: false
  scope: Cluster
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: ModifySet allows the user to modify non-keyed lists, such as
          the list of arguments to a container.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            properties:
              name:
                maxLength: 63
                type: string
            type: object
          spec:
            description: ModifySetSpec defines the desired state of ModifySet.
            properties:
              applyTo:
                description: ApplyTo lists the specific groups, versions and kinds
                  a mutation will be applied to. This is necessary because every mutation
                  implies part of an object schema and object schemas are associated
                  with specific GVKs.
                items:
                  description: ApplyTo determines what GVKs items the mutation should
                    apply to. Globs are not allowed.
                  properties:
                    groups:
                      items:
                        type: string
                      type: array
                    kinds:
                      items:
                        type: string
                      type: array
                    versions:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              location:
                description: 'Location describes the path to be mutated, for example:
                  ` + "`" + `spec.containers[name: main].args` + "`" + `.'
                type: string
              match:
                description: Match allows the user to limit which resources get mutated.
                  Individual match criteria are AND-ed together. An undefined match
                  criteria matches everything.
                properties:
                  excludedNamespaces:
                    description: 'ExcludedNamespaces is a list of namespace names.
                      If defined, a constraint only applies to resources not in a
                      listed namespace. ExcludedNamespaces also supports a prefix
                      or suffix based glob.  For example, ` + "`" + `excludedNamespaces: [kube-*]` + "`" + `
                      matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and ` + "`" + `excludedNamespaces:
                      [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: 'LabelSelector is the combination of two optional
                      fields: ` + "`" + `matchLabels` + "`" + ` and ` + "`" + `matchExpressions` + "`" + `.  These two fields
                      provide different methods of selecting or excluding k8s objects
                      based on the label keys and values included in object metadata.  All
                      selection expressions from both sections are ANDed to determine
                      if an object meets the cumulative requirements of the selector.'
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  name:
                    description: 'Name is the name of an object.  If defined, it will
                      match against objects with the specified name.  Name also supports
                      a prefix or suffix glob.  For example, ` + "`" + `name: pod-*` + "`" + ` would match
                      both ` + "`" + `pod-a` + "`" + ` and ` + "`" + `pod-b` + "`" + `, and ` + "`" + `name: *-pod` + "`" + ` would match both
                      ` + "`" + `a-pod` + "`" + ` and ` + "`" + `b-pod` + "`" + `.'
                    pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                    type: string
                  namespaceSelector:
                    description: NamespaceSelector is a label selector against an
                      object's containing namespace or the object itself, if the object
                      is a namespace.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    description: 'Namespaces is a list of namespace names. If defined,
                      a constraint only applies to resources in a listed namespace.  Namespaces
                      also supports a prefix or suffix based glob.  For example, ` + "`" + `namespaces:
                      [kube-*]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and
                      ` + "`" + `namespaces: [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  scope:
                    description: Scope determines if cluster-scoped and/or namespaced-scoped
                      resources are matched.  Accepts ` + "`" + `*` + "`" + `, ` + "`" + `Cluster` + "`" + `, or ` + "`" + `Namespaced` + "`" + `.
                      (defaults to ` + "`" + `*` + "`" + `)
                    type: string
                  source:
                    description: Source determines whether generated or original resources
                      are matched. Accepts ` + "`" + `Generated` + "`" + `|` + "`" + `Original` + "`" + `|` + "`" + `All` + "`" + ` (defaults
                      to ` + "`" + `All` + "`" + `). A value of ` + "`" + `Generated` + "`" + ` will only match generated
                      resources, while ` + "`" + `Original` + "`" + ` will only match regular resources.
                    enum:
                    - All
                    - Generated
                    - Original
                    type: string
                type: object
              parameters:
                description: Parameters define the behavior of the mutator.
                properties:
                  operation:
                    default: merge
                    description: Operation describes whether values should be merged
                      in ("merge"), or pruned ("prune"). Default value is "merge"
                    enum:
                    - merge
                    - prune
                    type: string
                  pathTests:
                    description: PathTests are a series of existence tests that can
                      be checked before a mutation is applied
                    items:
                      description: "PathTest allows the user to customize how the
                        mutation works if parent paths are missing. It traverses the
                        list in order. All sub paths are tested against the provided
                        condition, if the test fails, the mutation is not applied.
                        All ` + "`" + `subPath` + "`" + ` entries must be a prefix of ` + "`" + `location` + "`" + `. Any
                        glob characters will take on the same value as was used to
                        expand the matching glob in ` + "`" + `location` + "`" + `. \n Available Tests:
                        * MustExist    - the path must exist or do not mutate * MustNotExist
                        - the path must not exist or do not mutate."
                      properties:
                        condition:
                          description: Condition describes whether the path either
                            MustExist or MustNotExist in the original object
                          enum:
                          - MustExist
                          - MustNotExist
                          type: string
                        subPath:
                          type: string
                      type: object
                    type: array
                  values:
                    description: Values describes the values provided to the operation
                      as ` + "`" + `values.fromList` + "`" + `.
                    type: object
                    x-kubernetes-preserve-unknown-fields: true
                type: object
            type: object
          status:
            description: ModifySetStatus defines the observed state of ModifySet.
            properties:
              byPod:
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus.
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system.
                        properties:
                          message:
                            type: string
                          type:
                            description: Type indicates a specific class of error
                              for use by controller code. If not present, the error
                              should be treated as not matching any known type.
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: ModifySet allows the user to modify non-keyed lists, such as
          the list of arguments to a container.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ModifySetSpec defines the desired state of ModifySet.
            properties:
              applyTo:
                description: ApplyTo lists the specific groups, versions and kinds
                  a mutation will be applied to. This is necessary because every mutation
                  implies part of an object schema and object schemas are associated
                  with specific GVKs.
                items:
                  description: ApplyTo determines what GVKs items the mutation should
                    apply to. Globs are not allowed.
                  properties:
                    groups:
                      items:
                        type: string
                      type: array
                    kinds:
                      items:
                        type: string
                      type: array
                    versions:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              location:
                description: 'Location describes the path to be mutated, for example:
                  ` + "`" + `spec.containers[name: main].args` + "`" + `.'
                type: string
              match:
                description: Match allows the user to limit which resources get mutated.
                  Individual match criteria are AND-ed together. An undefined match
                  criteria matches everything.
                properties:
                  excludedNamespaces:
                    description: 'ExcludedNamespaces is a list of namespace names.
                      If defined, a constraint only applies to resources not in a
                      listed namespace. ExcludedNamespaces also supports a prefix
                      or suffix based glob.  For example, ` + "`" + `excludedNamespaces: [kube-*]` + "`" + `
                      matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and ` + "`" + `excludedNamespaces:
                      [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: 'LabelSelector is the combination of two optional
                      fields: ` + "`" + `matchLabels` + "`" + ` and ` + "`" + `matchExpressions` + "`" + `.  These two fields
                      provide different methods of selecting or excluding k8s objects
                      based on the label keys and values included in object metadata.  All
                      selection expressions from both sections are ANDed to determine
                      if an object meets the cumulative requirements of the selector.'
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  name:
                    description: 'Name is the name of an object.  If defined, it will
                      match against objects with the specified name.  Name also supports
                      a prefix or suffix glob.  For example, ` + "`" + `name: pod-*` + "`" + ` would match
                      both ` + "`" + `pod-a` + "`" + ` and ` + "`" + `pod-b` + "`" + `, and ` + "`" + `name: *-pod` + "`" + ` would match both
                      ` + "`" + `a-pod` + "`" + ` and ` + "`" + `b-pod` + "`" + `.'
                    pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                    type: string
                  namespaceSelector:
                    description: NamespaceSelector is a label selector against an
                      object's containing namespace or the object itself, if the object
                      is a namespace.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    description: 'Namespaces is a list of namespace names. If defined,
                      a constraint only applies to resources in a listed namespace.  Namespaces
                      also supports a prefix or suffix based glob.  For example, ` + "`" + `namespaces:
                      [kube-*]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and
                      ` + "`" + `namespaces: [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  scope:
                    description: Scope determines if cluster-scoped and/or namespaced-scoped
                      resources are matched.  Accepts ` + "`" + `*` + "`" + `, ` + "`" + `Cluster` + "`" + `, or ` + "`" + `Namespaced` + "`" + `.
                      (defaults to ` + "`" + `*` + "`" + `)
                    type: string
                  source:
                    description: Source determines whether generated or original resources
                      are matched. Accepts ` + "`" + `Generated` + "`" + `|` + "`" + `Original` + "`" + `|` + "`" + `All` + "`" + ` (defaults
                      to ` + "`" + `All` + "`" + `). A value of ` + "`" + `Generated` + "`" + ` will only match generated
                      resources, while ` + "`" + `Original` + "`" + ` will only match regular resources.
                    enum:
                    - All
                    - Generated
                    - Original
                    type: string
                type: object
              parameters:
                description: Parameters define the behavior of the mutator.
                properties:
                  operation:
                    default: merge
                    description: Operation describes whether values should be merged
                      in ("merge"), or pruned ("prune"). Default value is "merge"
                    enum:
                    - merge
                    - prune
                    type: string
                  pathTests:
                    description: PathTests are a series of existence tests that can
                      be checked before a mutation is applied
                    items:
                      description: "PathTest allows the user to customize how the
                        mutation works if parent paths are missing. It traverses the
                        list in order. All sub paths are tested against the provided
                        condition, if the test fails, the mutation is not applied.
                        All ` + "`" + `subPath` + "`" + ` entries must be a prefix of ` + "`" + `location` + "`" + `. Any
                        glob characters will take on the same value as was used to
                        expand the matching glob in ` + "`" + `location` + "`" + `. \n Available Tests:
                        * MustExist    - the path must exist or do not mutate * MustNotExist
                        - the path must not exist or do not mutate."
                      properties:
                        condition:
                          description: Condition describes whether the path either
                            MustExist or MustNotExist in the original object
                          enum:
                          - MustExist
                          - MustNotExist
                          type: string
                        subPath:
                          type: string
                      type: object
                    type: array
                  values:
                    description: Values describes the values provided to the operation
                      as ` + "`" + `values.fromList` + "`" + `.
                    type: object
                    x-kubernetes-preserve-unknown-fields: true
                type: object
            type: object
          status:
            description: ModifySetStatus defines the observed state of ModifySet.
            properties:
              byPod:
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus.
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system.
                        properties:
                          message:
                            type: string
                          type:
                            description: Type indicates a specific class of error
                              for use by controller code. If not present, the error
                              should be treated as not matching any known type.
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: ModifySet allows the user to modify non-keyed lists, such as
          the list of arguments to a container.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ModifySetSpec defines the desired state of ModifySet.
            properties:
              applyTo:
                description: ApplyTo lists the specific groups, versions and kinds
                  a mutation will be applied to. This is necessary because every mutation
                  implies part of an object schema and object schemas are associated
                  with specific GVKs.
                items:
                  description: ApplyTo determines what GVKs items the mutation should
                    apply to. Globs are not allowed.
                  properties:
                    groups:
                      items:
                        type: string
                      type: array
                    kinds:
                      items:
                        type: string
                      type: array
                    versions:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              location:
                description: 'Location describes the path to be mutated, for example:
                  ` + "`" + `spec.containers[name: main].args` + "`" + `.'
                type: string
              match:
                description: Match allows the user to limit which resources get mutated.
                  Individual match criteria are AND-ed together. An undefined match
                  criteria matches everything.
                properties:
                  excludedNamespaces:
                    description: 'ExcludedNamespaces is a list of namespace names.
                      If defined, a constraint only applies to resources not in a
                      listed namespace. ExcludedNamespaces also supports a prefix
                      or suffix based glob.  For example, ` + "`" + `excludedNamespaces: [kube-*]` + "`" + `
                      matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and ` + "`" + `excludedNamespaces:
                      [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: 'LabelSelector is the combination of two optional
                      fields: ` + "`" + `matchLabels` + "`" + ` and ` + "`" + `matchExpressions` + "`" + `.  These two fields
                      provide different methods of selecting or excluding k8s objects
                      based on the label keys and values included in object metadata.  All
                      selection expressions from both sections are ANDed to determine
                      if an object meets the cumulative requirements of the selector.'
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  name:
                    description: 'Name is the name of an object.  If defined, it will
                      match against objects with the specified name.  Name also supports
                      a prefix or suffix glob.  For example, ` + "`" + `name: pod-*` + "`" + ` would match
                      both ` + "`" + `pod-a` + "`" + ` and ` + "`" + `pod-b` + "`" + `, and ` + "`" + `name: *-pod` + "`" + ` would match both
                      ` + "`" + `a-pod` + "`" + ` and ` + "`" + `b-pod` + "`" + `.'
                    pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                    type: string
                  namespaceSelector:
                    description: NamespaceSelector is a label selector against an
                      object's containing namespace or the object itself, if the object
                      is a namespace.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    description: 'Namespaces is a list of namespace names. If defined,
                      a constraint only applies to resources in a listed namespace.  Namespaces
                      also supports a prefix or suffix based glob.  For example, ` + "`" + `namespaces:
                      [kube-*]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `kube-public` + "`" + `, and
                      ` + "`" + `namespaces: [*-system]` + "`" + ` matches both ` + "`" + `kube-system` + "`" + ` and ` + "`" + `gatekeeper-system` + "`" + `.'
                    items:
                      description: 'A string that supports globbing at its front or
                        end. Ex: "kube-*" will match "kube-system" or "kube-public",
                        "*-system" will match "kube-system" or "gatekeeper-system".  The
                        asterisk is required for wildcard matching.'
                      pattern: ^(\*|\*-)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\*|-\*)?$
                      type: string
                    type: array
                  scope:
                    description: Scope determines if cluster-scoped and/or namespaced-scoped
                      resources are matched.  Accepts ` + "`" + `*` + "`" + `, ` + "`" + `Cluster` + "`" + `, or ` + "`" + `Namespaced` + "`" + `.
                      (defaults to ` + "`" + `*` + "`" + `)
                    type: string
                  source:
                    description: Source determines whether generated or original resources
                      are matched. Accepts ` + "`" + `Generated` + "`" + `|` + "`" + `Original` + "`" + `|` + "`" + `All` + "`" + ` (defaults
                      to ` + "`" + `All` + "`" + `). A value of ` + "`" + `Generated` + "`" + ` will only match generated
                      resources, while ` + "`" + `Original` + "`" + ` will only match regular resources.
                    enum:
                    - All
                    - Generated
                    - Original
                    type: string
                type: object
              parameters:
                description: Parameters define the behavior of the mutator.
                properties:
                  operation:
                    default: merge
                    description: Operation describes whether values should be merged
                      in ("merge"), or pruned ("prune"). Default value is "merge"
                    enum:
                    - merge
                    - prune
                    type: string
                  pathTests:
                    description: PathTests are a series of existence tests that can
                      be checked before a mutation is applied
                    items:
                      description: "PathTest allows the user to customize how the
                        mutation works if parent paths are missing. It traverses the
                        list in order. All sub paths are tested against the provided
                        condition, if the test fails, the mutation is not applied.
                        All ` + "`" + `subPath` + "`" + ` entries must be a prefix of ` + "`" + `location` + "`" + `. Any
                        glob characters will take on the same value as was used to
                        expand the matching glob in ` + "`" + `location` + "`" + `. \n Available Tests:
                        * MustExist    - the path must exist or do not mutate * MustNotExist
                        - the path must not exist or do not mutate."
                      properties:
                        condition:
                          description: Condition describes whether the path either
                            MustExist or MustNotExist in the original object
                          enum:
                          - MustExist
                          - MustNotExist
                          type: string
                        subPath:
                          type: string
                      type: object
                    type: array
                  values:
                    description: Values describes the values provided to the operation
                      as ` + "`" + `values.fromList` + "`" + `.
                    type: object
                    x-kubernetes-preserve-unknown-fields: true
                type: object
            type: object
          status:
            description: ModifySetStatus defines the observed state of ModifySet.
            properties:
              byPod:
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus.
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system.
                        properties:
                          message:
                            type: string
                          type:
                            description: Type indicates a specific class of error
                              for use by controller code. If not present, the error
                              should be treated as not matching any known type.
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_modifysetMutationsGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_modifysetMutationsGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_modifysetMutationsGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_modifysetMutationsGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_modifyset.mutations.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: mutatorpodstatuses.status.gatekeeper.sh
spec:
  group: status.gatekeeper.sh
  names:
    kind: MutatorPodStatus
    listKind: MutatorPodStatusList
    plural: mutatorpodstatuses
    singular: mutatorpodstatus
  preserveUnknownFields: false
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: MutatorPodStatus is the Schema for the mutationpodstatuses API.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          status:
            description: MutatorPodStatusStatus defines the observed state of MutatorPodStatus.
            properties:
              enforced:
                type: boolean
              errors:
                items:
                  description: MutatorError represents a single error caught while
                    adding a mutator to a system.
                  properties:
                    message:
                      type: string
                    type:
                      description: Type indicates a specific class of error for use
                        by controller code. If not present, the error should be treated
                        as not matching any known type.
                      type: string
                  required:
                  - message
                  type: object
                type: array
              id:
                type: string
              mutatorUID:
                description: Storing the mutator UID allows us to detect drift, such
                  as when a mutator has been recreated after its CRD was deleted out
                  from under it, interrupting the watch
                type: string
              observedGeneration:
                format: int64
                type: integer
              operations:
                items:
                  type: string
                type: array
            type: object
        type: object
    served: true
    storage: true
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_mutatorpodstatuses.status.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_providersExternaldataGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.10.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: providers.externaldata.gatekeeper.sh
spec:
  group: externaldata.gatekeeper.sh
  names:
    kind: Provider
    listKind: ProviderList
    plural: providers
    singular: provider
  preserveUnknownFields: false
  scope: Cluster
  versions:
  - deprecated: true
    deprecationWarning: externaldata.gatekeeper.sh/v1alpha1 is deprecated. Use externaldata.gatekeeper.sh/v1beta1
      instead.
    name: v1alpha1
    schema:
      openAPIV3Schema:
        description: Provider is the Schema for the Provider API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: Spec defines the Provider specifications.
            properties:
              caBundle:
                description: CABundle is a base64-encoded string that contains the
                  TLS CA bundle in PEM format. It is used to verify the signature
                  of the provider's certificate.
                type: string
              timeout:
                description: Timeout is the timeout when querying the provider.
                type: integer
              url:
                description: URL is the url for the provider. URL is prefixed with
                  http:// or https://.
                type: string
            type: object
        type: object
    served: true
    storage: true
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: Provider is the Schema for the providers API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: Spec defines the Provider specifications.
            properties:
              caBundle:
                description: CABundle is a base64-encoded string that contains the
                  TLS CA bundle in PEM format. It is used to verify the signature
                  of the provider's certificate.
                type: string
              timeout:
                description: Timeout is the timeout when querying the provider.
                type: integer
              url:
                description: URL is the url for the provider. URL is prefixed with
                  http:// or https://.
                type: string
            type: object
        type: object
    served: true
    storage: false
`)

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_providersExternaldataGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_providersExternaldataGatekeeperShYaml, nil
}

func configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_providersExternaldataGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_providersExternaldataGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_providers.externaldata.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApps_v1_deployment_gatekeeperAuditYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: audit-controller
    gatekeeper.sh/operation: audit
    gatekeeper.sh/system: "yes"
  name: gatekeeper-audit
  namespace: gatekeeper-system
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: audit-controller
      gatekeeper.sh/operation: audit
      gatekeeper.sh/system: "yes"
  template:
    metadata:
      annotations:
        container.seccomp.security.alpha.kubernetes.io/manager: runtime/default
      labels:
        control-plane: audit-controller
        gatekeeper.sh/operation: audit
        gatekeeper.sh/system: "yes"
    spec:
      automountServiceAccountToken: true
      containers:
      - args:
        - --operation=audit
        - --operation=status
        - --operation=mutation-status
        - --logtostderr
        - --disable-opa-builtin={http.send}
        command:
        - /manager
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: CONTAINER_NAME
          value: manager
        image: openpolicyagent/gatekeeper:v3.11.1
        imagePullPolicy: Always
        livenessProbe:
          httpGet:
            path: /healthz
            port: 9090
        name: manager
        ports:
        - containerPort: 8888
          name: metrics
          protocol: TCP
        - containerPort: 9090
          name: healthz
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /readyz
            port: 9090
        resources:
          limits:
            cpu: 1000m
            memory: 512Mi
          requests:
            cpu: 100m
            memory: 512Mi
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
            - ALL
          readOnlyRootFilesystem: true
          runAsGroup: 999
          runAsNonRoot: true
          runAsUser: 1000
          seccompProfile:
            type: RuntimeDefault
        volumeMounts:
        - mountPath: /certs
          name: cert
          readOnly: true
        - mountPath: /tmp/audit
          name: tmp-volume
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      serviceAccountName: gatekeeper-admin
      terminationGracePeriodSeconds: 60
      volumes:
      - name: cert
        secret:
          defaultMode: 420
          secretName: gatekeeper-webhook-server-cert
      - emptyDir: {}
        name: tmp-volume
`)

func configGatekeeperRenderedApps_v1_deployment_gatekeeperAuditYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApps_v1_deployment_gatekeeperAuditYaml, nil
}

func configGatekeeperRenderedApps_v1_deployment_gatekeeperAuditYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApps_v1_deployment_gatekeeperAuditYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apps_v1_deployment_gatekeeper-audit.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedApps_v1_deployment_gatekeeperControllerManagerYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
    gatekeeper.sh/operation: webhook
    gatekeeper.sh/system: "yes"
  name: gatekeeper-controller-manager
  namespace: gatekeeper-system
spec:
  replicas: 3
  selector:
    matchLabels:
      control-plane: controller-manager
      gatekeeper.sh/operation: webhook
      gatekeeper.sh/system: "yes"
  template:
    metadata:
      annotations:
        container.seccomp.security.alpha.kubernetes.io/manager: runtime/default
      labels:
        control-plane: controller-manager
        gatekeeper.sh/operation: webhook
        gatekeeper.sh/system: "yes"
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: gatekeeper.sh/operation
                  operator: In
                  values:
                  - webhook
              topologyKey: kubernetes.io/hostname
            weight: 100
      automountServiceAccountToken: true
      containers:
      - args:
        - --port=8443
        - --logtostderr
        - --exempt-namespace=gatekeeper-system
        - --operation=webhook
        - --operation=mutation-webhook
        - --disable-opa-builtin={http.send}
        command:
        - /manager
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: CONTAINER_NAME
          value: manager
        image: openpolicyagent/gatekeeper:v3.11.1
        imagePullPolicy: Always
        livenessProbe:
          httpGet:
            path: /healthz
            port: 9090
        name: manager
        ports:
        - containerPort: 8443
          name: webhook-server
          protocol: TCP
        - containerPort: 8888
          name: metrics
          protocol: TCP
        - containerPort: 9090
          name: healthz
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /readyz
            port: 9090
        resources:
          limits:
            cpu: 1000m
            memory: 512Mi
          requests:
            cpu: 100m
            memory: 512Mi
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
            - ALL
          readOnlyRootFilesystem: true
          runAsGroup: 999
          runAsNonRoot: true
          runAsUser: 1000
          seccompProfile:
            type: RuntimeDefault
        volumeMounts:
        - mountPath: /certs
          name: cert
          readOnly: true
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      serviceAccountName: gatekeeper-admin
      terminationGracePeriodSeconds: 60
      volumes:
      - name: cert
        secret:
          defaultMode: 420
          secretName: gatekeeper-webhook-server-cert
`)

func configGatekeeperRenderedApps_v1_deployment_gatekeeperControllerManagerYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedApps_v1_deployment_gatekeeperControllerManagerYaml, nil
}

func configGatekeeperRenderedApps_v1_deployment_gatekeeperControllerManagerYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedApps_v1_deployment_gatekeeperControllerManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/apps_v1_deployment_gatekeeper-controller-manager.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedPolicy_v1_poddisruptionbudget_gatekeeperControllerManagerYaml = []byte(`apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-controller-manager
  namespace: gatekeeper-system
spec:
  minAvailable: 1
  selector:
    matchLabels:
      control-plane: controller-manager
      gatekeeper.sh/operation: webhook
      gatekeeper.sh/system: "yes"
`)

func configGatekeeperRenderedPolicy_v1_poddisruptionbudget_gatekeeperControllerManagerYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedPolicy_v1_poddisruptionbudget_gatekeeperControllerManagerYaml, nil
}

func configGatekeeperRenderedPolicy_v1_poddisruptionbudget_gatekeeperControllerManagerYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedPolicy_v1_poddisruptionbudget_gatekeeperControllerManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/policy_v1_poddisruptionbudget_gatekeeper-controller-manager.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-manager-role
rules:
- apiGroups:
  - '*'
  resources:
  - '*'
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - admissionregistration.k8s.io
  resourceNames:
  - gatekeeper-mutating-webhook-configuration
  resources:
  - mutatingwebhookconfigurations
  verbs:
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - config.gatekeeper.sh
  resources:
  - configs
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - config.gatekeeper.sh
  resources:
  - configs/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - constraints.gatekeeper.sh
  resources:
  - '*'
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - externaldata.gatekeeper.sh
  resources:
  - providers
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - mutations.gatekeeper.sh
  resources:
  - '*'
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - policy
  resourceNames:
  - gatekeeper-admin
  resources:
  - podsecuritypolicies
  verbs:
  - use
- apiGroups:
  - status.gatekeeper.sh
  resources:
  - '*'
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - templates.gatekeeper.sh
  resources:
  - constrainttemplates
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - templates.gatekeeper.sh
  resources:
  - constrainttemplates/finalizers
  verbs:
  - delete
  - get
  - patch
  - update
- apiGroups:
  - templates.gatekeeper.sh
  resources:
  - constrainttemplates/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - admissionregistration.k8s.io
  resourceNames:
  - gatekeeper-validating-webhook-configuration
  resources:
  - validatingwebhookconfigurations
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
`)

func configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml, nil
}

func configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: gatekeeper-manager-role
subjects:
- kind: ServiceAccount
  name: gatekeeper-admin
  namespace: gatekeeper-system
`)

func configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml, nil
}

func configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-manager-role
  namespace: gatekeeper-system
rules:
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - security.openshift.io
  resourceNames:
  - anyuid
  resources:
  - securitycontextconstraints
  verbs:
  - use
`)

func configGatekeeperRenderedRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml, nil
}

func configGatekeeperRenderedRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-manager-rolebinding
  namespace: gatekeeper-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: gatekeeper-manager-role
subjects:
- kind: ServiceAccount
  name: gatekeeper-admin
  namespace: gatekeeper-system
`)

func configGatekeeperRenderedRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml, nil
}

func configGatekeeperRenderedRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedV1_namespace_gatekeeperSystemYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  labels:
    admission.gatekeeper.sh/ignore: no-self-managing
    control-plane: controller-manager
    gatekeeper.sh/system: "yes"
    pod-security.kubernetes.io/audit: restricted
    pod-security.kubernetes.io/audit-version: latest
    pod-security.kubernetes.io/enforce: restricted
    pod-security.kubernetes.io/enforce-version: v1.24
    pod-security.kubernetes.io/warn: restricted
    pod-security.kubernetes.io/warn-version: latest
  name: gatekeeper-system
`)

func configGatekeeperRenderedV1_namespace_gatekeeperSystemYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedV1_namespace_gatekeeperSystemYaml, nil
}

func configGatekeeperRenderedV1_namespace_gatekeeperSystemYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedV1_namespace_gatekeeperSystemYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/v1_namespace_gatekeeper-system.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedV1_resourcequota_gatekeeperCriticalPodsYaml = []byte(`apiVersion: v1
kind: ResourceQuota
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-critical-pods
  namespace: gatekeeper-system
spec:
  hard:
    pods: "100"
  scopeSelector:
    matchExpressions:
    - operator: In
      scopeName: PriorityClass
      values:
      - system-cluster-critical
`)

func configGatekeeperRenderedV1_resourcequota_gatekeeperCriticalPodsYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedV1_resourcequota_gatekeeperCriticalPodsYaml, nil
}

func configGatekeeperRenderedV1_resourcequota_gatekeeperCriticalPodsYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedV1_resourcequota_gatekeeperCriticalPodsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/v1_resourcequota_gatekeeper-critical-pods.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedV1_secret_gatekeeperWebhookServerCertYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-webhook-server-cert
  namespace: gatekeeper-system
`)

func configGatekeeperRenderedV1_secret_gatekeeperWebhookServerCertYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedV1_secret_gatekeeperWebhookServerCertYaml, nil
}

func configGatekeeperRenderedV1_secret_gatekeeperWebhookServerCertYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedV1_secret_gatekeeperWebhookServerCertYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/v1_secret_gatekeeper-webhook-server-cert.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedV1_service_gatekeeperWebhookServiceYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-webhook-service
  namespace: gatekeeper-system
spec:
  ports:
  - name: https-webhook-server
    port: 443
    targetPort: webhook-server
  selector:
    control-plane: controller-manager
    gatekeeper.sh/operation: webhook
    gatekeeper.sh/system: "yes"
`)

func configGatekeeperRenderedV1_service_gatekeeperWebhookServiceYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedV1_service_gatekeeperWebhookServiceYaml, nil
}

func configGatekeeperRenderedV1_service_gatekeeperWebhookServiceYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedV1_service_gatekeeperWebhookServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/v1_service_gatekeeper-webhook-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRenderedV1_serviceaccount_gatekeeperAdminYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-admin
  namespace: gatekeeper-system
`)

func configGatekeeperRenderedV1_serviceaccount_gatekeeperAdminYamlBytes() ([]byte, error) {
	return _configGatekeeperRenderedV1_serviceaccount_gatekeeperAdminYaml, nil
}

func configGatekeeperRenderedV1_serviceaccount_gatekeeperAdminYaml() (*asset, error) {
	bytes, err := configGatekeeperRenderedV1_serviceaccount_gatekeeperAdminYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper-rendered/v1_serviceaccount_gatekeeper-admin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config/gatekeeper-rendered/admissionregistration.k8s.io_v1_mutatingwebhookconfiguration_gatekeeper-mutating-webhook-configuration.yaml":     configGatekeeperRenderedAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml,
	"config/gatekeeper-rendered/admissionregistration.k8s.io_v1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml": configGatekeeperRenderedAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_assign.mutations.gatekeeper.sh.yaml":                            configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_assignmetadata.mutations.gatekeeper.sh.yaml":                    configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_configs.config.gatekeeper.sh.yaml":                              configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml":                configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml":        configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml":               configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_expansiontemplate.expansion.gatekeeper.sh.yaml":                 configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_expansiontemplateExpansionGatekeeperShYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_modifyset.mutations.gatekeeper.sh.yaml":                         configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_modifysetMutationsGatekeeperShYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_mutatorpodstatuses.status.gatekeeper.sh.yaml":                   configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml,
	"config/gatekeeper-rendered/apiextensions.k8s.io_v1_customresourcedefinition_providers.externaldata.gatekeeper.sh.yaml":                      configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_providersExternaldataGatekeeperShYaml,
	"config/gatekeeper-rendered/apps_v1_deployment_gatekeeper-audit.yaml":                                                                        configGatekeeperRenderedApps_v1_deployment_gatekeeperAuditYaml,
	"config/gatekeeper-rendered/apps_v1_deployment_gatekeeper-controller-manager.yaml":                                                           configGatekeeperRenderedApps_v1_deployment_gatekeeperControllerManagerYaml,
	"config/gatekeeper-rendered/policy_v1_poddisruptionbudget_gatekeeper-controller-manager.yaml":                                                configGatekeeperRenderedPolicy_v1_poddisruptionbudget_gatekeeperControllerManagerYaml,
	"config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml":                                           configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml,
	"config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml":                             configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml,
	"config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml":                                                  configGatekeeperRenderedRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml,
	"config/gatekeeper-rendered/rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml":                                    configGatekeeperRenderedRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml,
	"config/gatekeeper-rendered/v1_namespace_gatekeeper-system.yaml":                                                                             configGatekeeperRenderedV1_namespace_gatekeeperSystemYaml,
	"config/gatekeeper-rendered/v1_resourcequota_gatekeeper-critical-pods.yaml":                                                                  configGatekeeperRenderedV1_resourcequota_gatekeeperCriticalPodsYaml,
	"config/gatekeeper-rendered/v1_secret_gatekeeper-webhook-server-cert.yaml":                                                                   configGatekeeperRenderedV1_secret_gatekeeperWebhookServerCertYaml,
	"config/gatekeeper-rendered/v1_service_gatekeeper-webhook-service.yaml":                                                                      configGatekeeperRenderedV1_service_gatekeeperWebhookServiceYaml,
	"config/gatekeeper-rendered/v1_serviceaccount_gatekeeper-admin.yaml":                                                                         configGatekeeperRenderedV1_serviceaccount_gatekeeperAdminYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"config": {nil, map[string]*bintree{
		"gatekeeper-rendered": {nil, map[string]*bintree{
			"admissionregistration.k8s.io_v1_mutatingwebhookconfiguration_gatekeeper-mutating-webhook-configuration.yaml":     {configGatekeeperRenderedAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml, map[string]*bintree{}},
			"admissionregistration.k8s.io_v1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml": {configGatekeeperRenderedAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_assign.mutations.gatekeeper.sh.yaml":                            {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_assignmetadata.mutations.gatekeeper.sh.yaml":                    {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_configs.config.gatekeeper.sh.yaml":                              {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml":                {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml":        {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml":               {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_expansiontemplate.expansion.gatekeeper.sh.yaml":                 {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_expansiontemplateExpansionGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_modifyset.mutations.gatekeeper.sh.yaml":                         {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_modifysetMutationsGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_mutatorpodstatuses.status.gatekeeper.sh.yaml":                   {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_providers.externaldata.gatekeeper.sh.yaml":                      {configGatekeeperRenderedApiextensionsK8sIo_v1_customresourcedefinition_providersExternaldataGatekeeperShYaml, map[string]*bintree{}},
			"apps_v1_deployment_gatekeeper-audit.yaml":                                                                        {configGatekeeperRenderedApps_v1_deployment_gatekeeperAuditYaml, map[string]*bintree{}},
			"apps_v1_deployment_gatekeeper-controller-manager.yaml":                                                           {configGatekeeperRenderedApps_v1_deployment_gatekeeperControllerManagerYaml, map[string]*bintree{}},
			"policy_v1_poddisruptionbudget_gatekeeper-controller-manager.yaml":                                                {configGatekeeperRenderedPolicy_v1_poddisruptionbudget_gatekeeperControllerManagerYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml":                                           {configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml":                             {configGatekeeperRenderedRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml":                                                  {configGatekeeperRenderedRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml":                                    {configGatekeeperRenderedRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml, map[string]*bintree{}},
			"v1_namespace_gatekeeper-system.yaml":                                                                             {configGatekeeperRenderedV1_namespace_gatekeeperSystemYaml, map[string]*bintree{}},
			"v1_resourcequota_gatekeeper-critical-pods.yaml":                                                                  {configGatekeeperRenderedV1_resourcequota_gatekeeperCriticalPodsYaml, map[string]*bintree{}},
			"v1_secret_gatekeeper-webhook-server-cert.yaml":                                                                   {configGatekeeperRenderedV1_secret_gatekeeperWebhookServerCertYaml, map[string]*bintree{}},
			"v1_service_gatekeeper-webhook-service.yaml":                                                                      {configGatekeeperRenderedV1_service_gatekeeperWebhookServiceYaml, map[string]*bintree{}},
			"v1_serviceaccount_gatekeeper-admin.yaml":                                                                         {configGatekeeperRenderedV1_serviceaccount_gatekeeperAdminYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
