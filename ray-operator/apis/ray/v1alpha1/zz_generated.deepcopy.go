//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalerOptions) DeepCopyInto(out *AutoscalerOptions) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.IdleTimeoutSeconds != nil {
		in, out := &in.IdleTimeoutSeconds, &out.IdleTimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.UpscalingMode != nil {
		in, out := &in.UpscalingMode, &out.UpscalingMode
		*out = new(UpscalingMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalerOptions.
func (in *AutoscalerOptions) DeepCopy() *AutoscalerOptions {
	if in == nil {
		return nil
	}
	out := new(AutoscalerOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashBoardStatus) DeepCopyInto(out *DashBoardStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.HealthLastUpdateTime.DeepCopyInto(&out.HealthLastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashBoardStatus.
func (in *DashBoardStatus) DeepCopy() *DashBoardStatus {
	if in == nil {
		return nil
	}
	out := new(DashBoardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeadGroupSpec) DeepCopyInto(out *HeadGroupSpec) {
	*out = *in
	if in.EnableIngress != nil {
		in, out := &in.EnableIngress, &out.EnableIngress
		*out = new(bool)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.RayStartParams != nil {
		in, out := &in.RayStartParams, &out.RayStartParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeadGroupSpec.
func (in *HeadGroupSpec) DeepCopy() *HeadGroupSpec {
	if in == nil {
		return nil
	}
	out := new(HeadGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayActorOptionSpec) DeepCopyInto(out *RayActorOptionSpec) {
	*out = *in
	if in.RuntimeEnv != nil {
		in, out := &in.RuntimeEnv, &out.RuntimeEnv
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.NumCpus != nil {
		in, out := &in.NumCpus, &out.NumCpus
		*out = new(float64)
		**out = **in
	}
	if in.NumGpus != nil {
		in, out := &in.NumGpus, &out.NumGpus
		*out = new(float64)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(int32)
		**out = **in
	}
	if in.ObjectStoreMemory != nil {
		in, out := &in.ObjectStoreMemory, &out.ObjectStoreMemory
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayActorOptionSpec.
func (in *RayActorOptionSpec) DeepCopy() *RayActorOptionSpec {
	if in == nil {
		return nil
	}
	out := new(RayActorOptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayCluster) DeepCopyInto(out *RayCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayCluster.
func (in *RayCluster) DeepCopy() *RayCluster {
	if in == nil {
		return nil
	}
	out := new(RayCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterList) DeepCopyInto(out *RayClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RayCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterList.
func (in *RayClusterList) DeepCopy() *RayClusterList {
	if in == nil {
		return nil
	}
	out := new(RayClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterSpec) DeepCopyInto(out *RayClusterSpec) {
	*out = *in
	in.HeadGroupSpec.DeepCopyInto(&out.HeadGroupSpec)
	if in.WorkerGroupSpecs != nil {
		in, out := &in.WorkerGroupSpecs, &out.WorkerGroupSpecs
		*out = make([]WorkerGroupSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableInTreeAutoscaling != nil {
		in, out := &in.EnableInTreeAutoscaling, &out.EnableInTreeAutoscaling
		*out = new(bool)
		**out = **in
	}
	if in.AutoscalerOptions != nil {
		in, out := &in.AutoscalerOptions, &out.AutoscalerOptions
		*out = new(AutoscalerOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterSpec.
func (in *RayClusterSpec) DeepCopy() *RayClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RayClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterStatus) DeepCopyInto(out *RayClusterStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterStatus.
func (in *RayClusterStatus) DeepCopy() *RayClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RayClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayService) DeepCopyInto(out *RayService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayService.
func (in *RayService) DeepCopy() *RayService {
	if in == nil {
		return nil
	}
	out := new(RayService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayServiceList) DeepCopyInto(out *RayServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RayService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayServiceList.
func (in *RayServiceList) DeepCopy() *RayServiceList {
	if in == nil {
		return nil
	}
	out := new(RayServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayServiceSpec) DeepCopyInto(out *RayServiceSpec) {
	*out = *in
	if in.ServeConfigSpecs != nil {
		in, out := &in.ServeConfigSpecs, &out.ServeConfigSpecs
		*out = make([]ServeConfigSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.RayClusterSpec.DeepCopyInto(&out.RayClusterSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayServiceSpec.
func (in *RayServiceSpec) DeepCopy() *RayServiceSpec {
	if in == nil {
		return nil
	}
	out := new(RayServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayServiceStatus) DeepCopyInto(out *RayServiceStatus) {
	*out = *in
	if in.ServeStatuses != nil {
		in, out := &in.ServeStatuses, &out.ServeStatuses
		*out = make([]ServeDeploymentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.DashBoardStatus.DeepCopyInto(&out.DashBoardStatus)
	in.RayClusterStatus.DeepCopyInto(&out.RayClusterStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayServiceStatus.
func (in *RayServiceStatus) DeepCopy() *RayServiceStatus {
	if in == nil {
		return nil
	}
	out := new(RayServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleStrategy) DeepCopyInto(out *ScaleStrategy) {
	*out = *in
	if in.WorkersToDelete != nil {
		in, out := &in.WorkersToDelete, &out.WorkersToDelete
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleStrategy.
func (in *ScaleStrategy) DeepCopy() *ScaleStrategy {
	if in == nil {
		return nil
	}
	out := new(ScaleStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServeConfigSpec) DeepCopyInto(out *ServeConfigSpec) {
	*out = *in
	if in.InitArgs != nil {
		in, out := &in.InitArgs, &out.InitArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InitKwargs != nil {
		in, out := &in.InitKwargs, &out.InitKwargs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NumReplicas != nil {
		in, out := &in.NumReplicas, &out.NumReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxConcurrentQueries != nil {
		in, out := &in.MaxConcurrentQueries, &out.MaxConcurrentQueries
		*out = new(int32)
		**out = **in
	}
	if in.UserConfig != nil {
		in, out := &in.UserConfig, &out.UserConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AutoscalingConfig != nil {
		in, out := &in.AutoscalingConfig, &out.AutoscalingConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.GracefulShutdownWaitLoopS != nil {
		in, out := &in.GracefulShutdownWaitLoopS, &out.GracefulShutdownWaitLoopS
		*out = new(int32)
		**out = **in
	}
	if in.GracefulShutdownTimeoutS != nil {
		in, out := &in.GracefulShutdownTimeoutS, &out.GracefulShutdownTimeoutS
		*out = new(int32)
		**out = **in
	}
	if in.HealthCheckPeriodS != nil {
		in, out := &in.HealthCheckPeriodS, &out.HealthCheckPeriodS
		*out = new(int32)
		**out = **in
	}
	if in.HealthCheckTimeoutS != nil {
		in, out := &in.HealthCheckTimeoutS, &out.HealthCheckTimeoutS
		*out = new(int32)
		**out = **in
	}
	in.RayActorOptions.DeepCopyInto(&out.RayActorOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServeConfigSpec.
func (in *ServeConfigSpec) DeepCopy() *ServeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ServeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServeDeploymentStatus) DeepCopyInto(out *ServeDeploymentStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.HealthLastUpdateTime.DeepCopyInto(&out.HealthLastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServeDeploymentStatus.
func (in *ServeDeploymentStatus) DeepCopy() *ServeDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ServeDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServeDeploymentStatuses) DeepCopyInto(out *ServeDeploymentStatuses) {
	*out = *in
	if in.Statuses != nil {
		in, out := &in.Statuses, &out.Statuses
		*out = make([]ServeDeploymentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServeDeploymentStatuses.
func (in *ServeDeploymentStatuses) DeepCopy() *ServeDeploymentStatuses {
	if in == nil {
		return nil
	}
	out := new(ServeDeploymentStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerGroupSpec) DeepCopyInto(out *WorkerGroupSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	if in.RayStartParams != nil {
		in, out := &in.RayStartParams, &out.RayStartParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	in.ScaleStrategy.DeepCopyInto(&out.ScaleStrategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerGroupSpec.
func (in *WorkerGroupSpec) DeepCopy() *WorkerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerGroupSpec)
	in.DeepCopyInto(out)
	return out
}
