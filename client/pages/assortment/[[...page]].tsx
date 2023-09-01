import Grid from '@mui/material/Grid';
import Pagination from '@mui/material/Pagination';
import { GetStaticPaths, GetStaticProps } from 'next';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';

import { Assortment, fetchAssortment } from 'api/requests/assortment';

import ProductCard, { ProductCardSkeleton } from 'components/ProductCard';

export type TAssortmentPageProps = {
    data: Assortment | null;
    error: string | null;
};

export default function AssortmentPage({ data, error }: TAssortmentPageProps) {
    const [isFetching, setIsFetching] = useState(false);

    const router = useRouter();

    const handleOnPageChange = (event: React.ChangeEvent<unknown>, value: number) => {
        router.push(`/assortment/${value}`);
        setIsFetching(true);
    };

    useEffect(() => setIsFetching(false), [data]);

    return (
        <>
            {!!data && !!data.products && (
                <>
                    <Grid container columns={12} spacing="24px">
                        {data.products.map((product) => (
                            <React.Fragment key={`ProductCardList__card-${product.id}`}>
                                <Grid item xs={12} sm={6} md={4}>
                                    {isFetching ? (
                                        <ProductCardSkeleton />
                                    ) : (
                                        <ProductCard
                                            id={product.id}
                                            name={product.name}
                                            imagesUrls={product.imagesUrls}
                                            price={product.price}
                                        />
                                    )}
                                </Grid>
                            </React.Fragment>
                        ))}
                    </Grid>
                    <Pagination
                        count={data.pagination.pageCount}
                        page={data.pagination.currentPage ?? 1}
                        onChange={handleOnPageChange}
                    />
                </>
            )}
        </>
    );
}

export const getStaticPaths: GetStaticPaths = async () => {
    return {
        paths: [{ params: { page: [] } }],
        fallback: true,
    };
};

export const getStaticProps: GetStaticProps<TAssortmentPageProps> = async ({ params }) => {
    const limit = 18;

    const offset = (() => {
        if (params && params.page) {
            return limit * Number(params.page[0]) - limit;
        }

        return 0;
    })();

    const { data, error } = await fetchAssortment({ limit, offset });

    return {
        props: {
            data: data ?? null,
            error: error ?? null,
        },
    };
};
