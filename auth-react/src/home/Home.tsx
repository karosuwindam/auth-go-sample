import { ViewPage } from "../home/ViewPage";

export const Home = () => {
  return (
    <ViewPage roles={["admin", "user"]}>
      <h1>Home</h1>
    </ViewPage>
  );
};
