import { createTheme } from '@mui/material/styles';

export const MuiSkeleton = createTheme({
    components: {
        MuiSkeleton: {
            defaultProps: {
                animation: 'wave',
            },
        },
    },
});
