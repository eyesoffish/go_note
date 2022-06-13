import 'dart:io';
import 'lib/app/config/page.dart' show userBusinessPageConfig, PageConfig;

const String route_path = "lib/app/router/router_test.dart";
const String pages_path = "lib/app/router/pages_test.dart";

void main() {
  // _test();
  _start();
}

void _start() {
  String _pagesContent = "class Pages {\n";
  String _routeContent = '''
// ignore_for_file: always_specify_types
import 'package:flutter/material.dart';
import 'package:supermonkeyapp/${pages_path.replaceAll("lib/", "")}';
''';

  String _importList = "";
  String _getUserRoute = '''
Route<dynamic>? getUserRoute(RouteSettings settings, buildRoute) {
  switch (settings.name) {\n''';

  userBusinessPageConfig.forEach(
    (String pageName, PageConfig value) {
      _pagesContent += "   static const String $pageName = '/$pageName';\n";
      _importList += "import 'package:supermonkeyapp/${value.filePath}/$pageName.dart';\n";
      _getUserRoute += '''    case Pages.$pageName:\n''';
      if (value.params == null) {
        _getUserRoute += '''      return buildRoute(settings, ${_capitalize(pageName)}());\n''';
      } else {
        _getUserRoute += '''
      final Map arguments = settings.arguments as Map;
      return buildRoute(settings,
        ${_capitalize(pageName)}(\n${_getParams(value.params!)}
        ),
      );\n''';
      }

      /// 创建对应的page页面
      _addClass(pageName, value);
    },
  );

  // 脚本拼接
  _pagesContent += '}';
  _routeContent += _importList;

  _getUserRoute += '''

    default:
      return null;
  }
}''';
  _routeContent += _getUserRoute;

  /// 执行写入脚本
  _createFile(_pagesContent, pages_path);
  _createFile(_routeContent, route_path);
}

/// 生成对应的类
void _addClass(String pageName, PageConfig config) {
  final String path = "lib/${config.filePath}/$pageName.dart";
  final File file = File(path);
  if (file.existsSync()) {
    _rewriteClass(pageName, config, file);
    return;
  }
  String params = "";
  final String className = _capitalize(pageName);
  if (config.params != null) {
    String _thisParams = "";
    config.params!.forEach((String key, dynamic value) {
      params += "  final ${_checkType(value)} $key;\n";
      _thisParams += " this.$key,";
    });
    params += '''  $className({Key? key,$_thisParams}) : super(key: key);''';
  }
  final String content = '''
import 'package:flutter/material.dart';
import 'package:supermonkeyapp/core/page.dart';
import 'package:supermonkeyapp/${pages_path.replaceAll("lib/", "")}';

class $className extends RoutePage {
$params
  @override
  final String name = Pages.$pageName;

  @override
  final String cnName = '${config.name}';

  @override
  _${className}State createState() => _${className}State();
}

class _${className}State extends PageState<$className>{

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(),
    );
  }
}
 ''';
  file.writeAsStringSync(content);
  file.createSync();
}

void _rewriteClass(String pageName, PageConfig config, File file) {
  final String content = file.readAsStringSync();
  final String className = _capitalize(pageName);
  final RegExp exp = RegExp(r"class " + className + r" extends RoutePage {(\w|\W)*;\n}", multiLine: true);
  final String? classHeaderContent = exp.stringMatch(content);
  if (classHeaderContent == null) return;
  String params = "";
  if (config.params != null) {
    String _thisParams = "";
    config.params!.forEach((String key, dynamic value) {
      params += "  final ${_checkType(value)} $key;\n";
      _thisParams += " this.$key,";
    });
    params += '''  $className({Key? key,$_thisParams}) : super(key: key);''';
  }
  final String newClassHeader = '''
class $className extends RoutePage {
$params
  @override
  final String name = Pages.$pageName;

  @override
  final String cnName = '${config.name}';

  @override
  _${className}State createState() => _${className}State();
}''';
  final String res = content.replaceFirst(classHeaderContent, newClassHeader);
  file.writeAsStringSync(res);
}

String _checkType(dynamic val) {
  if (val == 'dynamic') {
    return 'dynamic';
  } else if (val is int) {
    return "int?";
  } else if (val is double) {
    return "double?";
  } else if (val is String) {
    return "String?";
  } else if (val is Map) {
    return "Map<String, dynamic>?";
  } else if (val is List) {
    return "List<dynamic>?";
  } else if (val is bool) {
    return 'bool?';
  } else {
    return "dynamic";
  }
}

/// 获取参数
String _getParams(Map<String, dynamic> sender) {
  String res = "";
  int index = 0;
  sender.forEach((String key, dynamic value) {
    res += "          $key: arguments['$key'],${index == sender.length - 1 ? '' : '\n'}";
    index++;
  });

  return res;
}

/// 下划线转驼峰
String _capitalize(String sender) {
  final List<String> _list = sender.split('_');
  String res = "";
  _list.forEach((String element) {
    res += "${element[0].toUpperCase()}${element.substring(1)}";
  });
  return res;
}

void _createFile(String content, String path) {
  final File file = File(path);
  file.writeAsStringSync(content);
  file.createSync();
}
