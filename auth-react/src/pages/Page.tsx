import { ViewPage } from "./ViewPage";

export const Page = () => {
  return (
    <ViewPage roles={["admin", "user", "guest"]}>
      <h1>Page1</h1>
    </ViewPage>
  );
};
