```plantuml
interface IManifest {
  version: string;
  name: string;
  root: string;
  variables: Array<OneOf<ILiteralVariableData, ITemplateVariableData, IEnvironmentVariableData>>;
  casts: Record<string, AllOf<IBaseCastData, Partial<OneOf<IIfConditionalCast, IUnlessConditionalCast>>, Partial<AllOf<ICollectionCastData, Partial<OneOf<IIncludeCollectionCastData, IOmitCollectionCastData>>>>>>;
}

interface IBaseVariableData {
  name: string;
}

interface ILiteralVariableData<T> extends IBaseVariableData {
  value: T;
}

interface ITemplatedVariableData extends IBaseVariableData {
  template: string;
}

interface IEnvironmentVariableData extends IBaseVariableData {
  env: string;
}

interface IBaseCastData {
  from: OneOf<string, IFileSourceData, IDocumentsSourceData, IMagicSourceData>;
  to: string;
  variables: Array<OneOf<ILiteralVariableData, ITemplateVariableData, IEnvironmentVariableData>>;
}

interface IIfConditionalCastData {
  if: string | JSONLogic;
}

interface IUnlessConditionalCastData {
  unless: string | JSONLogic;
}

interface ICollectionCastData {
  each: string;
  as?: string;
}

interface IIncludeCollectionCastData {
  include: string | JSONLogic;
}

interface IOmitCollectionCastData {
  omit: string | JSONLogic;
}

interface IFileSourceData {
  file: string;
}

interface IDocumentsSourceData {
  document: string;
}

interface IMagicSourceData {
  Magic: string;
}
```