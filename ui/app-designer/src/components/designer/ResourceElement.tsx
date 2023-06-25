import { type Resource } from '@apibrew/core-lib'

export interface ResourceElementProps {
    resource: Resource
}

export function ResourceElement(props: ResourceElementProps) {
    return <g className={`resource-${props.resource.name}`}>
        <rect fill="#c3e0e0" x="1" y="1.76916" width="198" height="278" stroke="#000" rx="3"
            filter="url(#svg_1_blur)" strokeWidth="2" />
        <g transform='translate(10, 10)'>
            <g>
                <rect className='resource-head' fill="#c3e0e0" width="177" height="32.68937" strokeWidth="0.5"
                    stroke="#34a8a0" />
                <text fill="#000000" stroke="#000" strokeWidth="0" x="10" y="21" fontSize="16"
                    fontFamily="Noto Sans JP" textAnchor="start" fontWeight="normal" fontStyle="normal">
                    {props.resource.name} {props.resource.namespace && <>({props.resource.namespace})</>}
                </text>
            </g>
            <g transform='translate(0, 50)'>
                <g>
                    <rect fill="#c3e0e0" strokeWidth="0.5" width="177" y='10' height="198"
                        stroke="#000" />
                    <text fill="#000000" stroke="#000" strokeWidth="0" x="4.23848" fontSize="10"
                        fontFamily="Noto Sans JP" textAnchor="start">Properties:
                    </text>
                </g>
                <g transform='translate(0, 10)'>
                    {props.resource.properties?.map((property, index) => {
                        return <g key={property.name} transform={`translate(0, ${25 * index})`}
                            className={`resource-property-${property.name}`}>
                            <text fill="#000000" stroke="#000" strokeWidth="0" x="3.46932" y="12.28861"
                                fontSize="10" fontFamily="Noto Sans JP" textAnchor="start"
                                fontWeight="normal">
                                <tspan fontWeight='bold'>{property.name}</tspan>
                                <tspan>&nbsp; </tspan>
                                <tspan fill='red'>[{property.type?.toLowerCase()}]</tspan>
                            </text>
                            <line strokeDasharray="2,2" stroke="#000" y2="17.51724" x2="176.55172" y1="17.51724"
                                x1="0" strokeWidth="0.5" fill="none" />
                            {property.type === 'REFERENCE' && <g>
                                <rect className={'left-ref'} stroke="#000" height="7.71414" width="4.87209"
                                    y="6"
                                    x="-3.18594"
                                    strokeWidth="0.5" fill="#137a7f" />

                                <rect className={'right-ref'} stroke="#000" height="7.71414" width="4.87209"
                                    y="6"
                                    x="174.00214"
                                    strokeWidth="0.5" fill="#137a7f" />
                            </g>}
                        </g>
                    })}
                </g>
            </g>
        </g>
    </g>
}
