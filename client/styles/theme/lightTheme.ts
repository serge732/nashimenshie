import { createTheme } from '@mui/material/styles';
import { deepmerge } from '@mui/utils';

import components from './components';

export default createTheme(
    deepmerge(
        {
            palette: {
                mode: 'light',
            },
        },
        components
    )
);
