import { styled } from '@mui/material/styles';
import { ListItemIcon, ListItemButton } from '@mui/material';

// ----------------------------------------------------------------------

export interface StyledNavItemProps {
    children: JSX.Element | JSX.Element[]
}

export const StyledNavItem = styled((props: any) => <ListItemButton disableGutters {...props} />)(({ theme }) => ({
    ...theme.typography.body2,
    height: 48,
    position: 'relative',
    textTransform: 'capitalize',
    color: theme.palette.text.secondary,
    borderRadius: theme.shape.borderRadius,
}));

export const StyledNavItemIcon = styled(ListItemIcon)(({ theme }) => ({
    width: 22,
    height: 22,
    color: 'inherit',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    marginRight: theme.spacing(1),
}));

