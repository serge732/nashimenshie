import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Skeleton from '@mui/material/Skeleton';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Image from 'next/image';

import { Product } from 'api/requests/product';

type TProductCardProps = Omit<Product, 'description'>;

export default function ProductCard({ id, imagesUrls, price, name }: TProductCardProps) {
    return (
        <Stack justifyContent="space-between" height="100%" spacing="28px">
            <Box>
                <Box textAlign="center" sx={{ position: 'relative', width: '100%', height: 200 }}>
                    <Image src={imagesUrls[0]} layout="fill" objectFit="contain" />
                </Box>
                <Typography>{price}</Typography>
                <Typography>{name}</Typography>
            </Box>
            <Button id={id} variant="outlined">
                В корзину
            </Button>
        </Stack>
    );
}

export function ProductCardSkeleton() {
    return (
        <Box>
            <Skeleton variant="rectangular" width="100%" height={200} />
            <Skeleton variant="text" width="21%" />
            <Skeleton variant="text" />
            <Skeleton variant="text" />
        </Box>
    );
}
