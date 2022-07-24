interface SquareProps {
  value?: string | null;
  onClick: () => void;
}

const Square = (props: SquareProps) => {
  return (
    <button className="square" onClick={props.onClick}>
      {props.value}
    </button>
  );
};

export default Square;
