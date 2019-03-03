import { createMuiTheme } from '@material-ui/core';
import grey from '@material-ui/core/colors/grey';
import { primary } from '../styles/primary';

export const theme = createMuiTheme({
  palette: {
    type: 'dark',
    primary,
    secondary: grey,
  },
});
