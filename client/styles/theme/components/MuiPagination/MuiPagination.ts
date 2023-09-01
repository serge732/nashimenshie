import { createTheme } from '@mui/material/styles';

export const MuiPagination = createTheme({
    components: {
        MuiPagination: {
            defaultProps: {
                shape: 'rounded',
                siblingCount: 3,
            },
        },
    },
});
