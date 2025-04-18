import { IDeployment } from "npm:kubernetes-models/apps/v1";

export function createDeployment(name: string): IDeployment {
  return {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
      name: name,
    },
    spec: {
      replicas: 1,
      selector: {
        matchLabels: {
          app: name,
        },
      },
      template: {
        metadata: {
          labels: {
            app: name,
          },
        },
        spec: {
          containers: [
            {
              name: name,
              image: "nginx:latest",
              imagePullPolicy: "IfNotPresent",
            },
          ],
          affinity: {
            nodeAffinity: {
              preferredDuringSchedulingIgnoredDuringExecution: [
                {
                  weight: 1,
                  preference: {
                    matchExpressions: [
                      {
                        key: "node-role.kubernetes.io/worker",
                        operator: "In",
                        values: ["true"],
                      },
                    ],
                  },
                },
              ],
            },
          },
        },
      },
    },
  };
}
