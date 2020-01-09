interface IError {
  message: string;
  type: string | undefined;
}

interface ErrorProps {
  messages?: string[];
  includeLogin?: boolean;
}