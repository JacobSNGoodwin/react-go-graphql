import React from "react";

interface EditProductProps {
  show: boolean;
  initProduct?: IProduct;
  editing: boolean;
  close: () => void;
  // editSelectedUser: (product: IProduct) => void;
}

const EditProduct: React.FC<EditProductProps> = props => {
  const editMode = props.initProduct ? true : false;
  const title = editMode ? "Edit Product" : "Create Product";
  return (
    <div
      className={"has-text-centered modal" + (props.show ? " is-active" : "")}
    >
      <div className="modal-background"></div>
      <div className="modal-card">
        <header className="modal-card-head">
          <p className="modal-card-title">{title}</p>
          <button
            onClick={props.close}
            className="delete"
            aria-label="close"
          ></button>
        </header>
        <section className="modal-card-body">Edit Product</section>
        <footer className="modal-card-foot"></footer>
      </div>
    </div>
  );
};

export default EditProduct;
