import { moyskladInstance } from 'api/instance';
import { RequestResult } from 'api/types';

export type ProductQueryVariables = {
    id: string;
};

export type Product = {
    id: string;
    name: string;
    description: string;
    imagesUrls: string;
    price: number;
};

export type ProductRequestResult = RequestResult<Product>;

export const fetchProduct = async (variables: ProductQueryVariables): Promise<ProductRequestResult> => {
    return await moyskladInstance.put<{}, ProductRequestResult, ProductQueryVariables>('product', {
        ...variables,
    });
};
