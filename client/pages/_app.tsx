import { CacheProvider, EmotionCache } from '@emotion/react';
import { CssBaseline } from '@mui/material';
import { ThemeProvider } from '@mui/material/styles';
import type { AppProps } from 'next/app';

import MainLayout from 'layouts/MainLayout';

import '../styles/global.css';
import lightTheme from '../styles/theme/lightTheme';
import createEmotionCache from '../utils/createEmotionCache';

const clientSideEmotionCache = createEmotionCache();

interface MyAppProps extends AppProps {
    emotionCache?: EmotionCache;
}

export default function MyApp({ Component, emotionCache = clientSideEmotionCache, pageProps }: MyAppProps) {
    return (
        <CacheProvider value={emotionCache}>
            <ThemeProvider theme={lightTheme}>
                <MainLayout>
                    <CssBaseline />
                    <Component {...pageProps} />
                </MainLayout>
            </ThemeProvider>
        </CacheProvider>
    );
}
