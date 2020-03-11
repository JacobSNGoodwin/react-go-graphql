import React from "react";
import { useQuery, useMutation } from "@apollo/react-hooks";
import { ApolloError } from "apollo-boost";

import { GET_PRODUCTS } from "../../gql/queries";
import { CREATE_PRODUCT } from "../../gql/mutations";
import Spinner from "../ui/Spinner";
import ErrorsList from "../ErrorsList";
import Product from "../Product";
import EditProduct from "../EditProduct";

const Products: React.FC = props => {
  const [createActive, setCreateActive] = React.useState<boolean>(false);
  const [apolloError, setApolloError] = React.useState<ApolloError | undefined>(
    undefined
  );

  /*
   * Create User
   */
  const [createProductMutation, { loading: creatingProduct }] = useMutation<
    { createdUser: IProduct },
    { user: IProduct }
  >(CREATE_PRODUCT, {
    refetchQueries: [
      {
        query: GET_PRODUCTS,
        variables: {
          limit: 10
        }
      }
    ],
    onError: error => {
      setApolloError(error);
    }
  });

  /*
   * Read (get) Users
   */
  const { loading: loadingProducts, data } = useQuery<
    IProductData,
    IProductVars
  >(GET_PRODUCTS, {
    variables: {
      limit: 10
    },
    onError: error => {
      setApolloError(error);
    }
  });

  if (loadingProducts) {
    return (
      <div className="container">
        <div className="columns is-centered">
          <Spinner radius={50} />
        </div>
      </div>
    );
  }

  if (apolloError) {
    return <ErrorsList error={apolloError} />;
  } else {
    const productList =
      data &&
      data.products.map(product => {
        return (
          <div key={product.id} className="column is-half">
            <Product product={product} />
          </div>
        );
      });
    return (
      <>
        <h1 className="title is-1 has-text-centered">Products</h1>
        <div className="container">
          <div className="columns is-centered">
            <button
              className="button is-warning"
              onClick={() => {
                setCreateActive(true);
              }}
            >
              Create
            </button>
          </div>
          <div className="columns is-vcentered is-multiline">{productList}</div>
        </div>

        <EditProduct
          show={createActive}
          editing={false}
          close={() => setCreateActive(false)}
        />
      </>
    );
  }
};

export default Products;
