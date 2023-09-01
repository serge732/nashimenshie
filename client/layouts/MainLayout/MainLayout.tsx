import ShoppingCartOutlinedIcon from '@mui/icons-material/ShoppingCartOutlined';
import Badge from '@mui/material/Badge';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';

export default function MainLayout({ children }: React.PropsWithChildren<{}>) {
    return (
        <>
            <header>
                <Container>
                    <Box display="flex" justifyContent="space-between">
                        <Typography>Наши Меньшие</Typography>
                        <IconButton aria-label="shopping cart" size="large" color="primary">
                            <Badge badgeContent={4} color="primary">
                                <ShoppingCartOutlinedIcon fontSize="inherit" />
                            </Badge>
                        </IconButton>
                    </Box>
                </Container>
            </header>
            <main>
                <Container>{children}</Container>
            </main>
        </>
    );
}
