import SvgIcon, {SvgIconProps} from '@mui/material/SvgIcon';
import {alpha, styled} from '@mui/material/styles';
import TreeView from '@mui/lab/TreeView';
import TreeItem, {TreeItemProps, treeItemClasses} from '@mui/lab/TreeItem';
import Collapse from '@mui/material/Collapse';
import {TransitionProps} from '@mui/material/transitions';
import {Article} from "@mui/icons-material";
import {NavigationItem} from "./navigation.ts";

function MinusSquare(props: SvgIconProps) {
    return (
        <SvgIcon fontSize="inherit" style={{width: 14, height: 14}} {...props}>
            {/* tslint:disable-next-line: max-line-length */}
            <path
                d="M22.047 22.074v0 0-20.147 0h-20.12v0 20.147 0h20.12zM22.047 24h-20.12q-.803 0-1.365-.562t-.562-1.365v-20.147q0-.776.562-1.351t1.365-.575h20.147q.776 0 1.351.575t.575 1.351v20.147q0 .803-.575 1.365t-1.378.562v0zM17.873 11.023h-11.826q-.375 0-.669.281t-.294.682v0q0 .401.294 .682t.669.281h11.826q.375 0 .669-.281t.294-.682v0q0-.401-.294-.682t-.669-.281z"/>
        </SvgIcon>
    );
}

function PlusSquare(props: SvgIconProps) {
    return (
        <SvgIcon fontSize="inherit" style={{width: 14, height: 14}} {...props}>
            {/* tslint:disable-next-line: max-line-length */}
            <path
                d="M22.047 22.074v0 0-20.147 0h-20.12v0 20.147 0h20.12zM22.047 24h-20.12q-.803 0-1.365-.562t-.562-1.365v-20.147q0-.776.562-1.351t1.365-.575h20.147q.776 0 1.351.575t.575 1.351v20.147q0 .803-.575 1.365t-1.378.562v0zM17.873 12.977h-4.923v4.896q0 .401-.281.682t-.682.281v0q-.375 0-.669-.281t-.294-.682v-4.896h-4.923q-.401 0-.682-.294t-.281-.669v0q0-.401.281-.682t.682-.281h4.923v-4.896q0-.401.294-.682t.669-.281v0q.401 0 .682.281t.281.682v4.896h4.923q.401 0 .682.281t.281.682v0q0 .375-.281.669t-.682.294z"/>
        </SvgIcon>
    );
}

function TransitionComponent(props: TransitionProps) {

    return (
        <Collapse {...props} />
    );
}

const StyledTreeItem = styled((props: TreeItemProps) => (
    <TreeItem {...props} TransitionComponent={TransitionComponent}/>
))(({theme}) => ({
    [`& .${treeItemClasses.iconContainer}`]: {
        '& .close': {
            opacity: 0.3,
        },
    },
    [`& .${treeItemClasses.group}`]: {
        marginLeft: 15,
        paddingLeft: 18,
        borderLeft: `1px dashed ${alpha(theme.palette.text.primary, 0.4)}`,
    },
}));

export interface NavigationTreeProps {
    items: NavigationItem[]
    onClick?: (item: NavigationItem) => void
}

export interface NavigationTreeComponentProps {
    item: NavigationItem
    path: string
}

export function NavigationTreeComponent(props: NavigationTreeComponentProps) {
    const path = props.path + '.' + props.item.name;

    return <StyledTreeItem nodeId={path} label={props.item.title}>
        {props.item.children && props.item.children.map(subItem => {
            return <NavigationTreeComponent path={path} item={subItem}/>;
        })}
    </StyledTreeItem>
}

export function NavigationTree(props: NavigationTreeProps) {
    return (
        <TreeView
            aria-label="customized"
            defaultExpanded={['1', '3']}
            defaultCollapseIcon={<MinusSquare/>}
            defaultExpandIcon={<PlusSquare/>}
            defaultEndIcon={<Article/>}
            sx={{height: 264, flexGrow: 1, maxWidth: 400, overflowY: 'auto'}}
        >
            <StyledTreeItem nodeId="1" label="SDK">
                {props.items.map(item => <NavigationTreeComponent path={item.name} item={item}/>)}
            </StyledTreeItem>
        </TreeView>
    );
}