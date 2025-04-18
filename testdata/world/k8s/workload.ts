import { createDeployment } from "./util.ts";

await new Promise((resolve) => setTimeout(resolve, 1000));

export default [createDeployment("workload-1")];
