import PropTypes from 'prop-types'
import { NavLink as RouterLink } from 'react-router-dom'
import { Box, List, ListItemText } from '@mui/material'
import { StyledNavItem, StyledNavItemIcon } from './styles'

interface NavItemProps {
    item: {
        title: string
        path: string
        icon?: React.ReactNode
        info?: React.ReactNode
    }
}

NavItem.propTypes = {
    item: PropTypes.object
}

function NavItem({ item }: NavItemProps) {
    const { title, path, icon } = item

    return (
        <StyledNavItem
            component={RouterLink}
            to={path}
            sx={{
                '&.active': {
                    color: 'text.primary',
                    bgcolor: 'action.selected',
                    fontWeight: 'fontWeightBold'
                }
            }}>
            <StyledNavItemIcon>{icon}</StyledNavItemIcon>
            <ListItemText disableTypography primary={title} />
        </StyledNavItem>
    )
}

interface NavSectionProps {
    data?: Array<{
        title: string
        path: string
        icon?: React.ReactNode
        info?: React.ReactNode
    }>
}

NavSection.propTypes = {
    data: PropTypes.array
}

export default function NavSection({ data = [], ...other }: NavSectionProps) {
    return (
        <Box {...other}>
            <List disablePadding sx={{ p: 1 }}>
                {data.map((item) => (
                    <NavItem key={item.title} item={item} />
                ))}
            </List>
        </Box>
    )
}
