import 'package:flutter/material.dart';

import 'views/login_view.dart';
import 'views/users_view.dart';


void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});  // Add const constructor

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'User Management App',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const LoginView(),
      routes: {
        '/users': (context) => UsersView(),
      },
    );
  }
}