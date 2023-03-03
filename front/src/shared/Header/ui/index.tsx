import { Template } from "./template";

interface Props {
  children: React.ReactNode;
}

export const Header = (props: Props) => {
  return <Template {...props} />;
};
