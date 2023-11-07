import { ViewPage } from "../home/ViewPage";

export const Page1 = () => {
  return (
    <ViewPage roles={["admin", "user", "guest"]}>
      <h1>Page1</h1>
    </ViewPage>
  );
};
