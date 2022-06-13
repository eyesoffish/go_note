// ignore_for_file: avoid_print

import 'dart:io';

const String pages_path = 'lib/app/pages';
const String pages_config_path = 'lib/app/config/page.dart';
void main() {
  _start();
}

void _test() {
  final String res = _readFile("lib/app/pages/home/home_page.dart");
  print(res);
}

void _start() {
  final Directory d = Directory('lib/app/pages');
  if (!d.existsSync()) return;
  String config = "";
  d.list(recursive: true, followLinks: false).listen(
    (FileSystemEntity event) {
      if (event.path.endsWith(".dart")) {
        final String res = _readFile(event.path);
        if (res.isNotEmpty) {
          config += res;
          print(event.path);
        }
      }
    },
    onDone: () {
      print("========>执行完成<=========");
      _writePagesContent(config);
    },
  );
}

String _readFile(String path) {
  final File file = File(path);
  final String content = file.readAsStringSync();
  final String? classHeaderContent =
      RegExp(r"class (\w*)Page extends RoutePage {(\w|\W)*;\n\}(?=(\w|\W)*extends PageState)", multiLine: true)
          .stringMatch(content);
  // 'test_page': PageConfig(
  // name: '测试页面',
  // params: <String, dynamic>{'classId': typeInt, "className": typeString},
  // filePath: 'app/config',
  // ),
  String? className;
  String? name;
  String params = "";
  if (classHeaderContent != null) {
    className = RegExp(r"(?<=class )\w*(?= extends)").stringMatch(classHeaderContent);
    name = RegExp(r"(?<=').*(?=')").stringMatch(classHeaderContent);

    RegExp(r"final.*[^']", multiLine: true).allMatches(classHeaderContent).forEach((RegExpMatch element) {
      final String? pro = element.group(0); // 属性
      if (pro != null) {
        final String? proName = RegExp(r'\w*(?=;)').stringMatch(pro);
        final String proType = RegExp(r' (\w|\W)* ').stringMatch(pro) ?? "";
        if (proType.trim() != "String name =" && proType.trim() != "String cnName =") {
          params += "'$proName': ${_checkType(proType)}, ";
        }
      }
    });
  } else {
    return "";
  }
  if (className == null) return "";
  final String res = '''
  '${_pathName(className)}': PageConfig(
    name: '$name',${params.isNotEmpty ? "\n    params: <String, dynamic>{$params}," : ""}
    filePath: '${path.replaceAll("lib/", "")}',
  ),
''';
  return res;
}

String _checkType(String val) {
  if (val.contains('int')) {
    return "typeInt";
  } else if (val.contains('double')) {
    return "typeDouble";
  } else if (val.contains('String')) {
    return "typeString";
  } else if (val.contains('Map')) {
    return "typeMap";
  } else if (val.contains('List')) {
    return "typeList";
  } else if (val.contains('bool')) {
    return "typeBool";
  } else {
    return "typeDynamic";
  }
}

String _pathName(String className) {
  final RegExp exp = RegExp(r'(?<=[a-z])[A-Z]');
  return className.replaceAllMapped(exp, (Match m) => '_${m.group(0)}').toLowerCase();
}

void _writePagesContent(String newPageConfigContent) {
  final File file = File(pages_config_path);
  final String content = file.readAsStringSync();
  final RegExp exp = RegExp(
    r"final Map<RouteName, PageConfig> userBusinessPageConfig = <String, PageConfig>{(\w|\W)*};",
    multiLine: true,
  );
  final String? oldPageConfigContent = exp.stringMatch(content);
  if (oldPageConfigContent == null) return;
  final String _res = '''
final Map<RouteName, PageConfig> userBusinessPageConfig = <String, PageConfig>{
$newPageConfigContent};
''';
  final String res = content.replaceFirst(oldPageConfigContent, _res);
  file.writeAsStringSync(res);
}
