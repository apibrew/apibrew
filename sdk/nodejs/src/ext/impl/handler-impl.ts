import {Client} from "../../client";
import {ExtensionService} from "../extension-service";
import {ExtensionInfo} from "../../extension-info";
import {Event} from "../../model/extension";
import {Entity} from "../../entity";
import {EntityInfo} from "../../entity-info";
import {Handler} from "../handler";

type Predicate<T extends Entity> = (event: Event, entity: T) => boolean;

export class HandlerImpl<T extends Entity> implements Handler<T>{
    constructor(private client: Client,
                private extensionService: ExtensionService,
                private extensionInfo: ExtensionInfo,
                private predicates: Predicate<Entity>[],
                private entityInfo: EntityInfo) {
    }

    public withExtensionInfo(extensionInfo: ExtensionInfo): Handler<T> {
        return new HandlerImpl(this.client, this.extensionService, this.extensionInfo, this.predicates, this.entityInfo)
            .withExtensionInfo(extensionInfo);
    }

}

// import static io.apibrew.client.model.Extension.*;
//
// public class HandlerImpl<T extends Entity> implements Handler<T> {
//     private final Client client;
//     private final ExtensionService extensionService;
//
//     private final ExtensionInfo extensionInfo;
//     private final List<BiPredicate<Event, T>> predicates;
//
//     private final EntityInfo<T> entityInfo;
//
//     private final ObjectMapper objectMapper = new ObjectMapper();
//
//     public HandlerImpl(Client client, ExtensionService extensionService, EntityInfo<T> entityInfo) {
//         this(
//                 client,
//                 extensionService,
//                 entityInfo,
//                 new ExtensionInfo().withNamespace(entityInfo.getNamespace()).withResource(entityInfo.getResource()),
//                 new ArrayList<>()
//         );
//     }
//
//     public HandlerImpl(Client client, ExtensionService extensionService, Class<T> entityClass) {
//         this(client, extensionService, EntityInfo.fromEntityClass(entityClass));
//     }
//
//     public HandlerImpl(Client client, ExtensionService extensionService, EntityInfo<T> entityInfo, ExtensionInfo extensionInfo, List<BiPredicate<Event, T>> predicates) {
//         this.client = client;
//         this.extensionService = extensionService;
//         this.entityInfo = entityInfo;
//         this.extensionInfo = extensionInfo.withSealResource(true);
//         this.predicates = predicates;
//
//         objectMapper.registerModule(new JavaTimeModule());
//         objectMapper.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
//     }
//
//     public Handler<T> withExtensionInfo(ExtensionInfo extensionInfo) {
//         return new HandlerImpl<>(client, extensionService, entityInfo, extensionInfo, predicates);
//     }
//
//     public Handler<T> withPredicates(List<BiPredicate<Event, T>> predicates) {
//         return new HandlerImpl<>(client, extensionService, entityInfo, extensionInfo, predicates);
//     }
//
//     public Handler<T> withPredicate(BiPredicate<Event, T> predicate) {
//         ArrayList<BiPredicate<Event, T>> predicatesCopy = new ArrayList<>(predicates);
//         predicatesCopy.add(predicate);
//
//         return withPredicates(predicatesCopy);
//     }
//
//     @Override
//     public Handler<T> when(Condition<T> condition) {
//         return ((HandlerImpl<T>) withExtensionInfo(condition.configureExtensionInfo(extensionInfo)))
//                 .withPredicate(condition::eventMatches);
//     }
//
//     @Override
//     public Handler<T> configure(Function<ExtensionInfo, ExtensionInfo> configurer) {
//         return withExtensionInfo(configurer.apply(extensionInfo));
//     }
//
//     private T recordToEntity(Record entity) {
//         if (entity == null) {
//             return null;
//         }
//
//         return objectMapper.convertValue(entity.getProperties(), entityInfo.getEntityClass());
//     }
//
//     private Record recordFromEntity(T entity) {
//         if (entity == null) {
//             return null;
//         }
//
//         Record record = new Record();
//         record.setProperties(objectMapper.convertValue(entity, Object.class));
//         record.setId(entity.getId());
//
//         return record;
//     }
//
//     private boolean checkPredicates(Event event, T entity) {
//         for (BiPredicate<Event, T> predicate : predicates) {
//             if (!predicate.test(event, entity)) {
//                 return false;
//             }
//         }
//
//         return true;
//     }
//
//     @Override
//     public String operate(Operator<T> entityOperator) {
//         return entityOperator.operate(this);
//     }
//
//     @Override
//     public String operate(BiFunction<Event, T, T> entityOperator) {
//         return extensionService.registerExtensionWithOperator(extensionInfo, (event, record) -> {
//             T castedEntity = recordToEntity(record);
//             if (!checkPredicates(event, castedEntity)) {
//                 return record;
//             }
//
//             T result = entityOperator.apply(event, castedEntity);
//
//             return recordFromEntity(result);
//         });
//     }
//
//     @Override
//     public void unRegister(String id) {
//         extensionService.unRegisterOperator(id);
//     }
// }