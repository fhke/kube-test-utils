package health

import v1 "k8s.io/api/core/v1"

func PodReady(pod v1.Pod) bool {
	return podConditionsReady(pod.Status.Conditions)
}

func podConditionsReady(conditions []v1.PodCondition) bool {
	for _, condition := range conditions {
		if condition.Type == v1.ContainersReady {
			if condition.Status == v1.ConditionTrue {
				// all containers are ready
				return true
			} else {
				return false
			}
		}
	}
	return false
}
