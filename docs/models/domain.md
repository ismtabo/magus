```plantuml
interface IRenderizable {
  async render(context: Context): VFile[];
}

class Magic implements IRenderizable {
  manifest: VFile;
  config: Config;
  casts: ICast[];
  variables: VariablesMap;
}
Magic --> VFile
Magic --> Config
Magic *-right- ICast : renders >
Magic -down-> VariablesMap

class Config {
  tag: Tag;
}
Config -down-> Tag
Config -[hidden]right- HelpersMap

class Tag {
  start: string;
  end: string;
}

class Context {
  config: Config;
  cwd: string;
  variables: VariablesMap;
  helpers: HelpersMap;
}
Context --> Config
Context --> VariablesMap
Context --> HelpersMap

class VariablesMap<string, IVariable> extends Map {}
VariablesMap o-- IVariable

interface IVariable<T> {
  async value(context: Context): T;
}

class LiteralVariable<T> implements IVariable {
  literal: T;

  async value(context: Context): T;
}

class EnvVariable<string> implements IVariable {
  name: string;

  async value(context: Context): string;
}

class TemplateVariable<string> implements IVariable {
  template: TemplateString;

  async value(context: Context): string;
}
TemplateVariable --> TemplateString

class TemplateString {
  template: string;
  async render(ctx): string;
}

class HelpersMap<string, Helper> extends Map {}
HelpersMap o-- Helper

class Helper<Args, Result> {
  call(context: Context, ...args: Args): Result;
}

interface ICast extends IRenderizable {
  name: string;
  output: TemplateString;
  source: ISource;
  variables: VariablesMap;
}
ICast --> ISource : compiles >
ICast --> VariablesMap
ICast --> TemplateString : renders >

class Cast implements ICast {}

interface ICollectionCast extends ICast {
  each: JSONPath;
  key: string;
  filter?: ICondition;
}
ICollectionCast --> ICondition : evaluates >

interface IConditionalCast extends ICast {
  condition: ICondition;
}
IConditionalCast --> ICondition : evaluates >
IConditionalCast -[hidden]right- Cast

class CollectionCast extends Cast implements ICollectionCast {}

class ConditionalCast extends Cast implements IConditionalCast {}

class ConditionalCollectionCast extends Cast implements IConditionalCast, ICollectionCast {}

interface ICondition {
  async evaluate(context: Context): boolean
}

class JSONLogicCondition implements ICondition {
  statement: JSONLogic;
}

class TemplateCondition implements ICondition {
  template: TemplateString;
}
TemplateCondition --> TemplateString

interface ISource {
  async compile(context: Context): File[];
}

class DocumentSource implements ISource {
  file: VFile;
}
DocumentSource --> VFile

class FileSource implements ISource {
  file: VFile;
}
FileSource --> VFile
FileSource -[hidden]right- TemplateSource

class MagicSource implements ISource {
  magic: Magic;
  casts?: string[];
}
MagicSource --> Magic : renders >

class TemplateSource implements ISource {
  template: TemplateString;
}
TemplateSource --> TemplateString

namespace vfile {
  class VFile {
    path: string;
    value: string;
  }
}
```