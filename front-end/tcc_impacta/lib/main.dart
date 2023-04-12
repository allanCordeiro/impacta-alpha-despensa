import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:tcc_impacta/page/List_Items_home.dart';
import 'package:tcc_impacta/page/item_register_page.dart';
import 'package:tcc_impacta/routes/app_routes.dart';

import 'provider/item.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(
          create: (ctx) => ItemProvider(),
        )
      ],
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        title: 'Flutter Demo',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        routes: {
          AppRoutes.HOME: (_) => const ListItems(),
          AppRoutes.ITEMS_FORM: (_) => const ItemRegister(),
        },
      ),
    );
  }
}
