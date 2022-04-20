import 'package:flutter/material.dart';
import 'package:flutter_application_1/models/product.dart';
import 'package:flutter_application_1/services/product_services.dart';

void main() {
  runApp(AppDemo());
}

class AppDemo extends StatefulWidget {
  AppDemo({Key key}) : super(key: key);

  @override
  _AppDemoState createState() => _AppDemoState();
}

class _AppDemoState extends State<AppDemo> {
  var productService = ProductService();
  var listem = <ProductModel>[];
  @override
  void initState() {
    super.initState();
    getir();
  }

  getir() async {
    listem = await productService.getAllProducts();
    setState(() {});
    print(listem.first.name);
  }

  final _notifier = ValueNotifier<ThemeModel>(ThemeModel(ThemeMode.light));

  @override
  Widget build(BuildContext context) {
    return ValueListenableBuilder<ThemeModel>(
      valueListenable: _notifier,
      builder: (_, model, __) {
        final mode = model.mode;
        return MaterialApp(
          theme: ThemeData.light(), // Provide light theme.
          darkTheme: ThemeData.dark(), // Provide dark theme.
          themeMode: mode, // Decides which theme to show.
          home: Scaffold(
            appBar: AppBar(
              backgroundColor: Color.fromARGB(255, 12, 25, 22),
              title: Row(children: [
                Text(""),
                Expanded(
                  child: Center(
                    child: CircleAvatar(
                      radius: 50,
                      backgroundColor: Color.fromARGB(255, 43, 212, 142),
                      child: IconButton(
                        color: Color.fromARGB(255, 15, 7, 7),
                        padding: EdgeInsets.all(20),
                        iconSize: 50,
                        icon: Icon(Icons.dark_mode_sharp),
                        onPressed: () => _notifier.value = ThemeModel(
                            mode == ThemeMode.light
                                ? ThemeMode.dark
                                : ThemeMode.light),
                      ),
                    ),
                  ),
                ),
              ]),
            ),
            body: Padding(
              padding: const EdgeInsets.all(8.0),
              child: ListView.builder(
                itemCount: listem != null ? listem.length : 0,
                itemBuilder: (BuildContext context, int index) {
                  return ListTile(
                    title: Text(listem[index].name),
                    subtitle: Text(listem[index].id.toString()),
                    leading: FlutterLogo(size: 50.0),
                    onTap: () => print(listem[index].name),
                    contentPadding: EdgeInsets.all(10),
                    shape: RoundedRectangleBorder(
                      side: BorderSide(color: Colors.black, width: 1),
                      borderRadius: BorderRadius.circular(5),
                    ),
                  );
                },
              ),
            ),
          ),
        );
      },
    );
  }
}

class ThemeModel with ChangeNotifier {
  final ThemeMode _mode;
  ThemeMode get mode => _mode;

  ThemeModel(this._mode);
}
