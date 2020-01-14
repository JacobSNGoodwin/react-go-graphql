import React from "react";

import styles from "./Spinner.module.scss";

interface ISpinnerProps {
  radius?: number | undefined;
  color?: string | undefined;
}

const Spinner: React.FC<ISpinnerProps> = props => {
  const thickness: number = props.radius ? props.radius / 10 : 16;
  const radius: number = props.radius ? props.radius : 160;
  const color: string = props.color ? props.color : "#f44800";

  const spinnerStyle: React.CSSProperties = {
    border: `${thickness}px solid #f3f3f3`,
    borderTop: `${thickness}px solid ${color}`,
    borderRadius: "50%",
    width: radius,
    height: radius,
    animation: `${styles.spin} 2s linear infinite`
  };

  return <div style={spinnerStyle}></div>;
};

export default Spinner;
