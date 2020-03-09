import React from "react";
import { useQuery, useMutation } from "@apollo/react-hooks";
import { ApolloError } from "apollo-boost";

import { GET_PRODUCTS } from "../../gql/queries";
import Spinner from "../ui/Spinner";
import ErrorsList from "../ErrorsList";

const Products: React.FC = props => {
  const [createActive, setCreateActive] = React.useState<boolean>(false);
  const [apolloError, setApolloError] = React.useState<ApolloError | undefined>(
    undefined
  );

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
    console.log(data?.products);
    return (
      <>
        <h1 className="title is-1 has-text-centered">Products</h1>
        <div className="container">
          Products List
          {/* <div className="columns is-centered">
                <button
                  className="button is-warning"
                  onClick={() => {
                    setCreateActive(true);
                  }}
                >
                  Create
                </button>
              </div>
              <div className="columns is-vcentered is-multiline">{userList}</div> */}
        </div>
      </>
    );
  }
};

export default Products;
