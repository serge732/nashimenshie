import axios, { AxiosResponse } from 'axios';

import { moyskladInstance } from '../instance';
import { RequestResult } from '../types';
import { handleError } from './../utils/handleError';
import { Product } from './product';

export type AssortmentRequestInput = {
    limit: number;
    offset: number;
};

export type Pagination = {
    pageCount: number;
    currentPage: number;
};

export type Assortment = {
    products: Product[] | null;
    pagination: Pagination;
};

export type AssortmentRequestResult = RequestResult<Assortment>;

export const fetchAssortment = async (variables: AssortmentRequestInput): Promise<AssortmentRequestResult> => {
    try {
        const response = await moyskladInstance.put<{}, AxiosResponse<AssortmentRequestResult>, AssortmentRequestInput>(
            'assortment',
            {
                ...variables,
            }
        );

        if (response.data.error) {
            throw new Error(response.data.error);
        }

        return response.data;
    } catch (error) {
        return handleError(error);
    }
};
