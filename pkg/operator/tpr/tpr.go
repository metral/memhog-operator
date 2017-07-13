package tpr

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

func CreateTPR(clientSet kubernetes.Interface, name, version, desc string) (*v1beta1.ThirdPartyResource, error) {
	// initialize third party resource if it does not exist
	tpr, err := clientSet.Extensions().ThirdPartyResources().Get(name, metav1.GetOptions{ResourceVersion: "0"})
	if err != nil {
		if errors.IsNotFound(err) {
			tpr := &v1beta1.ThirdPartyResource{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
				},
				Versions: []v1beta1.APIVersion{
					{Name: version},
				},
				Description: desc,
			}

			result, err := clientSet.Extensions().ThirdPartyResources().Create(tpr)
			if err != nil {
				return nil, err
			}
			fmt.Printf("CREATED: %#v\nFROM: %#v\n", result, tpr)
		} else {
			return nil, err
		}
	} else {
		fmt.Printf("SKIPPING: already exists %#v\n", tpr)
	}

	return tpr, nil
}
