import React from "react";

import placeholder from "../images/placeholder.png";

interface ProductProps {
  product: IProduct;
}

const Product: React.FC<ProductProps> = props => {
  return (
    <div className="card" key={props.product.id}>
      <div className="card-content has-text-centered">
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            padding: "1.5em 1em"
          }}
        >
          <figure className="image is-96x96">
            {!props.product.imageUri || props.product.imageUri === "" ? (
              <img className="is-rounded" src={placeholder} alt="No profile" />
            ) : (
              <img
                className="is-rounded"
                src={props.product.imageUri}
                alt="Product"
              />
            )}
          </figure>
        </div>

        <p className="title is-4">{props.product.name}</p>
        <p className="subtitle is-6">${props.product.price / 100}</p>
        <br />
        <div className="is-size-5 has-text-weight-bold">Categories</div>
      </div>
    </div>
  );
};

export default Product;
